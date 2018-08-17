package basiclist

import (
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestDoubleBasicList_Head(t *testing.T) {
	list := NewDoubleBasicList()
	list.DisplayAll()
}

func TestDoubleBasicList_Insert(t *testing.T) {
	list := NewDoubleBasicList()
	list.Insert(0, "zero1")
	list.Insert(0, "zero2")
	list.Insert(1, "one1")
	list.Insert(1, "one2")
	list.Insert(1, "one3")
	list.Insert(20, "tw1")
	list.Insert(19, "nin1")
	list.Insert(18, "eit1")
	list.Insert(17, "sev1")
	list.Insert(17, "sev2")
	list.Insert(18, "eit2")
	list.Insert(19, "nin2")
	list.Insert(20, "tw2")
	list.DisplayAll()
}

func TestDoubleBasicList_Tail(t *testing.T) {
	list := NewDoubleBasicList()
	list.Insert(10, "bigger")
	list.Insert(15, "biggest")
	list.Insert(5, "big")
	assertEqual(t, list.Tail().Key(), uint(15))
	assertEqual(t, *list.Tail().Value(), "biggest")
	list.DisplayAll()
}

func TestDoubleBasicList_Remove(t *testing.T) {
	list := NewDoubleBasicList()
	list.Insert(0, "zero1")
	list.Insert(0, "zero2")
	list.Insert(1, "one1")
	list.Insert(1, "one2")
	list.Insert(1, "one3")
	list.Remove(1, "one1")
	list.DisplayAll()
}

func TestDoubleBasicList_RemoveAll(t *testing.T) {
	list := NewDoubleBasicList()
	list.Insert(0, "zero1")
	list.Insert(0, "zero2")
	list.Insert(1, "one1")
	list.Insert(1, "one2")
	list.Insert(2, "two1")
	list.Insert(2, "two2")
	list.RemoveAll(0)
	list.DisplayAll()
}

func TestDoubleBasicList_RemoveAll2(t *testing.T) {
	list := NewDoubleBasicList()
	list.Insert(0, "zero1")
	list.Insert(0, "zero2")
	list.Insert(1, "one1")
	list.Insert(1, "one2")
	list.Insert(2, "two1")
	list.Insert(2, "two2")
	list.RemoveAll(1)
	list.DisplayAll()
}

func TestDoubleBasicList_RemoveAll3(t *testing.T) {
	list := NewDoubleBasicList()
	list.Insert(0, "zero1")
	list.Insert(0, "zero2")
	list.Insert(1, "one1")
	list.Insert(1, "one2")
	list.Insert(2, "two1")
	list.Insert(2, "two2")
	list.RemoveAll(2)
	list.DisplayAll()
}

func TestDoubleBasicList_Inde1(t *testing.T) {
	list := NewDoubleBasicList()
	list.Insert(0, "zero1")
	list.Insert(0, "zero2")
	list.Insert(1, "one1")
	list.Insert(1, "one2")
	list.Insert(2, "two1")
	list.Insert(2, "two2")
	list.DisplayAll()
	node, _ := list.Index(1)
	assertEqual(t, node, "zero1")
}

func TestDoubleBasicList_Index2(t *testing.T) {
	list := NewDoubleBasicList()
	list.Insert(0, "zero1")
	list.Insert(0, "zero2")
	list.Insert(1, "one1")
	list.Insert(1, "one2")
	list.Insert(2, "two1")
	list.Insert(2, "two2")
	list.DisplayAll()
	node, _ := list.Index(2)
	assertEqual(t, node, "zero2")
}

func TestDoubleBasicList_Index3(t *testing.T) {
	list := NewDoubleBasicList()
	list.Insert(0, "zero1")
	list.Insert(0, "zero2")
	list.Insert(1, "one1")
	list.Insert(1, "one2")
	list.Insert(2, "two1")
	list.Insert(2, "two2")
	list.DisplayAll()
	node, _ := list.Index(3)
	assertEqual(t, node, "one1")
}

func TestDoubleBasicList_Search(t *testing.T) {
	list := NewDoubleBasicList()
	list.Insert(0, "zero1")
	list.Insert(0, "zero2")
	list.Insert(1, "one1")
	list.Insert(1, "one2")
	list.Insert(2, "two1")
	list.Insert(2, "two2")
	node, _ := list.Search(2)
	assertEqual(t, node, "two1")
}