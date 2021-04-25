package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
func main() {

	n := 9
	for i := 1; i <= n; i++ {
		fmt.Println(fib(i))
	}

}
