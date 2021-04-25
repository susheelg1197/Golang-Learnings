package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background()
	ctx1, ctcancelFunc := context.WithCancel(ctx)

	go goRout(ctx1)
	time.Sleep(time.Second * 3)
	ctcancelFunc() // Its invoked from the parent function that created the context
	time.Sleep(time.Second * 1)

}

func goRout(ctx context.Context) {
	i := 1

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully exited")
			return
		default:
			time.Sleep(time.Second * 1)
			fmt.Println(i)
			i++
		}
	}
}
