package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go sender(ch)
	go receiver(ch)
	time.Sleep(time.Second * 2)
}
func sender(ch chan int) {
	ch <- 3
	fmt.Println("Data sent!!!")
}
func receiver(ch chan int) {
	time.Sleep(time.Second * 1)
	fmt.Println("Recieved Data:: ", <-ch)
}
