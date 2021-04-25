/*
Buffered channel unlike unbuffered, blocks the senders after the limit exceeds
Reciever is blocked when the channel is empty
*/

package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	fmt.Println("Sending into buffered channel")
	_ = <-ch
	fmt.Println("Data recieved:: ", <-ch)
	// fmt.Println("Data recieved:: ", <-ch)
}
