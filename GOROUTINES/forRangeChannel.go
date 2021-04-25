package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 3)
	ch <- 2
	ch <- 2
	ch <- 2
	close(ch) // close when using simple function
	sum(ch)   // if this is made a go routine no error

	time.Sleep(time.Millisecond * 100)

}

func sum(ch chan int) {
	for v := range ch {
		fmt.Println("Value: ", v, len(ch))
	}
}
