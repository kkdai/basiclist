package basiclist

type DoubleBasicNode struct {
	key  uint
	val  interface{}
	prev *DoubleBasicNode
	next *DoubleBasicNode
}

func NewDoubleBasicHead() *DoubleBasicNode {
	return &DoubleBasicNode{0, nil, nil, nil}
}

func NewDoubleBasicNode(key uint, val interface{}) *DoubleBasicNode {
	return &DoubleBasicNode{key, val, nil, nil}
}

func (node *DoubleBasicNode) Key() uint {
	return node.key
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
		(node.key == 0 && node.Previous().key > 0)
}
