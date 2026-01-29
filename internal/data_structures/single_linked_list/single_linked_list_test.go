package single_linked_list

import (
	"dsa/internal/data_structures"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {

	assert := assert.New(t)
	tcs := []struct {
		Name             string
		SingleLinkedList *SingleLinkedList[int]
		Index            int
		ExpectedResult   int
		ExpectedError    error
	}{
		{
			Name:             "Get first item",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3}),
			Index:            0,
			ExpectedResult:   1,
			ExpectedError:    nil,
		},
		{
			Name:             "Get last item",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3}),
			Index:            2,
			ExpectedResult:   3,
			ExpectedError:    nil,
		},
		{
			Name:             "Inbounds index",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3}),
			Index:            1,
			ExpectedResult:   2,
			ExpectedError:    nil,
		},
		{
			Name:             "Inbounds index",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3, 5, 6, 7}),
			Index:            4,
			ExpectedResult:   6,
			ExpectedError:    nil,
		},
		{
			Name:             "Inbounds index",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3}),
			Index:            0,
			ExpectedResult:   1,
			ExpectedError:    nil,
		},
		{
			Name:             "Inbounds index",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3}),
			Index:            2,
			ExpectedResult:   3,
			ExpectedError:    nil,
		},
		{
			Name:             "Out of bounds index",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3}),
			Index:            3,
			ExpectedResult:   0,
			ExpectedError:    data_structures.ErrIndexOutOfBounds,
		},
		{
			Name:             "Negative index",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3}),
			Index:            -1,
			ExpectedResult:   0,
			ExpectedError:    data_structures.ErrIndexOutOfBounds,
		},
	}

	for _, test := range tcs {
		t.Run(test.Name, func(t *testing.T) {
			v, err := test.SingleLinkedList.Get(test.Index)
			if test.ExpectedError != nil {
				assert.EqualError(err, test.ExpectedError.Error(), test.Name)
				return
			}

			assert.NoError(err, test.Name)
			assert.Equal(test.ExpectedResult, v, test.Name)
		})
	}
}

func TestSet(t *testing.T) {

	assert := assert.New(t)
	tcs := []struct {
		Name             string
		SingleLinkedList *SingleLinkedList[int]
		Index            int
		ExpectedResult   int
		ExpectedError    error
	}{
		{
			Name:             "Inbounds index",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3}),
			Index:            1,
			ExpectedResult:   10,
			ExpectedError:    nil,
		},
		{
			Name:             "Inbounds index",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3, 5, 6, 7}),
			Index:            4,
			ExpectedResult:   10,
			ExpectedError:    nil,
		},
		{
			Name:             "Inbounds index",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3}),
			Index:            0,
			ExpectedResult:   10,
			ExpectedError:    nil,
		},
		{
			Name:             "Inbounds index",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3}),
			Index:            2,
			ExpectedResult:   3,
			ExpectedError:    nil,
		},
		{
			Name:             "Out of bounds index",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3}),
			Index:            3,
			ExpectedResult:   10,
			ExpectedError:    data_structures.ErrIndexOutOfBounds,
		},
		{
			Name:             "Negative index",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3}),
			Index:            -1,
			ExpectedResult:   10,
			ExpectedError:    data_structures.ErrIndexOutOfBounds,
		},
	}

	for _, test := range tcs {
		t.Run(test.Name, func(t *testing.T) {
			err := test.SingleLinkedList.Set(test.Index, test.ExpectedResult)
			if test.ExpectedError != nil {
				assert.EqualError(err, test.ExpectedError.Error(), test.Name)
				return
			}
			assert.NoError(err, test.Name)

			v, err := test.SingleLinkedList.Get(test.Index)
			assert.NoError(err, test.Name)
			assert.Equal(test.ExpectedResult, v, test.Name)
		})
	}
}

