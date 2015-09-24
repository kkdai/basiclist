package basiclist

import (
	"errors"
	"fmt"
)

type BasicNote struct {
	key      int
	val      interface{}
	nextNode *BasicNote
}

type BasicList struct {
	head *BasicNote
}

//NewBasicList : Init structure for basic Sorted Linked List.
func NewBasicList() *BasicList {
	var empty interface{}
	return &BasicList{head: &BasicNote{key: 0, val: empty, nextNode: nil}}
}

func (b *BasicList) Insert(key int, value interface{}) {
	if b.head == nil {
		// fmt.Println("note is empty")
		b.head = &BasicNote{key: key, val: value, nextNode: nil}
	} else {
		var currentNode *BasicNote
		currentNode = b.head
		var previouNote *BasicNote
		var found bool
		newNode := &BasicNote{key: key, val: value, nextNode: nil}

		for {
			if currentNode.key > key {
				newNode.nextNode = previouNote.nextNode
				previouNote.nextNode = newNode
				found = true
				break
			}

			if currentNode.nextNode == nil {
				break
			}

			previouNote = currentNode
			currentNode = currentNode.nextNode
		}

		if found == false {
			currentNode.nextNode = newNode
		}
	}
}

func (b *BasicList) Search(key int) (interface{}, error) {
	currentNode := b.head
	for {
		if currentNode.key == key {
			return currentNode.val, nil
		}

		if currentNode.nextNode == nil {
			break
		}
		currentNode = currentNode.nextNode
	}
	return nil, errors.New("Not found.")
}

func (b *BasicList) Remove(key int) error {
	currentNode := b.head
	var previouNote *BasicNote
	for {
		if currentNode.key == key {
			previouNote.nextNode = currentNode.nextNode
			return nil
		}

		if currentNode.nextNode == nil {
			break
		}
		previouNote = currentNode
		currentNode = currentNode.nextNode
	}
	return errors.New("Not found key.")
}

func (b *BasicList) DisplayAll() {
	fmt.Println("")
	fmt.Printf("head->")
	currentNode := b.head
	for {
		fmt.Printf("[key:%d][val:%v]->", currentNode.key, currentNode.val)
		if currentNode.nextNode == nil {
			break
		}
		currentNode = currentNode.nextNode
	}
	fmt.Printf("nil\n")
}
