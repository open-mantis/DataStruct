package _01_linkedlist

import "fmt"

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head   *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func (this *ListNode) GetNext() *ListNode {
	return this.next
}

func (this *ListNode) GetValue() interface{} {
	return this.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

//在某个节点后面插入节点
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}
	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	this.length++
	return true
}

//在某个节点前面插入节点
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if p == nil || p == this.head {
		return false
	}

	//获取前继节点
	current := this.head.next
	prefix := this.head
	for current != nil {
		if current == p {
			break
		}
		prefix = current
		current = current.next
	}

	if current == nil {
		return false
	}

	newNode := NewListNode(v)
	prefix.next = newNode
	newNode.next = current
	this.length++
	return true
}

//在链表头部插入节点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertBefore(this.head, v)
}

//在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	current := this.head
	if current.next != nil {
		current = current.next
	}
	return this.InsertAfter(current, v)
}

//通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.length {
		return nil
	}

	current := this.head
	var i uint = 0
	for ; i < index; i++ {
		current = current.next
	}
	return current
}

//删除传入的节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}

	//获取前继节点
	current := this.head.next
	prefix := this.head
	for current != nil {
		if current == p {
			break
		}
		prefix = current
		current = current.next
	}

	if current == nil {
		return false
	}

	prefix.next = current.next
	p = nil
	this.length--
	return true
}

//打印链表
func (this *LinkedList) Print() {
	cur := this.head
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
