package drills

/*
	Linked Lists - Push & BuildOneTwoThree
	DESCRIPTION:
		Linked Lists - Push & BuildOneTwoThree

		Write push() and buildOneTwoThree() functions to easily update and initialize linked lists. Try to use the push() function within your buildOneTwoThree() function.

		Here's an example of push() usage:

		var chained = null
		chained = push(chained, 3)
		chained = push(chained, 2)
		chained = push(chained, 1)
		push(chained, 8) === 8 -> 1 -> 2 -> 3 -> null
		The push function accepts head and data parameters, where head is either a node object or null/None/nil. Your push implementation should be able to create a new linked list/node when head is null/None/nil.

		The buildOneTwoThree function should create and return a linked list with three nodes: 1 -> 2 -> 3 -> null
*/

type node[T comparable] struct {
	elem T
	next *node[T]
}

type list[T comparable] struct {
	head *node[T]
}

func Push[T comparable](l *list[T], value T) *list[T] {
	return &list[T]{head: &node[T]{next: l.head, elem: value}}
}

func BuildOneTwoThree[T comparable](first, second, third T) *list[T] {
	return &list[T]{head: &node[T]{elem: first, next: &node[T]{elem: second, next: &node[T]{elem: third}}}}
} 
