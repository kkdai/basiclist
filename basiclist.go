package basiclist

import (
	"errors"
	"fmt"
)

type BasicList struct {
	head *BasicNode
}

//NewBasicList : Init structure for basic Sorted Linked List.
func NewBasicList() *BasicList {
	var empty interface{}
	return &BasicList{head: NewBasicNode(0, empty)}
}

func (b *BasicList) Head() *BasicNode {
	return b.head
}

func (b *BasicList) Insert(key int, value interface{}) {
	if b.head == nil {
		// fmt.Println("note is empty")
		b.head = NewBasicNode(key, value)
	} else {
		var currentNode *BasicNode
		currentNode = b.head
		var previouNote *BasicNode
		var found bool
		newNode := NewBasicNode(key, value)

		for {
			if currentNode.key > key {
				newNode.next = previouNote.Next()
				previouNote.next = newNode
				found = true
				break
			}

			if currentNode.Next() == nil {
				break
			}

			previouNote = currentNode
			currentNode = currentNode.Next()
		}

		if found == false {
			currentNode.next = newNode
		}
	}
}

func (b *BasicList) Search(key int) (interface{}, error) {
	currentNode := b.head
	for {
		if currentNode.key == key {
			return currentNode.val, nil
		}

		if currentNode.Next() == nil {
			break
		}
		currentNode = currentNode.Next()
	}
	return nil, errors.New("key not found")
}

func (b *BasicList) Remove(key int) error {
	currentNode := b.head
	var previousNote *BasicNode
	for {
		if currentNode.key == key {
			previousNote.next = currentNode.Next()
			return nil
		}

		if currentNode.Next() == nil {
			break
		}
		previousNote = currentNode
		currentNode = currentNode.Next()
	}
	return errors.New("key not found")
}

func (b *BasicList) DisplayAll() {
	fmt.Println("")
	fmt.Printf("head->")
	currentNode := b.head
	for {
		fmt.Printf("[key:%d][val:%v]->", currentNode.key, currentNode.val)
		if currentNode.Next() == nil {
			break
		}
		currentNode = currentNode.Next()
	}
	fmt.Printf("nil\n")
}
