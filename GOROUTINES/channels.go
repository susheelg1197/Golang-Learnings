/*
Channels are like pipes that are used by go routines to communicate.
Channels don't make use of locks as they are implicitly managed via locks,
Channel provides synchroniation and communication


Major operatiosn in channel is send or recieve
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan int // it will create a nil channel

	// It is an unbuffered channel
	c = make(chan int) // with make the instance of hchan struct is
	//created and fields are initialized to zero

	// c <- 10 // send operations sends data into the channel

	// fmt.Println(<-c) // recieve operation reads data from the channel

	// Example to send data to one go routine and recieve from other go routine

	go send(c)
	go recieve(c)

	time.Sleep(3 * time.Second)
}

/*
In unbuffered channel the send is blocked until the reciever go routine reads
from channel

The reciever go routine is blocked until any sender go routine send data in
channel

*/
func send(c chan int) {
	c <- 100
	fmt.Println("Sending data completed")
}
func recieve(c chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Timeout done")
	_ = <-c
	return
}
