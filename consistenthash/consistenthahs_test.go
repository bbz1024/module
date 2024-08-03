package consistenthash_test

import (
	"fmt"
	"module/consistenthash"
	"strconv"
	"testing"
)

func TestAddNode(t *testing.T) {
	consistent := consistenthash.NewConsistent(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})
	nodes := []string{"1", "2", "3", "4"}
	consistent.Add(nodes...)
	fmt.Println(consistent.GetAllNodes())
	/*
		  10 11 12
		42	      20
		41		  21
		40        22
		  32 31 30
	*/
}
func TestGet(t *testing.T) {
	consistent := consistenthash.NewConsistent(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})
	nodes := []string{"1", "2", "3", "4"}
	consistent.Add(nodes...)
	for i := 0; i < 50; i++ {
		key2 := strconv.Itoa(i)
		fmt.Println(key2, consistent.Get(key2))
	}
}
func TestHashAdd(t *testing.T) {
	hash := consistenthash.NewConsistent(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	// Given the above hash function, this will give replicas with "hashes":
	// 10，11，12，20，21，22，30，31，32
	hash.Add("1", "2", "3")
	mapCase := map[string]string{
		"10": "1",
		"11": "1",
		"12": "1",
		"20": "2",
		"21": "2",
		"22": "2",
		"30": "3",
		"31": "3",
		"32": "3",
	}
	for k, v := range mapCase {
		if v != hash.Get(k) {
			t.Errorf("Get(%s) failed: expected %s, got %s", k, v, hash.Get(k))
		}
	}
	//	add node
	hash.Add("4")
	mapCase = map[string]string{
		"10": "1",
		"11": "1",
		"12": "1",
		"20": "2",
		"21": "2",
		"22": "2",
		"30": "3",
		"31": "3",
		"32": "3",
		"40": "4",
	}
	for k, v := range mapCase {
		if v != hash.Get(k) {
			t.Errorf("Get(%s) failed: expected %s, got %s", k, v, hash.Get(k))
		}
	}
	fmt.Println(hash.GetAllNodes())

}
func TestHashing(t *testing.T) {
	hash := consistenthash.NewConsistent(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	// Given the above hash function, this will give replicas with "hashes":
	// 10，11，12，20，21，22，30，31，32
	hash.Add("1", "2", "3")
	mapCase := map[string]string{
		"10": "1",
		"11": "1",
		"12": "1",
		"20": "2",
		"21": "2",
		"22": "2",
		"30": "3",
		"31": "3",
		"32": "3",
	}
	for k, v := range mapCase {
		if v != hash.Get(k) {
			t.Errorf("Get(%s) failed: expected %s, got %s", k, v, hash.Get(k))
		}
	}

}
