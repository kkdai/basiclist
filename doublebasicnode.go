package basiclist

type DoubleBasicNode struct {
	key  int
	val  interface{}
	prev *DoubleBasicNode
	next *DoubleBasicNode
	isRoot bool
}

func newDoubleBasicRoot() *DoubleBasicNode {
	return &DoubleBasicNode{-1, nil, nil, nil, true}
}

func NewDoubleBasicNode(key uint, val interface{}) *DoubleBasicNode {
	return &DoubleBasicNode{int(key), val, nil, nil, false}
}

func (node *DoubleBasicNode) Key() uint {
	return uint(node.key)
}

func (node *DoubleBasicNode) Value() *interface{} {
	return &node.val
}

func (node *DoubleBasicNode) HasPrevious() bool {
	return node != nil && !node.IsRoot() && node.Previous() != nil && !node.Previous().IsRoot()
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

func (node *DoubleBasicNode) IsRoot() bool {
	return node.isRoot
}
