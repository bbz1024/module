package main

import "fmt"

func main() {
	bigHeap := NewBigHeap([]int{8, 9, 1, 2, 3, 6, 5, 7, 4})
	fmt.Println(bigHeap.arr)

	bigHeap.Push(10)
	res := bigHeap.TopK(5)
	fmt.Println(res)

}

type SmallHeap struct {
	arr []int
}

func (b *SmallHeap) Push(v int) {
	b.arr = append(b.arr, v)
	b.adjustUp(len(b.arr) - 1)
}
func (b *SmallHeap) Pop() int {
	if len(b.arr) == 0 {
		return 0
	}
	v := b.arr[0]
	b.arr[0] = b.arr[len(b.arr)-1]
	b.arr = b.arr[:len(b.arr)-1]
	b.adjustDown(0)
	return v
}
func (b *SmallHeap) adjustUp(i int) {
	for {
		parent := (i - 1) / 2
		if parent < 0 || b.arr[parent] < b.arr[i] {
			break
		}
		b.arr[parent], b.arr[i] = b.arr[i], b.arr[parent]
		i = parent
	}

}

func NewBigHeap(arr []int) *SmallHeap {
	b := &SmallHeap{
		arr: arr,
	}
	b.construct()
	return b
}
func (b *SmallHeap) construct() {
	for i := len(b.arr)/2 - 1; i >= 0; i-- {
		b.adjustDown(i)
	}
}
func (b *SmallHeap) Size() int {
	return len(b.arr)
}
func (b *SmallHeap) TopK(k int) []int {
	var res []int
	for i := 0; i < k; i++ {
		res = append(res, b.Pop())
	}
	return res
}
func (b *SmallHeap) adjustDown(i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		min := i
		if left < len(b.arr) && b.arr[left] < b.arr[min] {
			min = left
		}
		if right < len(b.arr) && b.arr[right] < b.arr[min] {
			min = right
		}
		if min == i {
			break
		}
		b.arr[i], b.arr[min] = b.arr[min], b.arr[i]
		i = min
	}
}
