package algorithms

import (
	"cmp"
)

func InsertionSort[T cmp.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		if key >= arr[i-1] {
			continue
		}

		for j := i; j >= 1; j-- {
			if key > arr[j-1] {
				break
			}
			if arr[j] < arr[j-1] {
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}
