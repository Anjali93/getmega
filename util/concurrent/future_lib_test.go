package main

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	var executor Executor = GoRoutineExecutor{}
	future := executor.submit(func() interface{} {
		return 100
	})
	res, err := future.get()
	if err != nil {
		t.Error("Error running task")
	}
	if res != 100 {
		t.Error("Value error")
	}
}

func TestGetWithTimeout(t *testing.T) {
	var executor Executor = GoRoutineExecutor{}
	future := executor.submit(func() interface{} {
		time.Sleep(2 * time.Second)
		return 100
	})
	res, err := future.getWithTimeout(3, time.Second)

	if err != nil {
		t.Error("Error running task")
	}
	if res != 100 {
		t.Error("Value error")
	}
}

func TestGetWithTimeoutError(t *testing.T) {
	var executor Executor = GoRoutineExecutor{}
	future := executor.submit(func() interface{} {
		time.Sleep(2 * time.Second)
		return 100
	})
	res, err := future.getWithTimeout(1, time.Second)

	if err == nil {
		t.Error("Timeout error not thrown")
	}
	if res != nil {
		t.Error("Value received")
	}
}

func TestIsDone(t *testing.T) {
	var executor Executor = GoRoutineExecutor{}
	future := executor.submit(func() interface{} {
		return 100
	})
	time.Sleep(1 * time.Second)
	res := future.isDone()

	if !res {
		t.Error("is done is false")
	}
}

func TestIsCancelled(t *testing.T) {
	t.Fail()
}

func TestIsDoneAfterCancelling(t *testing.T) {
	t.Fail()
}

func TestGetAfterCancelling(t *testing.T) {
	t.Fail()
}
