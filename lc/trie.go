package main

import "fmt"

// AlphabeSize is the number of possible characters in the trie
const AlphabetSize = 26

// Node represents each node in the trie
type trieNode struct {
	children [AlphabetSize]*trieNode
	isEnd    bool
}

// Trie represent and as a pointer to its root node
type Trie struct {
	root *trieNode
}

// InitTrie will create new Trie
func InitTrie() *Trie {
	result := &Trie{root: &trieNode{}}
	return result
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	wordLen := len(w)
	currentNode := t.root
	for i := 0; i < wordLen; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &trieNode{}
		}

		currentNode = currentNode.children[charIndex]
	}

	currentNode.isEnd = true
}

// Search takes in a word and return true if that words is included in the trie
func (t *Trie) Search(w string) bool {
	wordLen := len(w)
	currentNode := t.root
	for i := 0; i < wordLen; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}

		currentNode = currentNode.children[charIndex]
	}

	if currentNode.isEnd {
		return true
	}

	return false
}

func triemain() {
	testTrie := InitTrie()

	toAdd := []string{"sishir", "sish", "sishi", "ishi", "hir"}

	for _, v := range toAdd {
		testTrie.Insert(v)
	}

	isTrue := testTrie.Search("ishi")

	fmt.Println(isTrue)

}
