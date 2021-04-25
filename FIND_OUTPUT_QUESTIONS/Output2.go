package main

import "fmt"

func main() {
	a := 5
	func(a int) {
		a = 10
		fmt.Println("HI")
	}(a)
	println(a)
}
