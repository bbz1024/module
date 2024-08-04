package skiplist

import (
	"math"
)

/*
go实现SkipList
*/

type SkipNode struct {
	Key         int
	Val         interface{}
	Right, Down *SkipNode
	IsDelete    bool
}

func NewSkipNode(key int, val interface{}) *SkipNode {
	return &SkipNode{Key: key, Val: val}
}

type SkipList struct {
	Head     *SkipNode
	Level    int         // 当前跳表索引层数
	Random   func() bool // 用于判断是否建立索引
	MaxLevel int         // 最大层，MaxLevel越大所建造的索引越大，占用的空间也越大
}

func NewSkipList(random func() bool, maxLevel int) *SkipList {
	return &SkipList{
		Head:     &SkipNode{},
		Random:   random,
		MaxLevel: maxLevel,
	}
}
func (s *SkipList) Add(k int, v interface{}) {
	// if exist to update
	node := s.search(k)
	if node != nil {
		node.IsDelete = false
		node.Val = v
		return
	}
	// set
	var stack []*SkipNode
	dummy := SkipNode{
		Right: s.Head,
	}
	cur := dummy.Right // head
	// find the insert position
	for cur != nil {
		if cur.Right == nil { //右边没有向下操作
			stack = append(stack, cur)
			cur = cur.Down // 要么就是head的下，要么就是前一个节点的下
		} else if cur.Right.Key > k { // 需要下降寻找
			stack = append(stack, cur)
			cur = cur.Down
		} else {
			// 向右
			cur = cur.Right
		}
	}
	// 找到插入位置了
	var level int
	var downNode *SkipNode
	for len(stack) != 0 {
		val := stack[len(stack)-1] //插入节点的左节点
		stack = stack[:len(stack)-1]
		node := NewSkipNode(k, v)
		node.Down = downNode
		downNode = node
		node.Right = val.Right
		val.Right = node
		// 如果大于了最大索引就不建立，是否需要向上建立索引

		if level < s.MaxLevel && !s.Random() {
			return
		}

		level++
		//创建头，因为到达了最顶level但是没有超过MaxLevel
		if level > s.Level {
			s.Level = level
			//	x需要创建一个新的节点
			node := NewSkipNode(math.MinInt32, math.MinInt32)
			node.Down = s.Head
			s.Head = node
			stack = append(stack, node)
		}
	}
}
func (s *SkipList) search(key int) *SkipNode {

	dummy := SkipNode{Right: s.Head}
	cur := dummy.Right
	for cur != nil {
		if cur.Right == nil { // 右侧没有向下找
			cur = cur.Down
		} else if cur.Right.Key > key { // 需要向下找
			cur = cur.Down
		} else if cur.Right.Key == key { // 找到了
			return cur.Right
		} else { // 右侧比较小，向右
			cur = cur.Right
		}
	}
	return nil
}
func (s *SkipList) Search(key int) *SkipNode {
	if node := s.search(key); node != nil && !node.IsDelete {
		return node
	}
	return nil
}
func (s *SkipList) Delete(key int) (node *SkipNode) {
	dummy := SkipNode{Right: s.Head}
	cur := dummy.Right
	for cur != nil {
		if cur.Right == nil {
			cur = cur.Down
		} else if cur.Right.Key > key {
			cur = cur.Down
		} else if cur.Right.Key == key {
			node = cur.Right
			//cur.Right = cur.Right.Right // 清理节点
			cur.Right.IsDelete = true
			cur = cur.Down
		} else {
			cur = cur.Right
		}
	}
	return
}
