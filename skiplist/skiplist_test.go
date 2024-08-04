package skiplist

import (
	"math/rand"
	"testing"
)

func TestSkipList(t *testing.T) {
	skipList := NewSkipList(func() bool {
		return rand.Float32() > 0.5
	}, 5)
	skipList.Add(1, 1)
	skipList.Add(2, 2)
	skipList.Add(3, 3)
	skipList.Add(4, 4)
	skipList.Add(4, 8)
	if skipList.Search(4).Val.(int) == 4 {
		t.Error("search error")
	}
	skipList.Delete(4)
	if skipList.Search(4) != nil {
		t.Error("search error")
	}
	skipList.Add(4, 8)
	if skipList.Search(4).Val.(int) != 8 {
		t.Error("search error")
	}
}
