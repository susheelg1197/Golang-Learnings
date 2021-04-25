package main

import "fmt"

func main() {
	a, b := 10, 20

	b, a = a, b
	fmt.Println("swapped:: ", a, b)
}
