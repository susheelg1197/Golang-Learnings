package main

import (
	"fmt"
	"math"
)

func longestSubSeq(n int, arr []int) int {

	max1 := 0
	flag1 := 0

	for i := 0; i < n; i++ {
		if flag1 == 0 {
			if arr[i]%2 == 1 {
				flag1 = 1
				max1++
			}
		} else {
			if arr[i]%2 == 0 {
				flag1 = 0
				max1++
			}
		}
	}

	max2 := 0
	flag2 := 0

	for i := 0; i < n; i++ {
		if flag2 == 1 {
			if arr[i]%2 == 1 {
				flag2 = 1
				max2++
			}
		} else {
			if arr[i]%2 == 0 {
				flag2 = 0
				max2++
			}
		}
	}
	fmt.Println(max1, max2)
	return (int)(math.Max(float64(max1), float64(max2)))
}

func main() {

	seq := []int{10, 3, 5, 5051, 25, 25}

	fmt.Println(longestSubSeq(len(seq), seq))
}
