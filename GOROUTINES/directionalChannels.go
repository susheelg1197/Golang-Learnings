/*

chan : bidirectional
chan<- send unidirectional only write
<-chan recieve unidirectional only read

*/
package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	processWrite(ch)
	processRead(ch)
	close(ch)
}

func processWrite(ch chan<- int) {
	ch <- 2
}
func processRead(ch <-chan int) {
	fmt.Println(<-ch)
}
