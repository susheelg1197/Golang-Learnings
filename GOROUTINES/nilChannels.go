/*

Nil channel blocks forever for sender and reciever function
*/
package main

import "fmt"

func main() {
	var ch chan int

	ch <- 2
	fmt.Println("VBlocked")
}
