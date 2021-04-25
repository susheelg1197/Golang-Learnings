package main

import "fmt"

type tnode struct {
	data  int
	lnode *tnode
	rnode *tnode
}
type tree struct {
	root *tnode
}

func treeinit() *tree {
	return &tree{}
}
func (t *tree) insert(ele int) {
	t.insertRec(t.root, ele)

}
func (t *tree) insertRec(node *tnode, ele int) *tnode {
	if t.root == nil {
		t.root = &tnode{
			data: ele,
		}
		return t.root
	}
	if node == nil {
		return &tnode{data: ele}
	}
	if ele > node.data {
		node.rnode = t.insertRec(node.rnode, ele)
	}
	if ele <= node.data {
		node.lnode = t.insertRec(node.lnode, ele)
	}

	return node

}

func (t *tree) inorder() {
	if t.root == nil {
		fmt.Println("Tree is empty")
	} else {

		t.inorderRec(t.root)
	}
}
func (t *tree) inorderRec(node *tnode) {

	if node != nil {
		t.inorderRec(node.lnode)
		fmt.Println(node.data)
		t.inorderRec(node.rnode)

	}
}
func main() {
	arr := []int{100, 90, 40, 70, 60, 55, 99, 15}
	t := treeinit()
	for _, val := range arr {
		t.insert(val)
	}
	t.inorder()
}
