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
	//task    func() interface{}
	channel chan interface{}
	state   TaskState
}

func (f ChannelFuture) get() (interface{}, error) {
	return <-f.channel, nil
}

func (f ChannelFuture) getWithTimeout(timeout time.Duration, unit time.Duration) (interface{}, error) {
	select {
	case res := <-f.channel:
		return res, nil
	case <-time.After(timeout * unit):
		return nil, errors.New("Timeout exception")
	}
}

func (f ChannelFuture) cancel(mayInterruptIfRunning bool) {
	f.state = CANCELLED
}

func (f ChannelFuture) isCancelled() bool {
	return f.state == CANCELLED
}

func (f ChannelFuture) isDone() bool {
	return f.state != RUNNING
}
