package main

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func printLinkedList(head *SinglyLinkedListNode) {
	current := head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
