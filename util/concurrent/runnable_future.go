package main

import (
	"errors"
	"time"
)

type TaskState int

const (
	RUNNING TaskState = iota
	COMPLETED
	CANCELLED
)

type ChannelFuture struct {
	channel  chan interface{}
	state    TaskState
	result   interface{}
}

func (f *ChannelFuture) get() (interface{}, error) {

	if f.state == CANCELLED {

		return nil, errors.New("task is cancelled")
	}

	if f.result == nil {
		f.result = <-f.channel
	}

	return f.result, nil
}

func (f *ChannelFuture) getWithTimeout(timeout time.Duration, unit time.Duration) (interface{}, error) {

	if f.state == CANCELLED {

		return nil, errors.New("task is cancelled")
	}

	if f.result != nil{
		return f.result, nil
	}

	select {
	case f.result = <-f.channel:
		if len(f.channel) == cap(f.channel){
			f.result = <-f.channel
		}
		return f.result, nil
	case <-time.After(timeout * unit):
		return nil, errors.New("Timeout exception")
	}
}

func (f *ChannelFuture) cancel() {
	if f.state == RUNNING {
		f.state = CANCELLED
	}
}

func (f ChannelFuture) isCancelled() bool {
	return f.state == CANCELLED
}

func (f ChannelFuture) isDone() bool {
	return f.state != RUNNING
}
