package _03queue

import "fmt"

type ListNode struct {
	val  interface{}
	next *ListNode
}

type LinkedListQueue struct {
	head   *ListNode
	tail   *ListNode
	length int
}

//不存在容量上限
func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{nil, nil, 0}
}

//基于链表实现是头出尾入

func (this *LinkedListQueue) EnQueue(v interface{}) {
	node := &ListNode{v, nil}
	//判断队列为空
	if this.head == nil {
		this.head = node
		this.tail = node
	}
	//插入新元素
	this.tail.next = node

	//移动尾部指针
	this.tail = this.tail.next

	this.length++

}

func (this *LinkedListQueue) DeQueue() interface{} {
	if this.head == nil {
		return nil
	}
	//取出元素
	val := this.head.val

	//移动头指针
	this.head = this.head.next
	return val
}

func (this *LinkedListQueue) Print() string {
	if this.head == nil {
		return "empty queue"
	}
	result := "head"
	for cur := this.head; cur != nil; cur = cur.next {
		fmt.Println("current", cur.val)
		result += fmt.Sprintf("<-%+v", cur.val)
	}
	result += "<-tail"
	return result

}
