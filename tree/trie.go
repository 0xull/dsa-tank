package tree

type TrieNode struct {
	children map[rune]*TrieNode
	isEndOfWord bool
}

type Trie struct {
	root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEndOfWord: false,
	}
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (trie *Trie) Insert(word string) {
	current := trie.root
	
	for _, char := range word {
		if _, found := current.children[char]; !found {
			current.children[char] = NewTrieNode()
		}
		current = current.children[char]
	}
	
	current.isEndOfWord = true
}

func (trie *Trie) SearchWord(word string) bool {
	curr := trie.root
	
	for _, char := range word {
		if _, found := curr.children[char]; !found {
			return false
		}
		curr = curr.children[char]
	}
	return curr.isEndOfWord
}

func (trie *Trie) StartsWith(prefix string) bool {
	curr := trie.root
	
	for _, char := range prefix {
		if _, found := curr.children[char]; !found {
			return false
		}
		curr = curr.children[char]
	}
	return true
}
