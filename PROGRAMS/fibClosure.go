package main

import "fmt"

func main() {

	generateFibs := fibAno()
	for i := 0; i < 10; i++ {
		fmt.Println(generateFibs())
	}
}

func fibAno() func() int {

	f1 := 0
	f2 := 1
	return func() int {
		f2, f1 = f1+f2, f2
		return f1

	}
}
