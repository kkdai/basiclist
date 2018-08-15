package basiclist

type DoubleBasicNode struct {
	key  uint
	val  interface{}
	prev *DoubleBasicNode
	next *DoubleBasicNode
}

func NewDoubleBasicNode(key uint, val interface{}) DoubleBasicNodeInterface {
	return &DoubleBasicNode{key, val, nil, nil}
}

func (node *DoubleBasicNode) Key() uint {
	return node.key
}

func (node *DoubleBasicNode) Value() *interface{} {
	return &node.val
}

func (node *DoubleBasicNode) HasPrevious() bool {
	return node.prev != nil
}

func (node *DoubleBasicNode) Previous() DoubleBasicNodeInterface {
	return node.prev
}

func (node *DoubleBasicNode) HasNext() bool {
	return node.next != nil
}

func (node *DoubleBasicNode) Next() DoubleBasicNodeInterface {
	return node.next
}

