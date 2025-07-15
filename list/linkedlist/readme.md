# Linked Lists ADT

## 6.2 Singly Linked Lists

## 6.3 Circular Linked Lists

A **circular linked list** is a type of linked list in which the last node does not point to `nil`. Instead, the `Next` pointer of the last node points back to the **first node** (the `Head`) of the list, forming a circle or a loop.

**Visual Representation:**

Imagine our singly linked list:
`Head -> NodeA -> NodeB -> NodeC -> nil`

In a circular linked list, it would look like this:
```
  Head
    |
    V
  +-------+------+    +-------+------+    +-------+------+
  | DataA | Next | Règne| DataB | Next | Règne| DataC | Next |---
  +-------+------+    +-------+------+    +-------+------+   |
    ^                                                        |
    |--------------------------------------------------------|
```
Here, `NodeC` (the last node) points back to `NodeA` (the `Head`).

**Key Characteristics & Differences:**

1.  **No `nil` End:** The most defining feature is the absence of a `nil` pointer at the end of the list (unless the list is empty, in which case `Head` itself would be `nil`).
2.  **Traversal:** You can start at any node and traverse the entire list, eventually returning to the starting node. This is useful for applications where processing needs to cycle through elements, like a round-robin scheduler or managing turns in a game.
3.  **Reaching any Node:** Every node is accessible from any other node in the list by repeatedly following the `Next` pointers.
4.  **Caution with Loops:** Traversal algorithms must be careful to avoid infinite loops. The condition to stop traversing is no longer checking for `nil`, but rather checking if we've returned to the starting point (e.g., the `Head` node).
5.  **Head/Tail Pointers:** While you can maintain just a `Head` pointer, circular lists sometimes use a pointer to the **last node** (let's call it `Tail`). If you have a `Tail` pointer in a circular list, you can easily access the `Head` via `Tail.Next` and also the `Tail` itself. This can make operations like inserting at the beginning or end very efficient ($O(1)$). For our discussion, we will primarily assume our `CircularLinkedList` struct still holds a `Head` pointer, and the circularity is maintained by the last node's `Next` pointer referencing this `Head`.

**Go Representation:**

The `Node` structure itself remains the same as for a singly linked list. The difference lies in the *logic* of how nodes are connected and how the list is managed.

```go
// Node structure (same as for singly linked list)
type Node struct {
    Data interface{}
    Next *Node
}

// CircularLinkedList structure
type CircularLinkedList struct {
    Head *Node // Points to the "first" node in the conceptual circle
    Size int   // Number of nodes in the list
}
```

**Traversal in a Circular Linked List:**

Traversing a circular linked list requires a different stopping condition. If we simply loop while `currentNode != nil`, we'd loop forever (unless the list is empty).

A common approach:
1.  If the list is empty (`list.Head == nil`), there's nothing to traverse.
2.  Start with `currentNode = list.Head`.
3.  Process `currentNode`.
4.  Move to `currentNode = currentNode.Next`.
5.  Repeat steps 3 and 4 *until* `currentNode` becomes equal to `list.Head` again.

This is often implemented using a `do-while` loop structure in languages that support it.
