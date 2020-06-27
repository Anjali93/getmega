package main

import "time"

type Future interface {
	get() (interface{}, error)
	getWithTimeout(timeout time.Duration, unit time.Duration) (interface{}, error)
	cancel(mayInterruptIfRunning bool)
	isCancelled() bool
	isDone() bool
}
