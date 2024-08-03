package consistenthash

import (
	"fmt"
	"hash/crc32"
	"log"
	"sort"
)

type Consistent struct {
	hash    func(key []byte) uint32
	replica int            // 每个节点虚拟节点的个数
	nodes   []int          // 节点列表 sorted
	mp      map[int]string // 虚拟节点与真实节点的映射
}

func NewConsistent(replica int, hash func(key []byte) uint32) *Consistent {
	c := &Consistent{
		replica: replica,
		hash:    hash,
		mp:      make(map[int]string),
	}
	if hash == nil {
		c.hash = crc32.ChecksumIEEE
	}
	return c
}

func (c *Consistent) Add(nodes ...string) {
	for _, node := range nodes {
		for i := 0; i < c.replica; i++ {
			// replica-name
			key := int(c.hash([]byte(fmt.Sprintf("%s%d", node, i))))
			c.nodes = append(c.nodes, key)
			c.mp[key] = node // mapping
			log.Printf("add node:%s, key:%d", node, key)

		}
	}
	sort.Ints(c.nodes) // 在环上进行排序
}
func (c *Consistent) Get(key string) string {
	if len(c.nodes) == 0 {
		return ""
	}
	hash := int(c.hash([]byte(key)))
	index := sort.Search(len(c.nodes), func(i int) bool {
		return c.nodes[i] >= hash
	})
	return c.mp[c.nodes[index%len(c.nodes)]]
}
func (c *Consistent) GetAllNodes() []int {
	return c.nodes
}

// DeleteNode 删除节点
func (c *Consistent) DeleteNode(node string) {
	// delete node
	for i := 0; i < c.replica; i++ {
		key := int(c.hash([]byte(fmt.Sprintf("%s%d", node, i))))
		index := sort.Search(len(c.nodes), func(i int) bool {
			return c.nodes[i] >= key
		})
		c.nodes = append(c.nodes[:index], c.nodes[index+1:]...)
	}
	// delete mapping
	for i := 0; i < c.replica; i++ {
		key := int(c.hash([]byte(fmt.Sprintf("%s%d", node, i))))
		delete(c.mp, key)
	}
	sort.Ints(c.nodes)
	log.Printf("delete node:%s", node)
}
