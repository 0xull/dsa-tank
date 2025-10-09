package leetcode


import (
	"testing"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr, rem := head, 0

	for l1 != nil || l2 != nil || rem != 0 {
		sum := 0 + rem
		
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		} 
		
	 	if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next

		if sum >= 10 {
			rem = 1
		} else {
			rem = 0
		}
	}

	return head.Next
}

func newLinkedList(list []int) *ListNode {
	ll := &ListNode{Val: list[0]}
	curr := ll
	for i := 1; i < len(list); i++ {
		curr.Next = &ListNode{Val: list[i]}
		curr = curr.Next
	}
	return ll
}

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		name string
		l1 *ListNode
		l2 *ListNode
	}{
		{
			name: "short addition",
			l1: newLinkedList([]int{2, 4, 3}), // number 342 stored as [2 -> 4 -> 3]
			l2: newLinkedList([]int{5, 6, 4}), // number 465 stored as [5 -> 6 -> 4]
		},
		{
			name: "long addition",
			l1: newLinkedList([]int{9, 9, 9, 9, 9, 9, 9}), // number 9999999 stored as [9 -> 9 -> 9 -> 9 -> 9 -> 9 -> 9]
			l2: newLinkedList([]int{9, 9, 9, 9}),          // number 9999 stored as [9 -> 9 -> 9 -> 9]
		},
	}
	
	expected := []*ListNode{
		newLinkedList([]int{7,0,8}),
		newLinkedList([]int{8,9,9,9,0,0,0,1}),
	} 
	
	for i, tc := range testCases {
		t.Run(tc.name,func(t *testing.T) {
			got := addTwoNumbers(tc.l1, tc.l2)
			lExpected := expected[i]
			for got != nil && lExpected != nil {
				if got.Val != lExpected.Val {
					t.Fatalf("expected '%d'; got '%d'", lExpected.Val, got.Val)
				}
				got, lExpected = got.Next, lExpected.Next
			}
			if got != nil || lExpected != nil {
				t.Fatal("expects both 'got' and 'expected' linked list to be same length")
			}
		})
	}

	
}
