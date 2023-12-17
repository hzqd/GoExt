package main

import (
	"fmt"
	"goExt/ext"
	"time"
)

func main() {
	// Tap
	a := &struct{ int8 }{0}
	fmt.Print(*a)
	ext.Tap(a, func(t *struct{ int8 }) {
		*t = struct{ int8 }{1}
	})
	fmt.Println(*a)

	// Pipe
	b := &struct{ int8 }{0}
	fmt.Print(*b)
	c := ext.Pipe(b, func(t *struct{ int8 }) int8 {
		*t = struct{ int8 }{1}
		return t.int8 + 1
	})
	fmt.Println(*b, c)

	// Spawn, Await
	async := ext.Spawn(func() int8 {
		time.Sleep(1)
		print("spawn ")
		return c
	})
	println("main")
	println(async.Await())
}
