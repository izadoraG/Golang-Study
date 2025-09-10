package main

type DoublyLinkedListNode struct {
	data int
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

func reverse(llist *DoublyLinkedListNode) *DoublyLinkedListNode {
	current := llist
	prev := current.prev

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}
