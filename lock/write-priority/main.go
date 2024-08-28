package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type CRWMutex struct {
	mutex      sync.Mutex
	readCount  int32
	writeCount int32
	isWriting  bool
	readCond   *sync.Cond
	writeCond  *sync.Cond
}

func NewCRWMutex() *CRWMutex {
	m := &CRWMutex{}
	m.readCond = sync.NewCond(&m.mutex)
	m.writeCond = sync.NewCond(&m.mutex)
	return m
}

func (m *CRWMutex) LockRead() {
	m.mutex.Lock()
	for atomic.LoadInt32(&m.writeCount) > 0 {
		m.readCond.Wait()
	}
	atomic.AddInt32(&m.readCount, 1)
	m.mutex.Unlock()
}

func (m *CRWMutex) UnlockRead() {
	m.mutex.Lock()
	if atomic.AddInt32(&m.readCount, -1) == 0 && atomic.LoadInt32(&m.writeCount) > 0 {
		m.writeCond.Signal()
	}
	m.mutex.Unlock()
}

func (m *CRWMutex) LockWrite() {
	m.mutex.Lock()
	atomic.AddInt32(&m.writeCount, 1)
	for atomic.LoadInt32(&m.readCount) > 0 || m.isWriting {
		m.writeCond.Wait()
	}
	m.isWriting = true
	m.mutex.Unlock()
}

func (m *CRWMutex) UnlockWrite() {
	m.mutex.Lock()
	if atomic.AddInt32(&m.writeCount, -1) == 0 {
		m.readCond.Broadcast()
	} else {
		m.writeCond.Signal()
	}
	m.isWriting = false
	m.mutex.Unlock()
}
func main() {
	m := NewCRWMutex()
	var cnt int
	var wg sync.WaitGroup
	wg.Add(200)
	m.LockWrite()
	m.UnlockWrite()
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			m.LockRead()
			fmt.Println("read")
			m.UnlockRead()
		}()
	}
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			m.LockWrite()
			fmt.Println("write")
			cnt++
			m.UnlockWrite()
		}()
	}
	wg.Wait()
	fmt.Println(cnt)
}
