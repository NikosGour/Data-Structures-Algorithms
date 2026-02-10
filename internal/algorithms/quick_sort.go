package algorithms

import (
	"cmp"
)

func QuickSort[T cmp.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	i := -1
	pivot := arr[len(arr)-1]
	j := 0
	for j = 0; j < len(arr)-1; j++ {
		if arr[j] < pivot {
			i++
			temp := arr[j]
			arr[j] = arr[i]
			arr[i] = temp
		}
	}

	i++
	temp := arr[j]
	arr[j] = arr[i]
	arr[i] = temp
	if i > 0 {
		QuickSort(arr[0:i])
	}
	if i < len(arr)-1 {
		QuickSort(arr[i+1:])
	}

	return arr
}
