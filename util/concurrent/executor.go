package main

type Executor interface {
	submit(func() interface{}) Future
}

type GoRoutineExecutor struct {
}

func (g GoRoutineExecutor) submit(f func() interface{}) Future {
	ch := make(chan interface{}, 1)
	future := &ChannelFuture{channel: ch, state: RUNNING}
	go func() {
		ch <- f()
		if future.state == RUNNING {
			future.state = COMPLETED
		}
	}()
	return future
}
