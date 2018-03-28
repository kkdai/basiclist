package basiclist

import (
	"fmt"
	"testing"
)

func TestLinkedNode(t *testing.T) {
	headNode := NewBasicNode(0, nil)
	headNode.next = NewBasicNode(1,  "node1")
	headNode.next.next = NewBasicNode(2,  "node2")

	fmt.Printf("head->")
	currentNode := headNode
	for {
		fmt.Printf("[key:%d][val:%v]->", currentNode.Key(), currentNode.Value())
		if currentNode.Next() == nil {
			break
		}
		currentNode = currentNode.Next()
	}
	fmt.Printf("nil\n")
}

func TestBasicList(t *testing.T) {
	bList := NewBasicList()
	bList.Insert(3, "string3")
	bList.Insert(4, "string4")
	bList.Insert(2, "string2")
	bList.Insert(5, "string5")
	bList.Insert(7, "string7")
	bList.Insert(9, "string9")
	bList.Insert(11, "string11")
	bList.DisplayAll()

	ret, err := bList.Search(5)
	if err != nil || ret != "string5" {
		t.Errorf("Search failed.\n")
	}

	bList.Remove(7)
	bList.DisplayAll()
}

func TestBasicListRemove(t *testing.T) {
	bList := NewBasicList()
	bList.Insert(4, "string4")
	bList.Insert(3, "string3")
	bList.DisplayAll()
	bList.Remove(3)
	bList.Remove(4)
	bList.DisplayAll()
	if bList.Remove(4) == nil {
		t.Errorf("Over remove\n")
	}
}
