package main

import (
	"fmt"

	linkedlist "github.com/IkehAkinyemi/dsa-tank/ds/list/linkedlist"
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
}
