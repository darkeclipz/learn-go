package main

import (
	"errors"
	"fmt"
)

type DoubleLinkedList[T any] struct {
	head *DoubleLinkedListNode[T]
	tail *DoubleLinkedListNode[T]
	size uint
}

func createDoubleLinkedList[T any](item T) *DoubleLinkedList[T] {
	root := &DoubleLinkedListNode[T]{item, nil, nil}
	return &DoubleLinkedList[T]{root, root, 1}
}

type DoubleLinkedListNode[T any] struct {
	item     T
	next     *DoubleLinkedListNode[T]
	previous *DoubleLinkedListNode[T]
}

func (list *DoubleLinkedList[T]) insert(x T) {
	item := DoubleLinkedListNode[T]{x, nil, nil}
	list.tail.next = &item
	item.previous = list.tail
	list.tail = &item
	list.size++
}

func (list *DoubleLinkedList[T]) find(predicate func(*DoubleLinkedListNode[T]) bool) (*DoubleLinkedListNode[T], error) {
	current := list.head
	for current != nil {
		if predicate(current) {
			return current, nil
		}
		current = current.next
	}
	return nil, errors.New("no node matching the predicate")
}

func main() {

	myLinkedList := createDoubleLinkedList(0) // := &DoubleLinkedListNode[int]{0, nil, nil}

	for i := range 10 {
		myLinkedList.insert(i + 1)
	}

	node := myLinkedList.head
	for node != nil {
		fmt.Println(node.item)
		node = node.next
	}

	fmt.Printf("The size of the list is %d\n", myLinkedList.size)

	foundNode, err := myLinkedList.find(func(node *DoubleLinkedListNode[int]) bool { return node.item == 5 })
	if foundNode != nil {
		fmt.Printf("Found node with item 5: %v\n", foundNode.item)
	} else {
		fmt.Printf("The node doesn't exists: %v\n", err)
	}
}
