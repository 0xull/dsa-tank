package main

import linkedlist "github.com/IkehAkinyemi/dsa-tank/ds/list/linkedlist/singly-linkedlist"

type Process struct {
	PID int
}

func main() {
	sll := linkedlist.SinglyLinkedList[*Process]{
		Head: &linkedlist.Node[*Process]{Data: &Process{1}, Next: &linkedlist.Node[*Process]{Data: &Process{2}, Next: nil}},
	}

	sll.Display()
}
