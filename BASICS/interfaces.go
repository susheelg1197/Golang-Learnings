/*
An interface is an abstract type.
It is a set of method signatures.
A value of interface type can hold any value that implements thos methods
A type implements an interface by implementing its methods.

There is no explicit declaration of intent, no "implements" keyword.
Under the hood, interface values can be thought of as a tuple of a value and a
concrete type:

(value, type)
An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its
underlying type.

If the concrete value inside the interface itself is nil, the method will be
called with a nil receiver.

In some languages this would trigger a null pointer exception, but in Go it
is common to write methods that gracefully handle being called with a nil
receiver (as with the method M in this example.)

Note that an interface value that holds a nil concrete value is itself non-nil.

A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no
type inside the interface tuple to indicate which concrete method to call.


*/
package main

import (
	"fmt"
	"math"
)

type impl interface {
	abs() float64
}

type Vertex1 struct {
	x float64
	y float64
}

func (v *Vertex1) abs() float64 {
	return math.Abs(v.x) + math.Abs(v.y)
}

func main() {
	var im impl
	v := &Vertex1{x: 1.1, y: 2.1}

	describe(im)
	// im.abs()// invalid memory address or pointer dereference
	im = v
	describe(im) // A tuple of Type and value

	// Value can be nil
	v1 := &Vertex1{}
	im = v1
	describe(im)
	fmt.Println(im.abs())

	// An empty interface
	// It is an interface with no method signatures
	// it takes any values that implements no methods
	// It is used to hold any type of values best example is fmt.println
	var empty interface{}
	empty = 42
	describe(empty)
	empty = "Hi"
	describe(empty)

	//Type assertion is used to provide access to interface value's
	//underlying concrete value

	variable, ok := empty.(int)
	fmt.Println("Value: ", variable, "IsPresent", ok)

	// type switch
	typeCheck(40)
	typeCheck(50.55)

}
func describe(im interface{}) {

	fmt.Printf("\n%T & %v\n", im, im)
}
func typeCheck(im interface{}) {
	switch v := im.(type) { // type keyword
	case int:
		fmt.Println("Integer: ", v)
	case float64:
		fmt.Println("Float: ", v)
	}
}
