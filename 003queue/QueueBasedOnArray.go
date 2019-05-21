package _03queue

import "fmt"

type ArrayQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

//存在容量上限
func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, capacity), capacity, 0, 0}
}

func (this *ArrayQueue) EnQueue(val interface{}) bool {
	//	判断容量上限
	if this.tail == this.capacity {
		return false
	}
	this.q[this.tail] = val
	this.tail++
	return true
}

func (this *ArrayQueue) DeQueue() interface{} {
	if this.head == this.tail {
		return nil
	}
	val := this.q[this.head]
	this.head++
	return val
}

func (this *ArrayQueue) Print() string {
	s := "head"
	for _, val := range this.q {
		s += "<-" + fmt.Sprintf("%v", val)
	}
	s += "<-tail"
	return s
}
