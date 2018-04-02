package basiclist

type BasicNode struct {
	key  int
	val  interface{}
	next *BasicNode
}

func NewBasicNode(key int, val interface{}) *BasicNode {
	return &BasicNode{key, val, nil}
}

func (node *BasicNode) Key() int {
	return node.key
}

func (node *BasicNode) Value() *interface{} {
	return &node.val
}

func (node *BasicNode) Next() *BasicNode {
	return node.next
}

func (node *BasicNode) HasNext() bool {
	return node.next != nil
}