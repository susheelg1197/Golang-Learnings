package main

import "fmt"

type str struct {
	name string
}

func (s *str) String() string {
	return fmt.Sprintf("\n My name is %v..", s.name)
}
func main() {

	sample := &str{name: "Susheel"}
	fmt.Println(sample)
}
