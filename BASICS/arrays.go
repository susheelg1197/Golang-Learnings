/*
What is an array?
An array is a data structure that consists of a collection of elements of a
single type or simply you can say a special variable, which can hold more than
one value at a time. The values an array holds are called its elements or items.
An array holds a specific number of elements, and it cannot grow or shrink.
Different data types can be handled as elements in arrays such as Int, String,
Boolean, and others. The index of the first element of any dimension of an array
is 0, the index of the second element of any array dimension is 1, and so on.

You cannot equate arrays of different lengths, and you also cannot assign
the value of one to the other

Go's arrays are values. An array variable denotes the entire array;
it is not a pointer to the first array element (as would be the case in C).
This means that when you assign or pass around an array value you will make a
copy of its contents. (To avoid the copy you could pass a pointer to the array,
but then that's a pointer to an array, not an array.) One way to think about
arrays is as a sort of struct but with indexed rather than named fields: a
fixed-size composite value.

An array literal can be specified like so:

b := [2]string{"Penn", "Teller"}

Or, you can have the compiler count the array elements for you:

b := [...]string{"Penn", "Teller"}
In both cases, the type of b is [2]string.

*/

package main

import (
	"fmt"
	"unsafe"
)

func main() {

	arr1 := [2]int{10, 20}
	arr2 := [...]int{1, 2} // the compiler counts the length of the array
	// arr2 = arr1 //cannot use arr1 (type [6]int) as type [2]int in assignment
	fmt.Println("Size of", unsafe.Sizeof(arr1), "  ", arr1, len(arr2))
}
