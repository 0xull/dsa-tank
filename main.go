package main

import (
	"fmt"

	"github.com/IkehAkinyemi/dsa-tank/ds/list/linkedlist"
	stack "github.com/IkehAkinyemi/dsa-tank/ds/stack/example"
)

type Process struct {
	PID int
}

func main() {
	sll := &linkedlist.SinglyLinkedList[Process]{
		Head: &linkedlist.Node[Process]{Data: Process{1}, Next: &linkedlist.Node[Process]{Data: Process{2}, Next: nil}},
	}

	sll.Display()
	fmt.Println(sll.Search(Process{2}))
	sll.Prepend(Process{0})
	sll.Append(Process{3})
	sll.InsertAfter(Process{0}, Process{4})
	sll.InsertBefore(Process{3}, Process{10})
	fmt.Printf("%v\n", sll.DeleteHead())
	fmt.Printf("%v\n", sll.DeleteTail())
	fmt.Printf("%v\n", sll.Delete(Process{2}))
	sll.Display()
	
	cll := &linkedlist.CircularLinkedList[Process]{}
	cll.Prepend(Process{0})
	cll.Append(Process{1})
	cll.Prepend(Process{-1})
	cll.Append(Process{2})
	
	cll.Display()
	fmt.Println(cll.Delete(Process{2}))
	
	dll := &linkedlist.DoublyLinkedList[Process]{}
	
	dll.Prepend(Process{0})
	dll.Display()
	

	origInt := []int{1, 2, 4, 5}
	reversed := stack.ReverseArray(origInt)
	fmt.Println(reversed)
	
	origStr := []rune{'s', 't', 'r', 'i', 'n', 'g', 's'}
	reversed1 := stack.ReverseArray(origStr)
	fmt.Println(string(reversed1))
	
	fmt.Println("--- Parentheses Checker ---")

	expressions := []string{
		"{([])}",          // Balanced
		"([{}])()",        // Balanced
		"((()))",          // Balanced
		"var x = func({a:1, b:[2,3]});", // Balanced (ignores non-parentheses)
		"((()",            // Unbalanced (too many open)
		"())",             // Unbalanced (closing before open)
		"{[}]",            // Unbalanced (mismatched and wrong order)
		"({[}])",          // Unbalanced (mismatched order)
		"([)]",            // Unbalanced (mismatched order)
		"",                // Balanced (empty string)
		"abc",             // Balanced (no parentheses)
		"[",               // Unbalanced
		"}",               // Unbalanced
	}

	for _, expr := range expressions {
		if stack.ValidateExpression(expr) {
			fmt.Printf("Expression '%s' is BALANCED.\n", expr)
		} else {
			fmt.Printf("Expression '%s' is UNBALANCED.\n", expr)
		}
	}
	expression := "( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )"
	fmt.Println(stack.EvalParenthesizedInfixExpression(expression))
}
