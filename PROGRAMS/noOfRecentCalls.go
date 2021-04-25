package main

import "fmt"

// Slice as queue implementation

type pings struct {
	queue []int
}

func Constructor() pings {
	return pings{[]int{}}
}
func (p *pings) PingCall(t int) int {
	p.queue = append(p.queue, t)
	for i, val := range p.queue {
		if val >= t-2000 {
			p.queue = p.queue[i:] // pop the left part
			break
		}
	}
	fmt.Println(p.queue)

	return len(p.queue)
}

func main() {
	obj := Constructor()
	list1 := [][]int{{}, {1}, {100}, {3001}, {3002}}

	for _, list := range list1 {
		for _, t := range list {
			param1 := obj.PingCall(t)
			fmt.Println(param1)
		}
	}

}
