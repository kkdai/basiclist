package basiclist

type DoubleBasicNodeInterface interface {
	Key() uint
	Value() *interface{}
	HasPrevious() bool
	Previous() DoubleBasicNodeInterface
	HasNext() bool
	Next() DoubleBasicNodeInterface
}
