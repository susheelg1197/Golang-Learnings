package main

import "fmt"

// Anonymous function - is a type of function without any specific name to it
var (
	area = func(l, b int) int {
		return l * b
	}
)

// User defined function

type Susheel func(int) int

func sum(x, y int) int {

	return x + y
}

// Higher order function
// Function as an argument and return function

func partialfunc(x int) func(y int) int {

	return func(y int) int {
		return sum(x, y)
	}
}

// return function from function
func multipleFuncs(x int) func(y int) func(z int) int {
	return func(y int) func(z int) int {
		return func(z int) int {
			return x + y + z
		}
	}
}

// Call by value
func valuetype(a, b int) (int, int) {

	return b, a
}

// Call be reference
func reftype(a, b *int) {
	*b, *a = *a, *b
}
func main() {
	a, b := 50, 100
	valuetype(a, b)
	fmt.Println(a, b)
	reftype(&a, &b)
	fmt.Println(a, b)

	// Passing arguments to an anonymous function
	answer := func(l, b int) int {
		return l * b
	}(a, b)
	fmt.Println(answer)
	//Closure Function in Golang
	// These are a special case of anonymous function that accesses variables
	// defined outside the body of the function
	var add int
	for i := 0; i < 10; i++ {
		func() {
			add += i
		}()
	}
	fmt.Printf("\n Addition: %d\n", add)

	// Higher order functions
	// An higher order function is a function that operate on the
	// other functions and taking them as arguments or by returning them
	funVar := partialfunc(10)
	fmt.Println("Sum of Two numbers:: ", funVar(5))

	// Returning multiple functions from other functions
	//Skip

	fmt.Println(multipleFuncs(10)(20)(30))

	// User defined functions
	x := 100
	answer1 := func(int) Susheel {
		return func(y int) int {
			return x + b + y
		}
	}(x)
	fmt.Println("Answer:: ", answer1(1))
}
