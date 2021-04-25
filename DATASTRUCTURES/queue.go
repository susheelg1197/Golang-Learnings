package main

import "fmt"

type queue struct {
	list []int
}

func (q *queue) enqueue(elem int) {

	q.list = append(q.list, elem)
}
func (q *queue) dequeue() error {

	if len(q.list) == 0 {
		return fmt.Errorf("Queue is Empty")
	}
	q.list = q.list[1:]
	return nil
}

func (q *queue) front() (int, error) {

	if len(q.list) == 0 {
		return 0, fmt.Errorf("Queue is Empty")
	}
	return q.list[0], nil
}
func main() {

	queueObj := &queue{
		list: make([]int, 0),
	}
	queueObj.enqueue(20)
	queueObj.enqueue(200)
	queueObj.enqueue(2099)
	fmt.Println(queueObj.list)
	queueObj.dequeue()
	fmt.Println(queueObj.list)

}
