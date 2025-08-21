package tree

type RadixNode struct {
	edgeLabel string
	children map[byte]*RadixNode
	isEndOfWord bool
}

type RadixTree struct {
	root *RadixNode
}

func NewRadixNode() *RadixNode {
	return &RadixNode{
		children: make(map[byte]*RadixNode),
	}
}

func NewRadixTree() *RadixTree {
	return &RadixTree{
		root: NewRadixNode(),
	}
}

func (radix *RadixTree) Insert(word string) {
	node := radix.root
	i := 0
	
	for i < len(word) {
		char := word[i]
		child, found := node.children[char]
		
		if !found {
			newNode := &RadixNode{
				edgeLabel: word[i:],
				children: make(map[byte]*RadixNode),
				isEndOfWord: true,
			}
			node.children[char] = newNode
			return
		}
		
		j := 0
		for j < len(child.edgeLabel) && i < len(word) && child.edgeLabel[j] == word[i] {
			j++
			i++
		}
		
		if j == len(child.edgeLabel) {
			node = child
		} else {
			splitNode := &RadixNode{
				edgeLabel: child.edgeLabel[:j],
				children: make(map[byte]*RadixNode),
				isEndOfWord: false,
			}
			child.edgeLabel = child.edgeLabel[j:]
			splitNode.children[child.edgeLabel[0]] = child
			
			if i < len(word) {
				nn := &RadixNode{
					edgeLabel: word[i:],
					children: make(map[byte]*RadixNode),
					isEndOfWord: true,
				}
				splitNode.children[nn.edgeLabel[0]] = nn
			} else {
				splitNode.isEndOfWord = true
			}
			
			node.children[splitNode.edgeLabel[0]] = splitNode
			return
		}
	}
}
