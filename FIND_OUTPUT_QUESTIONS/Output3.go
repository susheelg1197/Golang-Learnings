package main

import "fmt"

func main() {
	var a *int

	for i := 0; i <= 3; i++ {
		a = &i // It will print error cannot use &i (value of type *int) as int value in assignment
		fmt.Println(a)
	}

	for i := 0; i <= 3; i++ {
		a = &i
		fmt.Println(*a)
	}
}
