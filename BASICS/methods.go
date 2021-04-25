/*
A go method is a function having a limited scope and attached to a specific type.
This type is called as reciever type. Method provide a way to add behaviour to
these user defined types.
func (recv recv_type) methodName()
A method is a function with a receiver. A method declaration binds an identifier,
the method name, to a method,
and associates the method with the receiver's base type.

If you want to change the state of the receiver in a method, manipulating the
value of it, use a pointer receiver. It’s not possible with a value receiver,
which copies by value. Any modification to a value receiver is local to that copy.
If you don’t need to manipulate the receiver value, use a value receiver.
*/
package main

import "fmt"

type number int
type decimal float64

// Value reciever - the method will be acting against the copy of the
// value
func (m number) tenTimes() int { //must declare a single non-variadic parameter,
	return int(m * 10)
}

// Method overloading will be based on the reciever type
func (m decimal) tenTimes() int {
	return int(m * 10)
}

// Pointer reciever
func (m *number) tenTimes2() int {
	return int((*m) * 10)
}
func main() {
	res := number(10) // Create instance
	res2 := decimal(10.1)
	res3 := number(20)
	fmt.Println(res.tenTimes(), "and.. ", res2.tenTimes(), "pointer.. ", res3.tenTimes2(), res3)

}
