package main

type ListNode struct {
	value *Node
	next  *ListNode
}

func initForwardList(head **ListNode) {
	*head = nil
}

func addToTail(head **ListNode, value *Node) {
	newNode := &ListNode{value, nil}
	current := *head

	if *head == nil {
		*head = newNode
		return
	}

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func findNode(head *ListNode, symbol string) *ListNode {
	current := head

	for current != nil {
		if current.value.symbol == symbol {
			return current
		}
		current = current.next
	}

	return nil
}

func deleteHead(head **ListNode) {
	if *head == nil {
		return
	}

	toDelete := *head
	*head = toDelete.next
}

func deleteCurrent(head **ListNode, nodeFound *ListNode) {
	if nodeFound == *head {
		deleteHead(head)
		return
	}

	current := *head
	for current.next != nodeFound {
		current = current.next
	}

	current.next = nodeFound.next
}

func clear(head **ListNode) {
	for *head != nil {
		destroyTree(&(*head).value)
		deleteHead(head)
	}
}
