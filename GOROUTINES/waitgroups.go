package main

import (
	"fmt"
	"sync"
	"time"
)

func routine(s string, wg *sync.WaitGroup) {
	fmt.Println("Heelo ", s)
	time.Sleep(1000 * time.Millisecond)
	wg.Done()
}
func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go routine("HI", &wg)
	go routine("Bye", &wg)
	wg.Wait()

}
