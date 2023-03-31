package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	currentNode := linkedList
	for currentNode.Next != nil {
		if currentNode.Value == currentNode.Next.Value {
			currentNode.Next = currentNode.Next.Next
		} else {
			currentNode = currentNode.Next
		}
	}
	return linkedList
}
