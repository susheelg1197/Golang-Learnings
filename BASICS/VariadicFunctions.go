/*
A variadic function is a function that accepts a variable number of arguments.
In Golang, it is possible to pass a varying number of arguments of the same
type as referenced in the function signature. To declare a variadic function,
the type of the final parameter is preceded by an ellipsis, "...", which shows
that the function may be called with any number of arguments of this type.
This type of function is useful when you don't know the number of arguments
you are passing to the function, the best example is built-in Println function
of the fmt package which is a variadic function.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	final_str := varidexample("hi", "bye", "Yes")
	fmt.Println(final_str)
}

func varidexample(s ...string) string {
	str := ""

	str = strings.Join(s, str)
	return str

}

//Pass different types of arguments in variadic function
