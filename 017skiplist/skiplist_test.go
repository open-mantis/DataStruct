package _17skiplist

import (
	"fmt"
	"testing"
)

func TestSkipList(t *testing.T) {
	sl := NewSkipList()

	fmt.Println(sl.Insert("leo", 95))
	fmt.Println(sl.Insert("leo1", 951))
	fmt.Println(sl.Insert("leo1", 95))
	sl.Print()
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[1])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl.head.forwards[1].forwards[1])
	t.Log(sl)
	t.Log("-----------------------------")
	//
	sl.Insert("jack", 88)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0])
	t.Log(sl)
	t.Log("-----------------------------")

	sl.Insert("lily", 100)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0].forwards[0])
	t.Log(sl)
	t.Log("-----------------------------")

	t.Log(sl.Find("jack", 88))
	t.Log("-----------------------------")
	sl.Print()

	sl.Delete("leo", 95)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl.head.forwards[0].forwards[0].forwards[0])
	t.Log(sl)

	t.Log("-----------------------------")
}
