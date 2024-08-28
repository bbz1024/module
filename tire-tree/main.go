package main

import "fmt"

type TrieTree struct {
	root TrieNode
}

type TrieNode struct {
	next  map[rune]*TrieNode
	isEnd bool
}

func (t *TrieNode) insert(word string) {
	this := t
	for _, v := range word {
		if this.next[v] == nil {
			node := new(TrieNode)
			node.next = make(map[rune]*TrieNode)
			this.next[v] = node
		}
		this = this.next[v]
	}
	this.isEnd = true
}
func (t *TrieNode) search(word string) bool {
	this := t
	for _, v := range word {
		if this.next[v] == nil {
			return false
		}
		this = this.next[v]
	}
	return this.isEnd
}
func (t *TrieNode) startsWith(prefix string) bool {
	this := t
	for _, v := range prefix {
		if this.next[v] == nil {
			return false
		}
		this = this.next[v]
	}
	return true
}
func (t *TrieTree) Insert(word string) {
	t.root.insert(word)
}
func (t *TrieTree) Search(word string) bool {
	return t.root.search(word)
}
func (t *TrieTree) StartsWith(prefix string) bool {
	return t.root.startsWith(prefix)
}
func NewTrieTree() *TrieTree {
	return &TrieTree{
		root: TrieNode{
			next:  make(map[rune]*TrieNode),
			isEnd: false,
		},
	}
}

func main() {
	t := NewTrieTree()
	t.Insert("abc")
	t.Insert("abd")
	t.Insert("abe")
	t.Insert("bc")
	t.Insert("bd")
	t.Insert("aaa")
	fmt.Println(t.Search("abc"))    // true
	fmt.Println(t.StartsWith("a"))  // true
	fmt.Println(t.StartsWith("ba")) // false
	fmt.Println(t.StartsWith("b"))  // true
}
