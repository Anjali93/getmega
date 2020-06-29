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

func TestMultipleGet(t *testing.T) {
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

	time.Sleep(1 * time.Second)

	res1, err1 := future.get()
	if err1 != nil {
		t.Error("Error running task")
	}
	if res1 != 100 {
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

func TestGetWithMultipleTimeout(t *testing.T) {
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
	time.Sleep(1 * time.Second)
	res1, err1 := future.getWithTimeout(3, time.Second)

	if err1 != nil {
		t.Error("Error running task")
	}
	if res1 != 100 {
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
	var executor Executor = GoRoutineExecutor{}
	future := executor.submit(func() interface{} {
		time.Sleep(1 * time.Second)
		return 100
	})

	future.cancel()
	res := future.isCancelled()

	if !res {
		t.Error("is cancelled is false")
	}
}

func TestIsDoneAfterCancelling(t *testing.T) {

	var executor Executor = GoRoutineExecutor{}
	future := executor.submit(func() interface{} {
		time.Sleep(1 * time.Second)
		return 100
	})

	future.cancel()
	res := future.isCancelled()
	res1 := future.isDone()

	if !res {
		t.Error("is cancelled is false")
	}

	if !res1{
		t.Error("Is done is false after cancelling")
	}
}

func TestGetAfterCancelling(t *testing.T) {

	var executor Executor = GoRoutineExecutor{}
	future := executor.submit(func() interface{} {
		time.Sleep(1 * time.Second)
		return 100
	})

	future.cancel()
	res,err := future.get()

	if err == nil {
		t.Error("Cancellation error is not thrown")
	}
	if res != nil {
		t.Error("Value received")
	}
}
