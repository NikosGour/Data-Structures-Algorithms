package algorithms

import "cmp"

func BubbleSort[T cmp.Ordered](arr []T) []T {
	for i := 0; i < len(arr); i++ {
		for j := 0; j+i+1 < len(arr); j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}
