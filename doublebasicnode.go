package basiclist

type DoubleBasicNode struct {
	key  int
	val  interface{}
	prev *DoubleBasicNode
	next *DoubleBasicNode
}

func NewDoubleBasicHead() *DoubleBasicNode {
	return &DoubleBasicNode{-1, nil, nil, nil}
}

func NewDoubleBasicNode(key uint, val interface{}) *DoubleBasicNode {
	return &DoubleBasicNode{int(key), val, nil, nil}
}

func (node *DoubleBasicNode) Key() uint {
	return uint(node.key)
}

func (node *DoubleBasicNode) Value() *interface{} {
	return &node.val
}

func (node *DoubleBasicNode) HasPrevious() bool {
	return node != nil && !node.IsHead() && node.Previous() != nil && !node.Previous().IsHead()
}

func (node *DoubleBasicNode) Previous() *DoubleBasicNode {
	return node.prev
}

func (node *DoubleBasicNode) HasNext() bool {
	return node.next != nil
}

func (node *DoubleBasicNode) Next() *DoubleBasicNode {
	return node.next
}

func (node *DoubleBasicNode) IsHead() bool {
	return node.Previous() == nil ||
		(node.key == -1 && node.Previous().key > -1)
}
