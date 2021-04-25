package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/home/susheelg/go/src/SDCOM/PROGRAMS/countDigits.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println(string(f)) // cannot convert f (variable of type *os.File) to string
}
