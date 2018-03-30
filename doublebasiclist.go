package basiclist

import (
	"fmt"
	"errors"
)

type DoubleBasicList struct {
	head *DoubleBasicNode
}

//NewBasicList : Init structure for basic Sorted Doubly Linked List.
func NewDoubleBasicList() *DoubleBasicList {
	var empty interface{}
	return &DoubleBasicList{head: NewDoubleBasicNode(0, empty)}
}

func (b *DoubleBasicList) Head() *DoubleBasicNode {
	return b.head
}

func (b *DoubleBasicList) Tail() *DoubleBasicNode {
	return b.Head().Previous()
}

func (b *DoubleBasicList) Insert(key uint, value interface{}) error {

	newNode := NewDoubleBasicNode(key, value)

	if b.head == nil {
		return errors.New("insert failed due to unexpected state")
	} else if b.Tail() == nil { /* First Inserted Node */
		newNode.prev = b.Head()
		b.Head().next = newNode
		b.Head().prev = newNode
	} else if newNode.Key() >= b.Tail().Key() { /* Largest new tail    */
		newNode.prev = b.Tail()
		// order of operations matters here
		b.Tail().next = newNode
		b.Head().prev = newNode
	} else {
		current := b.Head()
		previous := b.Tail()

		for {
			if newNode.Key() < current.Key() {
				newNode.next = current
				newNode.prev = previous
				previous.next = newNode
				current.prev = newNode
				break
			} else {
				if current.HasNext() {
					previous = current
					current = current.Next()
				} else {
					return errors.New("insert failed due to unexpected state")
				}
			}
		}
	}

	return nil
}

func (b *DoubleBasicList) Search(key uint) (interface{}, error) {
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

func (b *DoubleBasicList) Remove(key uint, value interface{}) error {

	current := b.Head().Next()
	previous := b.Head()

	for {
		if key == current.Key() && value == *current.Value() {
			previous.next = current.Next()
			current.Next().prev = previous
			break
		} else {
			if current.HasNext() {
				previous = current
				current = current.Next()
			} else {
				return errors.New("key not found")
			}
		}
	}

	return nil
}

func (b *DoubleBasicList) RemoveAll(key uint) error {

	current := b.Head().Next()
	previous := b.Head()
	found := false

	for {
		if key < current.Key() && found {
			break
		} else if key == current.Key() {
			found = true
			previous.next = current.Next()
			if current.HasNext() {
				current.Next().prev = previous
				current = current.Next()
			} else {
				break
			}
		} else {
			if current.HasNext() {
				previous = current
				current = current.Next()
			} else {
				return errors.New("key not found")
			}
		}


	}

	return nil
}

func (b *DoubleBasicList) DisplayAll() {
	fmt.Println("")
	fmt.Printf("<-head->")
	current := b.head
	for {
		fmt.Printf("<-[key:%d][val:%v]->", current.key, current.val)
		if !current.HasNext() {
			break
		}
		current = current.Next()
	}
	fmt.Printf("nil\n")
}
