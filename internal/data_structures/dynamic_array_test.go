package data_structures_test

import (
	"dsa/internal/data_structures"
	"testing"

	"github.com/NikosGour/logging/log"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	log.Debug("nikos")
	assert := assert.New(t)
	tcs := []struct {
		Name           string
		DynamicArray   *data_structures.DynamicArray[int]
		Index          int
		ExpectedResult int
		ExpectedError  error
	}{
		{
			Name:           "Inbounds index",
			DynamicArray:   data_structures.NewDynamicArrayFromSlice([]int{1, 2, 3}),
			Index:          1,
			ExpectedResult: 2,
			ExpectedError:  nil,
		},
		{
			Name:           "Inbounds index",
			DynamicArray:   data_structures.NewDynamicArrayFromSlice([]int{1, 2, 3, 5, 6, 7}),
			Index:          4,
			ExpectedResult: 6,
			ExpectedError:  nil,
		},
		{
			Name:           "Inbounds index",
			DynamicArray:   data_structures.NewDynamicArrayFromSlice([]int{1, 2, 3}),
			Index:          0,
			ExpectedResult: 1,
			ExpectedError:  nil,
		},
		{
			Name:           "Inbounds index",
			DynamicArray:   data_structures.NewDynamicArrayFromSlice([]int{1, 2, 3}),
			Index:          2,
			ExpectedResult: 3,
			ExpectedError:  nil,
		},
		{
			Name:           "Out of bounds index",
			DynamicArray:   data_structures.NewDynamicArrayFromSlice([]int{1, 2, 3}),
			Index:          3,
			ExpectedResult: 0,
			ExpectedError:  data_structures.ErrIndexOutOfBounds,
		},
		{
			Name:           "Negative index",
			DynamicArray:   data_structures.NewDynamicArrayFromSlice([]int{1, 2, 3}),
			Index:          -1,
			ExpectedResult: 0,
			ExpectedError:  data_structures.ErrIndexOutOfBounds,
		},
	}

	for _, test := range tcs {
		t.Run(test.Name, func(t *testing.T) {
			v, err := test.DynamicArray.Get(test.Index)
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
	log.Debug("nikos")
	assert := assert.New(t)
	tcs := []struct {
		Name           string
		DynamicArray   *data_structures.DynamicArray[int]
		Index          int
		ExpectedResult int
		ExpectedError  error
	}{
		{
			Name:           "Inbounds index",
			DynamicArray:   data_structures.NewDynamicArrayFromSlice([]int{1, 2, 3}),
			Index:          1,
			ExpectedResult: 10,
			ExpectedError:  nil,
		},
		{
			Name:           "Inbounds index",
			DynamicArray:   data_structures.NewDynamicArrayFromSlice([]int{1, 2, 3, 5, 6, 7}),
			Index:          4,
			ExpectedResult: 10,
			ExpectedError:  nil,
		},
		{
			Name:           "Inbounds index",
			DynamicArray:   data_structures.NewDynamicArrayFromSlice([]int{1, 2, 3}),
			Index:          0,
			ExpectedResult: 10,
			ExpectedError:  nil,
		},
		{
			Name:           "Inbounds index",
			DynamicArray:   data_structures.NewDynamicArrayFromSlice([]int{1, 2, 3}),
			Index:          2,
			ExpectedResult: 3,
			ExpectedError:  nil,
		},
		{
			Name:           "Out of bounds index",
			DynamicArray:   data_structures.NewDynamicArrayFromSlice([]int{1, 2, 3}),
			Index:          3,
			ExpectedResult: 10,
			ExpectedError:  data_structures.ErrIndexOutOfBounds,
		},
		{
			Name:           "Negative index",
			DynamicArray:   data_structures.NewDynamicArrayFromSlice([]int{1, 2, 3}),
			Index:          -1,
			ExpectedResult: 10,
			ExpectedError:  data_structures.ErrIndexOutOfBounds,
		},
	}

	for _, test := range tcs {
		t.Run(test.Name, func(t *testing.T) {
			err := test.DynamicArray.Set(test.Index, test.ExpectedResult)
			if test.ExpectedError != nil {
				assert.EqualError(err, test.ExpectedError.Error(), test.Name)
				return
			}
			assert.NoError(err, test.Name)

			v, err := test.DynamicArray.Get(test.Index)
			assert.NoError(err, test.Name)
			assert.Equal(test.ExpectedResult, v, test.Name)
		})
	}
}

func TestAdd(t *testing.T) {
	log.Debug("nikos")
	assert := assert.New(t)
	tcs := []struct {
		Name          string
		DynamicArray  *data_structures.DynamicArray[int]
		TimesToAdd    int
		ValueToAdd    int
		ExpectedError error
	}{
		{
			Name:          "Add one item",
			DynamicArray:  data_structures.NewDynamicArrayFromSlice([]int{1, 2, 3}),
			TimesToAdd:    1,
			ValueToAdd:    10,
			ExpectedError: nil,
		},
		{
			Name:          "Add multiple items",
			DynamicArray:  data_structures.NewDynamicArrayFromSlice([]int{1, 2, 3, 5, 6, 7}),
			TimesToAdd:    3,
			ValueToAdd:    10,
			ExpectedError: nil,
		},
		{
			Name:          "Add to empty array",
			DynamicArray:  data_structures.NewDynamicArray[int](),
			TimesToAdd:    1,
			ValueToAdd:    10,
			ExpectedError: nil,
		},
		{
			Name:          "Add to empty array multiple items",
			DynamicArray:  data_structures.NewDynamicArray[int](),
			TimesToAdd:    3,
			ValueToAdd:    10,
			ExpectedError: nil,
		},
	}

	for _, test := range tcs {
		t.Run(test.Name, func(t *testing.T) {
			for i := range test.TimesToAdd {
				err := test.DynamicArray.Add(test.ValueToAdd * (i + 1))
				assert.NoError(err, test.Name)
			}

			for i := range test.TimesToAdd {
				v, err := test.DynamicArray.Get(test.DynamicArray.Length() - test.TimesToAdd + i)
				assert.NoError(err, test.Name)
				assert.Equal(test.ValueToAdd*(i+1), v)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	log.Debug("nikos")
	assert := assert.New(t)
	tcs := []struct {
		Name          string
		DynamicArray  *data_structures.DynamicArray[int]
		Index         []int
		ExpectedArray *data_structures.DynamicArray[int]
		ExpectedSize  int
		ExpectedError error
	}{
		{
			Name:          "Remove one item",
			DynamicArray:  data_structures.NewDynamicArrayFromSlice([]int{1, 2, 3}),
			Index:         []int{1},
			ExpectedArray: data_structures.NewDynamicArrayFromSlice([]int{1, 3}),
			ExpectedSize:  2,
			ExpectedError: nil,
		},
		{
			Name:          "Remove multiple items",
			DynamicArray:  data_structures.NewDynamicArrayFromSlice([]int{1, 2, 3, 5, 6, 7}),
			Index:         []int{3, 1},
			ExpectedArray: data_structures.NewDynamicArrayFromSlice([]int{1, 3, 6, 7}),
			ExpectedSize:  4,
			ExpectedError: nil,
		},
		{
			Name:          "Remove from empty array",
			DynamicArray:  data_structures.NewDynamicArray[int](),
			Index:         []int{1},
			ExpectedArray: data_structures.NewDynamicArray[int](),
			ExpectedSize:  0,
			ExpectedError: data_structures.ErrIndexOutOfBounds,
		},
		{
			Name:          "Remove multiple items from empty error",
			DynamicArray:  data_structures.NewDynamicArray[int](),
			Index:         []int{3, 1},
			ExpectedArray: data_structures.NewDynamicArray[int](),
			ExpectedSize:  0,
			ExpectedError: data_structures.ErrIndexOutOfBounds,
		},
	}

	for _, test := range tcs {
		t.Run(test.Name, func(t *testing.T) {
			for _, idx := range test.Index {
				expected_v := 0
				if test.DynamicArray.Length() > 0 {
					_expected_v, err := test.DynamicArray.Get(idx)
					expected_v = _expected_v
					assert.NoError(err, test.Name)

				}
				v, err := test.DynamicArray.Remove(idx)
				if test.ExpectedError != nil {
					assert.EqualError(err, test.ExpectedError.Error(), test.Name)
				} else {
					assert.Equal(expected_v, v, test.Name)
				}

			}
			for i := range test.ExpectedArray.Length() {
				expected_v, err := test.ExpectedArray.Get(i)
				assert.NoError(err, test.Name)
				v, err := test.DynamicArray.Get(i)
				assert.NoError(err, test.Name)
				assert.Equal(expected_v, v, test.Name)
			}

			assert.Equal(test.ExpectedSize, test.DynamicArray.Length(), test.Name)

		})
	}
}
