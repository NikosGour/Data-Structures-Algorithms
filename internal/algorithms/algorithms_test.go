package algorithms

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		desc        string
		startingArr []int
		sortedArr   []int
	}{
		{
			desc:        "Simple array",
			startingArr: []int{3, 10, 1, 2, 14, 9},
			sortedArr:   []int{1, 2, 3, 9, 10, 14},
		},
		{
			desc:        "Duplicate numbers",
			startingArr: []int{3, 10, 2, 2, 14, 9},
			sortedArr:   []int{2, 2, 3, 9, 10, 14},
		},
		{
			desc:        "Empty array",
			startingArr: []int{},
			sortedArr:   []int{},
		},
		{
			desc:        "Sorted array",
			startingArr: []int{1, 2, 3, 4, 5, 6, 7, 8},
			sortedArr:   []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			desc:        "Last number wrong array",
			startingArr: []int{1, 2, 3, 4, 5, 6, 7, 0},
			sortedArr:   []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			desc:        "First number wrong array",
			startingArr: []int{9, 2, 3, 4, 5, 6, 7, 0},
			sortedArr:   []int{0, 2, 3, 4, 5, 6, 7, 9},
		},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			sorted_arr := BubbleSort(test.startingArr)
			is_sorted := slices.Equal(sorted_arr, test.sortedArr)
			assert.True(is_sorted)
		})
	}
}
