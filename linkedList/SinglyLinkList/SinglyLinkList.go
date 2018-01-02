package SinglyLinkList

import "fmt"

type Link struct {
	Data interface{}
	Next *Link
}

type SinglyLinkedList struct {
	front  *Link
	length int
}

func (list *SinglyLinkedList) Init() *SinglyLinkedList {
	list.length = 0
	return list
}

func New() *SinglyLinkedList {
	return new(SinglyLinkedList).Init()
}

func (list *SinglyLinkedList) Front() *Link {
	return list.front
}

func (list *SinglyLinkedList) Back() *Link {
	currentNode := list.front
	for currentNode != nil && currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	return currentNode
}

func (list *SinglyLinkedList) Append(link *Link) {
	if list.front == nil {
		list.front = link
	} else {
		currentNode := list.front

		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}

		currentNode.Next = link
	}

	list.length++
}

func (list *SinglyLinkedList) Prepend(link *Link) {
	if list.front == nil {
		list.front = link
	} else {
		link.Next = list.front
		list.front = link
	}

	list.length++
}

func (list *SinglyLinkedList) Remove(link *Link) {
	if list.front == link {
		list.front = link.Next
		list.length--
	} else {
		currentNode := list.front

		for currentNode != nil && currentNode.Next != nil && currentNode.Next != link {
			currentNode = currentNode.Next
		}

		if currentNode.Next == link {
			currentNode.Next = currentNode.Next.Next
			list.length--
		}
	}
}

func (list *SinglyLinkedList) Find(value interface{}) *Link {
	currentNode := list.front
	for currentNode != nil && currentNode.Data != value && currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	return currentNode
}

func (list *SinglyLinkedList) Print() {
	currentNode := list.front
	for currentNode != nil && currentNode.Next != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}
	fmt.Println(currentNode.Data, "\n")
}

func (list *SinglyLinkedList) Length() int {
	return list.length
}
