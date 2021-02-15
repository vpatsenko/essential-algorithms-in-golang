package main

import (
	"errors"
	"fmt"
)

type node struct {
	data     int
	nextNode *node
	previous *node
}

type LinkedList struct {
	head *node
	tail *node
}

func NewLinkedList() *LinkedList {
	head := &node{}
	tail := &node{}

	head.nextNode = tail
	tail.previous = head

	return &LinkedList{
		head: head,
		tail: tail,
	}
}

func (ll *LinkedList) IterateFromTail() {
	current := ll.tail
	for current.previous != nil {
		current = current.previous
		if current == ll.head {
			break
		}
		if current.previous == ll.head {
			fmt.Println(current)
			break
		}
		fmt.Print(current, " ==> ")
	}
}

func (ll *LinkedList) IterateFromHead() {
	current := ll.head
	for current.nextNode != nil {
		current = current.nextNode
		if current == ll.tail {
			break
		}
		if current.nextNode == ll.tail {
			fmt.Println(current)
			break
		}
		fmt.Print(current, " ==> ")
	}
}

func (ll *LinkedList) FindNode(data int) (*node, error) {
	current := ll.head
	for current.nextNode != nil {
		if current.data == data {
			return current, nil
		}
		current = current.nextNode
	}
	return nil, errors.New("no data in list")
}

func (ll *LinkedList) FindNodeBefore(data int) (*node, error) {
	current := ll.head
	for current.nextNode != nil {
		if current.nextNode.data == data {
			return current, nil
		}
		current = current.nextNode
	}
	return nil, errors.New("no data in list")
}

func (ll *LinkedList) AddAtBegginning(data int) {
	nodeAfterHead := ll.head.nextNode
	newNode := &node{
		data:     data,
		previous: ll.head,
		nextNode: nodeAfterHead,
	}

	ll.head.nextNode = newNode
	nodeAfterHead.previous = newNode
}

func (ll *LinkedList) AddAtEnd(data int) {
	nodeBeforeTail := ll.tail.previous
	newNode := &node{
		data: data,
	}

	newNode.nextNode = ll.tail
	newNode.previous = nodeBeforeTail
	ll.tail.previous = newNode
	nodeBeforeTail.nextNode = newNode
}

func (ll *LinkedList) InsertNode(data int, nodeBeforeInsertion *node) {
	newNode := &node{
		data: data,
	}

	// update references to next node
	newNode.nextNode = nodeBeforeInsertion.nextNode
	nodeBeforeInsertion.nextNode = newNode

	// updated references to previous node
	newNode.nextNode.previous = newNode
	newNode.previous = nodeBeforeInsertion
}

func (ll *LinkedList) DeleteAfter(nodeBeforeDelete *node) {
	nodeToBeDeleted := nodeBeforeDelete.nextNode
	nodeAfterDeleted := nodeToBeDeleted.nextNode

	nodeBeforeDelete.nextNode = nodeAfterDeleted
	nodeAfterDeleted.previous = nodeBeforeDelete
	nodeToBeDeleted = nil
}

func (ll *LinkedList) CopyList() *LinkedList {
	llCopy := new(LinkedList)
	head := new(node)
	tail := new(node)

	head.data = 0
	head.nextNode = tail
	head.previous = nil

	tail.data = 0
	tail.previous = head
	tail.nextNode = nil

	llCopy.head = head
	llCopy.tail = tail

	current := ll.head.nextNode
	for current.nextNode != nil {
		llCopy.AddAtEnd(current.data)
		current = current.nextNode
	}

	return llCopy
}

func (ll *LinkedList) InsertionSort() *LinkedList {

	sorted := NewLinkedList()
	input := ll.head.nextNode

	for input != ll.tail {
		nextNode := input
		input = input.nextNode

		afterMe := sorted.head
		for afterMe.nextNode != sorted.tail && afterMe.nextNode.data < nextNode.data {
			afterMe = afterMe.nextNode
		}
		sorted.InsertNode(nextNode.data, afterMe)
	}

	return sorted
}

func main() {
	fmt.Println()
	ll := NewLinkedList()

	fmt.Println("=== AddAtBegginning ===")
	ll.AddAtBegginning(1)
	ll.AddAtBegginning(2)
	ll.AddAtBegginning(3)
	ll.IterateFromHead()
	fmt.Println()
	fmt.Println()

	ll.AddAtEnd(2)
	fmt.Println("=== AddAtEnd ===")
	ll.AddAtEnd(10)
	ll.IterateFromHead()
	fmt.Println()
	fmt.Println()

	fmt.Println("=== FindNode ===")
	node, _ := ll.FindNode(1)
	fmt.Println(node)
	fmt.Println()
	fmt.Println()

	fmt.Println("=== FindNodeBefore ===")
	node, _ = ll.FindNodeBefore(1)
	fmt.Println(node)
	fmt.Println()
	fmt.Println()

	fmt.Println("=== InsertNode ===")
	nodeBeforeInsertion, _ := ll.FindNode(2)
	ll.InsertNode(100, nodeBeforeInsertion)
	ll.IterateFromHead()
	fmt.Println()
	fmt.Println()

	fmt.Println("=== Delete ===")
	nodeBeforeInsertion, _ = ll.FindNode(2)
	ll.DeleteAfter(nodeBeforeInsertion)
	ll.IterateFromHead()
	fmt.Println()
	fmt.Println()

	fmt.Println("=== IterateFromHead ===")
	ll.IterateFromHead()
	fmt.Println()
	fmt.Println()

	fmt.Println("=== IterateFromTail ===")
	ll.IterateFromTail()
	fmt.Println()

	fmt.Println("=== CopyList ===")

	ll.IterateFromHead()
	fmt.Println()
	llCopy := ll.CopyList()
	llCopy.IterateFromHead()
	fmt.Println()

	fmt.Println()
	fmt.Println()
	sorted := llCopy.InsertionSort()
	sorted.IterateFromHead()
	sorted.IterateFromTail()
}