package data_structures

import (
	"errors"
)

var (
	ErrIndexOutOfBounds  = errors.New("index out of bounds")
	ErrNotAllItemsCopied = errors.New("not all items copied")
	ErrUnreachable       = errors.New("unreachable")
)

type DynamicArray[T any] struct {
	items    []T
	size     int
	capacity int
}

func NewDynamicArray[T any]() *DynamicArray[T] {
	da := &DynamicArray[T]{}
	da.items = make([]T, 0)
	return da
}

func NewDynamicArrayFromSlice[T any](arr []T) *DynamicArray[T] {
	da := &DynamicArray[T]{items: arr, size: len(arr), capacity: len(arr)}
	return da
}

func (da *DynamicArray[T]) Get(i int) (T, error) {
	var rv T
	if i >= da.size || i < 0 {
		return rv, ErrIndexOutOfBounds
	}

	return da.items[i], nil
}

func (da *DynamicArray[T]) Set(i int, v T) error {
	if i >= da.size || i < 0 {
		return ErrIndexOutOfBounds
	}

	da.items[i] = v
	return nil
}

func (da *DynamicArray[T]) Length() int {
	return da.size
}

func (da *DynamicArray[T]) Add(v T) error {
	if da.size+1 > da.capacity {
		err := da.doubleArray()
		if err != nil {
			return err
		}
	}

	da.size++
	return da.Set(da.size-1, v)
}

func (da *DynamicArray[T]) Remove(i int) (T, error) {
	var rv T
	if i >= da.size || i < 0 {
		return rv, ErrIndexOutOfBounds
	}

	saved_v, err := da.Get(i)
	if err != nil {
		return rv, err
	}

	expected_copied := da.size - i - 1
	copied_items := copy(da.items[i:], da.items[i+1:da.size])
	if copied_items != expected_copied {
		return rv, ErrNotAllItemsCopied
	}

	da.size--
	rv = saved_v

	return rv, nil

}

func (da *DynamicArray[T]) doubleArray() error {
	if da.capacity == 0 {
		da.capacity = 1
	} else {
		da.capacity = da.capacity * 2
	}
	new_arr := make([]T, da.capacity)
	copied := copy(new_arr, da.items)
	if copied != da.size {
		return ErrNotAllItemsCopied
	}

	da.items = new_arr
	return nil
}
