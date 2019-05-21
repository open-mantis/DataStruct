package _02stack

import "fmt"

type node struct {
	val  interface{}
	next *node
}

type LinkedListStack struct {
	topNode *node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (this *LinkedListStack) Push(v interface{}) {
	this.topNode = &node{next: this.topNode, val: v}
}

func (this *LinkedListStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	val := this.topNode.val
	this.topNode = this.topNode.next
	return val
}

func (this *LinkedListStack) IsEmpty() bool {
	if this.topNode == nil {
		return true
	}
	return false
}

func (this *LinkedListStack) Top() interface{} {
	if this.topNode == nil {
		return nil
	}
	return this.topNode.val
}

func (this *LinkedListStack) Flush() {
	this.topNode = nil
}
func (this *LinkedListStack) Print() {
	if this.IsEmpty() {
		fmt.Print("this stack empty")
	} else {
		current := this.topNode
		for current != nil {
			fmt.Print(current.val)
			current = current.next
		}
	}
}
