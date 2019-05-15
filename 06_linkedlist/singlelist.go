package _6_linkedlist

type Linklist struct {
	value interface{}
	next  *Linklist
}

func NewListNode(val interface{}) *Linklist {
	return &Linklist{val, nil}
}

//获取
