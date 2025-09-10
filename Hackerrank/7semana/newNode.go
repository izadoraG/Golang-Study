package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	newNode := &SinglyLinkedListNode{data: data}

	// Caso inserir na cabeça (posição 0)
	if position == 0 {
		newNode.next = llist
		return newNode
	}

	current := llist
	// Percorre até o nó anterior à posição
	for i := 0; i < int(position-1); i++ {
		current = current.next
	}

	// Insere o novo nó
	newNode.next = current.next
	current.next = newNode

	return llist
}
