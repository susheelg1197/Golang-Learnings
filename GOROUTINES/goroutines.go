/*
Goroutines are user-space threads conceptually similar to kernel threads
(managed by OS) but these goroutines are managed by GO Runtime

It has a smaller memory footprint.
Initial Goroutine stack = 2KB where as default thread stack = 8KB.
Faster creation, destruction, and context switches;
Goroutine switches = ~tens of ns; thread switches = ~1 microsecond

*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

func say() {

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Index: ", i)
	}
}
func main() {
	go say()
	say()

	// To get number of running go routines there is a function called NumGoroutines

	fmt.Println((runtime.NumGoroutine()))
}
