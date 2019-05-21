package _03queue

import (
	"fmt"
)

type CircularQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

func NewCircularQueue(n int) *CircularQueue {
	if n == 0 {
		return nil
	}
	return &CircularQueue{make([]interface{}, n), n, 0, 0}
}

func (this *CircularQueue) IsEmpty() bool {
	if this.head == this.tail {
		return true
	}
	return false
}

func (this *CircularQueue) IsFull() bool {
	return this.head == (this.tail+1)%this.capacity
}

func (this *CircularQueue) EnQueue(v interface{}) bool {
	if this.IsFull() {
		return false
	}
	this.q[this.tail] = v
	this.tail = (this.tail + 1) % this.capacity
	return false
}

func (this *CircularQueue) DeQueue() interface{} {
	if this.IsEmpty() {
		return false
	}
	val := this.q[this.head]
	this.head = (this.head + 1) % this.capacity
	return val
}

func (this *CircularQueue) String() string {
	if this.IsEmpty() {
		return "this is empty queue"
	}
	s := "head"
	tmp := this.head
	fmt.Println(this.tail)
	for tmp != this.tail {
		s += "<-" + fmt.Sprintf("%v", this.q[tmp])
		tmp = (tmp + 1) % this.capacity
	}
	s += "<-tail"
	return s
}
