package main

import (
	"fmt"
)

func main() {
	sr := "jooo"
	r := []rune(sr)

	r[0] -= 'a' - 'A' // To print the difference
	fmt.Println(string(r))
}
