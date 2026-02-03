package algorithms

import "cmp"

func SelectionSort[T cmp.Ordered](arr []T) []T {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		tmp := arr[min]
		arr[min] = arr[i]
		arr[i] = tmp
	}

	return arr
}
