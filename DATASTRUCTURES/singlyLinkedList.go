package main

import "fmt"

type snode struct {
	data int
	next *snode
}

type singleNode struct {
	length int
	head   *snode
}

func initList() *singleNode {
	return &singleNode{}
}
func (s *singleNode) addFront(ele int) {

	element := &snode{
		data: ele,
	}
	if s.head == nil {
		s.head = element
	} else {
		newEle := s.head
		element.next = newEle
		s.head = element
	}
	s.length++
}

func (s *singleNode) addBack(ele int) {

	element := &snode{
		data: ele,
	}
	if s.head == nil {
		s.head = element
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = element
	}
}
func (s *singleNode) traverse() {

	if s.head == nil {
		fmt.Println("Linked List is empty")
	} else {
		current := s.head
		for current != nil {
			fmt.Println(current.data)
			current = current.next
		}
	}
}

func (s *singleNode) headEle() (int, error) {

	if s.head == nil {
		return 0, fmt.Errorf("Linked list is empty")
	}

	return s.head.data, nil
}

func (s *singleNode) deleteFront() error {
	if s.head == nil {
		return fmt.Errorf("Linked list is empty")
	}

	s.head = s.head.next
	s.length--
	return nil
}
func (s *singleNode) deleteBack() error {
	if s.head == nil {
		return fmt.Errorf("Linked list is empty")
	}
	var prev *snode
	current := s.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		s.head = nil
	}

	return nil
}
func main() {
	node1 := initList()
	node1.addFront(20)
	node1.addFront(30)
	node1.addFront(40)
	node1.addBack(900)
	node1.addBack(1000)
	node1.deleteFront()
	node1.deleteBack()
	node1.deleteBack()
	node1.deleteBack()
	node1.deleteBack()
	node1.deleteFront()
	node1.traverse()
}
