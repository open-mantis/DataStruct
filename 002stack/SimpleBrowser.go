package _02stack

import "fmt"

type Browser struct {
	forwardStack Stack
	backStack    Stack
}

func NewBrowser() *Browser {
	return &Browser{
		forwardStack: NewLinkedListStack(),
		backStack:    NewLinkedListStack(),
	}
}

func (this *Browser) CanForward() bool {
	if this.forwardStack.IsEmpty() {
		return false
	}
	return true
}

func (this *Browser) CanBack() bool {
	if this.backStack.IsEmpty() {
		return false
	}
	return true
}

func (this *Browser) Open(addr string) {
	fmt.Printf("Open new addr %+v\n", addr)
	this.forwardStack.Flush()
}

func (this *Browser) Forward() {
	if !this.CanForward() {
		return
	}
	val := this.forwardStack.Pop()
	this.backStack.Push(val)

	fmt.Printf("forward to %+v\n", val)
}

func (this *Browser) Back() {
	if !this.CanBack() {
		return
	}
	val := this.backStack.Pop()
	this.forwardStack.Push(val)

	fmt.Printf("back to %+v\n", val)
}
