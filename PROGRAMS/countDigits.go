package main

import (
	"fmt"
)

func DigitsCount(num int) int {
	count := 0
	for num != 0 {
		num = num / 10
		count++
	}
	return count
}

func main() {

	arrayOfnum := []int{100, 200, 400, 55, 6666}

	for _, val := range arrayOfnum {

		// length := (int)(math.Log10(float64(val)) + 1) // Easy way to find length
		fmt.Println(DigitsCount(val))
		if DigitsCount(val)%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("ODD")
		}
	}
	fmt.Println(arrayOfnum)
}
