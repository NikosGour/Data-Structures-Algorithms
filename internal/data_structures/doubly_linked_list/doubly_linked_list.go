package doubly_linked_list

import (
	"dsa/internal/data_structures"
)

type DoublyLinkedList[T any] struct {
	head *Node[T]
	size int
}

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	sll := &DoublyLinkedList[T]{}
	return sll
}
func NewDoublyLinkedListFromSlice[T any](arr []T) *DoublyLinkedList[T] {
	sll := &DoublyLinkedList[T]{}
	if len(arr) > 0 {
		sll.head = NewNode(arr[0])
		sll.size++
	}

	n := sll.head
	for i, v := range arr {
		if i == 0 {
			continue
		}

		n.next = NewNode(v)
		n.next.prev = n
		sll.size++
		n = n.next
	}

	return sll
}

func NewNode[T any](v T) *Node[T] {
	n := &Node[T]{value: v}
	return n
}

func (sll *DoublyLinkedList[T]) Get(i int) (T, error) {
	var defaultVal T
	if i >= sll.size || i < 0 {
		return defaultVal, data_structures.ErrIndexOutOfBounds
	}
	idx := 0
	v := sll.head

	for ; idx <= i; idx++ {
		if v == nil {
			return defaultVal, data_structures.ErrIndexOutOfBounds
		}

		if idx == i {
			return v.value, nil
		}

		v = v.next
	}

	return defaultVal, data_structures.ErrUnreachable
}

func (sll *DoublyLinkedList[T]) Set(i int, v T) error {
	if i >= sll.size || i < 0 {
		return data_structures.ErrIndexOutOfBounds
	}

	idx := 0
	node := sll.head

	for ; idx <= i; idx++ {
		if node == nil {
			return data_structures.ErrIndexOutOfBounds
		}

		if idx == i {
			node.value = v
			return nil
		}

		node = node.next
	}

	return data_structures.ErrUnreachable
}

func (sll *DoublyLinkedList[T]) Add(v T) {
	if sll.head == nil {
		sll.head = NewNode(v)
		sll.size++
		return
	}

	n := sll.head
	for n.next != nil {
		n = n.next
	}

	n.next = NewNode(v)
	n.next.prev = n
	sll.size++
}
func (sll *DoublyLinkedList[T]) Remove(i int) (T, error) {
	var defaultVal T
	if i >= sll.size || i < 0 {
		return defaultVal, data_structures.ErrIndexOutOfBounds
	}

	idx := 0
	node := sll.head
	prev := sll.head

	for ; idx <= i; idx++ {
		if node == nil {
			return defaultVal, data_structures.ErrIndexOutOfBounds
		}

		if i == 0 {
			sll.head = node.next
			sll.head.prev = nil
			sll.size--
			return node.value, nil
		}

		if idx == i {
			prev.next = node.next
			if node.next != nil {
				node.next.prev = prev
			}
			sll.size--
			return node.value, nil
		}

		prev = node
		node = node.next
	}

	return defaultVal, data_structures.ErrUnreachable
}

func (sll *DoublyLinkedList[T]) Length() int {
	return sll.size
}