func TestAdd(t *testing.T) {

	assert := assert.New(t)
	tcs := []struct {
		Name             string
		SingleLinkedList *SingleLinkedList[int]
		TimesToAdd       int
		ValueToAdd       int
		ExpectedError    error
	}{
		{
			Name:             "Add one item",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3}),
			TimesToAdd:       1,
			ValueToAdd:       10,
			ExpectedError:    nil,
		},
		{
			Name:             "Add multiple items",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3, 5, 6, 7}),
			TimesToAdd:       3,
			ValueToAdd:       10,
			ExpectedError:    nil,
		},
		{
			Name:             "Add to empty array",
			SingleLinkedList: NewSingleLinkedList[int](),
			TimesToAdd:       1,
			ValueToAdd:       10,
			ExpectedError:    nil,
		},
		{
			Name:             "Add to empty array multiple items",
			SingleLinkedList: NewSingleLinkedList[int](),
			TimesToAdd:       3,
			ValueToAdd:       10,
			ExpectedError:    nil,
		},
	}

	for _, test := range tcs {
		t.Run(test.Name, func(t *testing.T) {
			for i := range test.TimesToAdd {
				test.SingleLinkedList.Add(test.ValueToAdd * (i + 1))
			}

			for i := range test.TimesToAdd {
				v, err := test.SingleLinkedList.Get(test.SingleLinkedList.Length() - test.TimesToAdd + i)
				assert.NoError(err, test.Name)
				assert.Equal(test.ValueToAdd*(i+1), v)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	assert := assert.New(t)
	tcs := []struct {
		Name             string
		SingleLinkedList *SingleLinkedList[int]
		Index            []int
		ExpectedArray    *SingleLinkedList[int]
		ExpectedSize     int
		ExpectedError    error
	}{
		{
			Name:             "Remove one item",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3}),
			Index:            []int{1},
			ExpectedArray:    NewSingleLinkedListFromSlice([]int{1, 3}),
			ExpectedSize:     2,
			ExpectedError:    nil,
		},
		{
			Name:             "Remove first item",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3}),
			Index:            []int{0},
			ExpectedArray:    NewSingleLinkedListFromSlice([]int{2, 3}),
			ExpectedSize:     2,
			ExpectedError:    nil,
		},
		{
			Name:             "Remove last item",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3}),
			Index:            []int{2},
			ExpectedArray:    NewSingleLinkedListFromSlice([]int{1, 2}),
			ExpectedSize:     2,
			ExpectedError:    nil,
		},
		{
			Name:             "Remove multiple items",
			SingleLinkedList: NewSingleLinkedListFromSlice([]int{1, 2, 3, 5, 6, 7}),
			Index:            []int{3, 1},
			ExpectedArray:    NewSingleLinkedListFromSlice([]int{1, 3, 6, 7}),
			ExpectedSize:     4,
			ExpectedError:    nil,
		},
		{
			Name:             "Remove from empty array",
			SingleLinkedList: NewSingleLinkedList[int](),
			Index:            []int{1},
			ExpectedArray:    NewSingleLinkedList[int](),
			ExpectedSize:     0,
			ExpectedError:    data_structures.ErrIndexOutOfBounds,
		},
		{
			Name:             "Remove multiple items from empty error",
			SingleLinkedList: NewSingleLinkedList[int](),
			Index:            []int{3, 1},
			ExpectedArray:    NewSingleLinkedList[int](),
			ExpectedSize:     0,
			ExpectedError:    data_structures.ErrIndexOutOfBounds,
		},
	}

	for _, test := range tcs {
		t.Run(test.Name, func(t *testing.T) {
			for _, idx := range test.Index {
				expected_v := 0
				if test.SingleLinkedList.Length() > 0 {
					_expected_v, err := test.SingleLinkedList.Get(idx)
					expected_v = _expected_v
					assert.NoError(err, test.Name)

				}
				v, err := test.SingleLinkedList.Remove(idx)
				if test.ExpectedError != nil {
					assert.EqualError(err, test.ExpectedError.Error(), test.Name)
				} else {
					assert.Equal(expected_v, v, test.Name)
				}

			}
			for i := range test.ExpectedArray.Length() {
				expected_v, err := test.ExpectedArray.Get(i)
				assert.NoError(err, test.Name)
				v, err := test.SingleLinkedList.Get(i)
				assert.NoError(err, test.Name)
				assert.Equal(expected_v, v, test.Name)
			}

			assert.Equal(test.ExpectedSize, test.SingleLinkedList.Length(), test.Name)

		})
	}
}
