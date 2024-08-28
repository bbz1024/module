package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// ReadPriorityLock 读写锁，读优先
type ReadPriorityLock struct {
	mutex   sync.Mutex
	readCnt atomic.Int32
}

func (r *ReadPriorityLock) Lock() {
	r.mutex.Lock()
}

func (r *ReadPriorityLock) Unlock() {
	r.mutex.Unlock()
}
func (r *ReadPriorityLock) RLock() {
	if r.readCnt.Load() == 0 {

		r.mutex.Lock()
	}
	r.readCnt.Add(1)

}
func (r *ReadPriorityLock) RUnlock() {
	r.readCnt.Add(-1)
	if r.readCnt.Load() == 0 {
		r.mutex.Unlock()
	}
}
func main() {
	cnt := 0
	lock := &ReadPriorityLock{}
	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(time.Second)
			lock.RLock()
			fmt.Println(cnt)
			lock.RUnlock()
		}()
	}
	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(time.Second)
			lock.Lock()
			cnt++
			lock.Unlock()
		}()
	}

	time.Sleep(time.Second * 5)
	fmt.Println(cnt)
}
