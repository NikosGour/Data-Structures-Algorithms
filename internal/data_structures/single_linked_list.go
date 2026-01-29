package data_structures

type SingleLinkedList[T any] struct {
	head *Node[T]
	size int
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

func NewSingleLinkedList[T any]() *SingleLinkedList[T] {
	sll := &SingleLinkedList[T]{}
	return sll
}

func (sll *SingleLinkedList[T]) Get(i int) (T, error) {
	var defaultVal T
	idx := 0
	v := sll.head

	for ; idx <= i; idx++ {
		if v == nil {
			return defaultVal, ErrIndexOutOfBounds
		}

		if idx == i {
			return v.value, nil
		}

		v = v.next
	}

	return defaultVal, ErrUnreachable
}

func (sll *SingleLinkedList[T]) Set(i int, v T) error {
	idx := 0
	node := sll.head

	for ; idx <= i; idx++ {
		if node == nil {
			return ErrIndexOutOfBounds
		}

		if idx == i {
			node.value = v
		}

		node = node.next
	}

	return ErrUnreachable
}
