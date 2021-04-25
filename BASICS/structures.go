/*
A structure is a heterogenous collection of elements. It is compare with the class
in OOP. It can be termed as lightweight class that not supports inheritence
but supports composition.

*/
package main

import "fmt"

type Salary struct {
	HRA         float64
	Da          float64
	isInsurance bool
}
type Address struct { // using the derived "type" to create the type Address
	name    string
	street  string
	city    string
	state   string
	Pincode int
	salary  Salary
}

func main() {
	var a Address // all fields are set to their zero value
	// a = Address{  // Too few values in struct literal
	// 	"SSG", "S1", "jalna", "MAha", 1212, a.salary.Da: 101.0, 20.02, false, //missing ',' before newline in
	// 	//  composite literal
	// }
	// The above syntax is perfectly valid. But when creating a struct without
	//  declaring the field names, you need to provide all field values in the
	// order of their appearance in the struct type

	b := Address{name: "Susheel", street: "Abc", city: ""}
	fmt.Println(a, b)

	//Anonymous struct
	/*
		Used to define a struct type without specifying the derived struct "type"
		keyword

	*/
	emp := struct {
		fname, lname string
		rollNumber   int
		isDeleted    bool
	}{
		fname: "susheel", lname: "Gounder", rollNumber: 9,
	}
	fmt.Println("Anonyomos struct print:: ", emp)
	/*
		Creating a derived type from the built-in struct type gives us the
		flexibility to reuse it without having to write complex syntax again and again.
	*/

	// Pointer to a struct

	//Syntax
	ptr := &Address{name: "SSSSS"}
	// to print the value we have to use dereferencing operatior
	fmt.Println("Pointer value:: ", (*ptr).name) // this will help the compiler to
	// understand between (*ptr).name and (*ptr.name)
	// Under the hood struct can we used without dereferencing even
	fmt.Println("UndertheHoood type:: ", ptr.name)

	//Anonymous fields

	/*
		we can define only the data type even while defining struct type
		Go will use the data type declaration as the field names
	*/
	// Example

	field1 := struct {
		int
		float64
		string
	}{1, 1.0, "abcd"}
	fmt.Println(field1.string)

	//Nested struct
	/*
		A struct field can have any data type, hence it can also be nested within
		another struct
	*/
}
