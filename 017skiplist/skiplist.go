package _17skiplist

import (
	"fmt"
	"math"
	"math/rand"
)

/*
我理解的还不是redis zset层面的数据结构时间，不能满足插入相关val 不同score数据问题
*/

//设置索引层数
const MaxLevel = 6

//跳表节点结构体
type skipListNode struct {
	//实际存储数据
	val interface{}

	//用于排序分值
	score int

	//层高
	level int

	//每层前进指针
	forwards []*skipListNode
}

//新建跳表节点
func newSkipListNode(val interface{}, sore, level int) *skipListNode {
	return &skipListNode{val: val, score: sore, level: level, forwards: make([]*skipListNode, level, level)}
}

//跳表结构体
type SkipList struct {
	head   *skipListNode
	level  int
	length int
}

func NewSkipList() *SkipList {
	head := newSkipListNode(0, math.MinInt32, MaxLevel)
	return &SkipList{head, 1, 0}
}

//获取跳表层级
func (sl *SkipList) Level() int {
	return sl.level
}

//获取跳表长度
func (sl *SkipList) Length() int {
	return sl.length
}

func (sl *SkipList) String() string {
	return fmt.Sprintf("level:%+v, length:%+v", sl.level, sl.length)
}

func (sl *SkipList) Insert(v interface{}, score int) int {
	if v == nil {
		return 1
	}
	//查找插入位置
	cur := sl.head

	//记录每次的路径
	update := [MaxLevel]*skipListNode{}

	i := MaxLevel - 1

	//从顶层查找
	for ; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].val == v {
				return 2
			}
			if cur.forwards[i].score > score {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		if nil == cur.forwards[i] {
			update[i] = cur

		}
	}

	//通过随机算法获取该节点层数
	level := 1
	for i := 1; i < MaxLevel; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	//创建新跳跃表节点
	newNode := newSkipListNode(v, score, level)

	//原有节点连接
	for i := 0; i <= level-1; i++ {
		next := update[i].forwards[i]
		update[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	//如果当前节点的层数大于之前跳表的层数
	//更新当前跳表层数
	if level > sl.level {
		sl.level = level
	}

	//更新跳表长度
	sl.length++

	return 0
}

func (sl *SkipList) Find(v interface{}, score int) *skipListNode {
	if v == nil || sl.length == 0 {
		return nil
	}
	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for cur.forwards[i] == nil {
			if cur.forwards[i].score == score && cur.forwards[i].val == v {
				return cur.forwards[i]

			} else if cur.forwards[i].score > score {
				break
			}
			cur = cur.forwards[i]
		}
	}
	return nil
}

func (sl *SkipList) Delete(v interface{}, score int) int {
	if v == nil {
		return 1
	}

	//	查找前驱节点
	cur := sl.head

	//记录所有的前驱节点
	update := [MaxLevel]*skipListNode{}
	for i := sl.level - 1; i >= 0; i-- {
		update[i] = sl.head
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].val == v {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
	}

	//开始删除节点
	cur = update[0].forwards[0]
	for i := cur.level - 1; i >= 0; i-- {
		if update[i] == sl.head && cur.forwards[i] == nil {
			sl.level = i
		}

		if nil != update[i].forwards[i] {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
	}

	sl.length--

	return 0
}

func (sl *SkipList) Print() {
	curr := sl.head
	fmt.Println("begin------------")
	for curr.forwards[0] != nil {
		fmt.Println("val", curr.forwards[0].val, "score", curr.score)
		curr = curr.forwards[0]
	}
	fmt.Println("end--------------")
}
