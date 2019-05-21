package _03queue

import "testing"

func TestListQueue_EnQueue(t *testing.T) {
	q := NewLinkedListQueue()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	t.Log(q)
	t.Log(q.Print())
}

func TestListQueue_DeQueue(t *testing.T) {
	q := NewLinkedListQueue()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	t.Log(q)
	t.Log(q.Print())
	var v interface{}
	v = q.DeQueue()
	t.Log(q)
	v = q.DeQueue()
	t.Log(v)
	t.Log(q)
	v = q.DeQueue()
	t.Log(q)
	t.Log(v)
	v = q.DeQueue()
	t.Log(q)
	t.Log(v)
	v = q.DeQueue()
	t.Log(q)
	t.Log(v)
}
