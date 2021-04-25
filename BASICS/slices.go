/*

Slice is a variable-length sequence which stores elements of a similar type,
you are not allowed to store different type of elements in the same slice.
It is just like an array having an index value and length, but the size of the
slice is resized they are not in fixed-size just like an array
*/

package main

import "fmt"

func main() {

	// Ways of declaration
	arr1 := []int{1, 2, 34, 5, 6, 7, 8}
	arr2 := make([]int, 5) // func make([]T, len, cap) []T
	// capacity is optional and the length of the slice is taken as the capacity

	fmt.Println("Printing slices:: ", arr1, cap(arr2))

	// Zero value of slice
	fmt.Println("Slice zero value:: ", []int{})

	/*
		A slice can also be formed by "slicing" an existing slice or array.
		Slicing is done by specifying a half-open range with two indices separated by a colon.
		For example, the expression b[1:4] creates a slice including elements 1 through 3 of b
		 (the indices of the resulting slice will be 0 through 2).
	*/

	fmt.Println("Sliced between:: ", arr1[1:3])              // index pos 1 and exclude index position of last
	fmt.Println("Sliced from start till index:: ", arr1[:3]) // index pos from start 0 and exclude index position of last
	fmt.Println("Sliced from index till last:: ", arr1[3:])  // index pos from start 3 and till last
	fmt.Println("Sliced only colon(:):: ", arr1[:])          // similar to original array

	//This is also the syntax to create a slice given an array:
	x := [3]string{"Лайка", "Белка", "Стрелка"}
	s := x[:] // a slice referencing the storage of x

	fmt.Println("Sliced Array:: ", len(s))

	/*
		Slice internals
		A slice is a descriptor of an array segment. It consists of a pointer to the array,
		the length of the segment, and its capacity (the maximum length of the segment).
	*/

	// While slicing into a new slice it creates a new slice instead of copying the value it points to the
	// original array
	s1 := arr1[2:4]
	s1[1] = 32
	fmt.Printf("\nS1:: %d -> %d -> %d ->> %p ---->>> %p\n", s1, len(s1), cap(s1), &s1, &s1[1])
	fmt.Printf("\nArr1:: %d -> %d -> %d ->> %p  ---->>> %p\n", arr1, len(arr1), cap(arr1), &arr1, &arr1[3])

	/*
		A slice cannot be grown beyond its capacity. Attempting to do so will cause a runtime panic,
		just as when indexing outside the bounds of a slice or array. Similarly, slices cannot be re-sliced
		 below zero to access earlier elements in the array.
	*/
	//Growing slices (the copy and append functions)
	t := make([]int, len(s1), (cap(s1)+1)*2)
	copy(t, s1)
	s1 = t

	fmt.Println(cap(s1))

	//Appending slice

	a := make([]int, 1)
	fmt.Println("Previous:: ", a, cap(a), len(a))

	a = append(a, 1, 2, 3)

	fmt.Println("Later:: ", a, cap(a), len(a))

	a = append(a, 1, 2, 3)

	fmt.Println("Later More:: ", a, cap(a), len(a))

	// Understanding passing slice as value and slice as pointer
	slice := []int{40, 50}
	fmt.Println("Slicedv Previous:: ", slice)
	slicedv(slice)
	fmt.Println("Slicedv:: ", slice)
	// slicedp(&slice)
	// fmt.Println("Slicedp:: ", slice)

	slice1 := make([]int, 2)
	fmt.Printf("%d and address %p and %p", slice1, &slice1, &slice[0])
	slice2 := append(slice1, 12, 34, 45, 56)
	slice1 = slice2
	slice2[0] = 45
	fmt.Printf("%d and address %p and %p", slice1, &slice1, &slice[0])

}

func slicedv(a []int) {
	fmt.Printf("\nAddres=%p\n", &a)
	a = append(a, 40, 45)
	a[0], a[1] = 10, 20
	fmt.Println("Inside func:: ", a)

}
func slicedp(a *[]int) {
	fmt.Printf("\nAddres=%p\n", a)

	(*a)[0], (*a)[1] = 10, 20

}
