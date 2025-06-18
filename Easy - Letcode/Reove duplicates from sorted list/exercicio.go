package main

// Definição da estrutura ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head

	for cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next // Pula os nós duplicados
		} else {
			cur = cur.Next // Avança apenas se não houver duplicata
		}
	}

	return head
}
