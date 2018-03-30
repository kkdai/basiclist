package basiclist

type DoubleBasicNode struct {
	key  uint
	val  interface{}
	prev *DoubleBasicNode
	next *DoubleBasicNode
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

func (node *DoubleBasicNode) Previous() *DoubleBasicNode {
	return node.prev
}

func (node *DoubleBasicNode) Next() *DoubleBasicNode {
	return node.next
}

func (node *DoubleBasicNode) HasNext() bool {
	return node.next != nil
}