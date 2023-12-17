package ext

import (
	"sync"
)

type Future[T any] struct {
	*future[T]
}

type future[T any] struct {
	runFn  func() T
	wait   sync.WaitGroup
	result T
}

func Spawn[T any](runFn func() T) Future[T] {
	f := Future[T]{&future[T]{
		runFn: runFn,
		wait:  sync.WaitGroup{},
	}}
	k := Tap(&f, func(c *Future[T]) {
		Pipe(&c.wait, func(wg *sync.WaitGroup) struct{} {
			wg.Add(1)
			return struct{}{}
		})
	})
	go Tap(k, func(x *Future[T]) {
		defer Pipe(&x.wait, func(wg *sync.WaitGroup) struct{} {
			wg.Done()
			return struct{}{}
		})
		Tap(x, func(a *Future[T]) {
			a.result = a.runFn()
			a.runFn = nil
		})
	})
	return *k
}

func (f Future[T]) Await() T {
	Pipe(&f.wait, func(wg *sync.WaitGroup) struct{} {
		wg.Wait()
		return struct{}{}
	})
	return f.result
}

func (f Future[T]) TryGet() (T, bool) {
	return f.result, f.runFn == nil
}
