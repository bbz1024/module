package main

// TimeSlotBitmap 用于表示一天中各个时间槽的占用情况
type TimeSlotBitmap uint16 // 假设一天有32个时间槽，使用uint32存储

// SetBit 设置指定位置的位为1，表示时间槽被占用
func (tb *TimeSlotBitmap) SetBit(slot int) {
	if slot < 0 || slot >= 16 {
		panic("Slot out of range")
	}
	*tb |= 1 << slot
}

// ClearBit 清除指定位置的位，表示释放时间槽
func (tb *TimeSlotBitmap) ClearBit(slot int) {
	if slot < 0 || slot >= 16 {
		panic("Slot out of range")
	}
	*tb &= ^(1 << slot)
}

// TestBit 检查指定位置的位是否为1，即时间槽是否被占用
func (tb *TimeSlotBitmap) TestBit(slot int) bool {
	if slot < 0 || slot >= 16 {
		panic("Slot out of range")
	}
	return (*tb>>(slot))&1 == 1
}

// Conflict 检测两个位图是否存在冲突的时间槽
func Conflict(a, b TimeSlotBitmap) bool {
	return a&b != 0
}
func main() {

	/*
		// 假设有两门课程，第一门课占用第2、4时间槽，第二门课占用第3、5时间槽
		course1 := new(TimeSlotBitmap)
		course1.SetBit(2)
		course1.SetBit(4)
		bit := course1.TestBit(9)
		fmt.Println(bit)

	*/
	//c := uint(8)

}
