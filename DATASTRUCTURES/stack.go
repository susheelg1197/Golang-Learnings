package main

import "fmt"

type stack struct {
	list []int
}

func (s *stack) push(elem int) {
	s.list = append(s.list, elem)
}
func (s *stack) pop() error {

	if len(s.list) == 0 {
		return fmt.Errorf("Stack is empty")
	}
	s.list = s.list[:len(s.list)-1]
	return nil
}

func (s *stack) Head() (int, error) {
	if len(s.list) == 0 {
		return 0, fmt.Errorf("Stack is empty")
	}

	return s.list[len(s.list)-1], nil
}
func (s *stack) size() int {

	return len(s.list)
}
func (s *stack) isEmpty() bool {

	return len(s.list) == 0
}
func main() {
	stackObj := &stack{
		list: make([]int, 0),
	}
	stackObj.push(1)
	fmt.Println(stackObj.list)
	stackObj.push(2)
	stackObj.push(3)
	fmt.Println(stackObj.list)
	stackObj.pop()
	fmt.Println(stackObj.list)
	ele, _ := stackObj.Head()
	fmt.Println("Head: ", ele)

}
