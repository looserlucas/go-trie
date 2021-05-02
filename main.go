package main

import "fmt"

type TrieNode struct {
	Children    []*TrieNode
	IsEndOfWord bool
}

const mASize int = 25

func NewNode() (n *TrieNode) {
	return &TrieNode{
		Children:    make([]*TrieNode, mASize),
		IsEndOfWord: false,
	}
}

func Insert(root *TrieNode, key string) {
	currentNode := root
	for i := 0; i < len(key); i++ {
		index := int(key[i]) - 97
		if currentNode.Children[index] == nil {
			currentNode.Children[index] = NewNode()
		}
		currentNode = currentNode.Children[index]
	}

	currentNode.IsEndOfWord = true
}

func Search(root *TrieNode, key string) (ok bool) {
	currentNode := root
	i := 0
	for i < len(key) {
		index := int(key[i]) - 97
		if currentNode.Children[index] == nil {
			return false
		} else {
			currentNode = currentNode.Children[index]
		}
	}

	return true
}

//func PrintTrie(root *TrieNode)

func main() {
	t1 := "asdasd"
	t2 := "asdasfasf"
	t3 := "asdwhrjkq"
	p := []string{"asd", "fd", "as", "da", "kq", "wr"}

	root := NewNode()
	Insert(root, t1)
	Insert(root, t2)
	Insert(root, t3)

	for i := 0; i < len(p); i++ {
		if Search(root, p[i]) {
			fmt.Printf("there is %s in the trie \n", p[i])
		} else {
			fmt.Printf("there isn't %s in the trie \n", p[i])
		}
	}
}
