package main

import "fmt"

func main() {
	bigHeap := NewBigHeap([]int{8, 9, 1, 2, 3, 6, 5, 7, 4})
	fmt.Println(bigHeap.arr)

	bigHeap.Push(10)
	res := bigHeap.TopK(5)
	fmt.Println(res)

}

type BigHeap struct {
	arr []int
}

func (b *BigHeap) Push(v int) {
	b.arr = append(b.arr, v)
	b.adjustUp(len(b.arr) - 1) // 向上调整
}
func (b *BigHeap) Pop() int {
	popV := b.arr[0]
	b.arr[0], b.arr[len(b.arr)-1] = b.arr[len(b.arr)-1], b.arr[0]
	b.arr = b.arr[:len(b.arr)-1]
	b.adjustDown(0)
	return popV
}
func (b *BigHeap) adjustUp(i int) {
	for i > 0 && b.arr[i] > b.arr[(i-1)/2] {
		b.arr[i], b.arr[(i-1)/2] = b.arr[(i-1)/2], b.arr[i]
		i = (i - 1) / 2
	}
}

func NewBigHeap(arr []int) *BigHeap {
	b := &BigHeap{
		arr: arr,
	}
	b.construct()
	return b
}
func (b *BigHeap) construct() {
	/*
		找到最后一个非叶子节点,调整该节点，然后再调整兄弟节点

		左右孩子节点：2i+1 2i+2
		父节点：(i-1)/2
		8,9,1,2,3,6,5,7,4
						8
				9				1
			2		3		6		5
		  7	  4
		最后一个非叶子节点：
		(9-1)/2=4 => 3
		[9 8 6 7 3 1 5 2 4]


						9
				8				6
			7		3		1		5
		  4	  2	  10

	*/
	for i := (len(b.arr) - 1) / 2; i >= 0; i-- {
		b.adjustDown(i)
	}
}
func (b *BigHeap) Size() int {
	return len(b.arr)
}
func (b *BigHeap) TopK(k int) []int {
	var res []int
	for i := 0; i < k; i++ {
		res = append(res, b.Pop())
	}
	return res
}
func (b *BigHeap) adjustDown(i int) {
	left := 2*i + 1
	right := 2*i + 2
	if left >= len(b.arr) {
		return
	}
	var maxIdx = i
	if b.arr[left] > b.arr[maxIdx] {
		maxIdx = left
	}
	if right < len(b.arr) && b.arr[right] > b.arr[maxIdx] {
		maxIdx = right
	}
	if maxIdx != i {
		b.arr[i], b.arr[maxIdx] = b.arr[maxIdx], b.arr[i]
		b.adjustDown(maxIdx)
	}
}
