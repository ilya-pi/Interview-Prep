package main

import "fmt"

func quickSort(arr []int) {

	/*
		Find pi, quick sort the parts
	*/

	partition := func(arr []int, left, right int) int {
		pivot := arr[right]

		i := left - 1                   // last placed element
		for j := left; j < right; j++ { // run for all elements except the pivot element
			if arr[j] < pivot {
				i++ // next place of the element
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		// place pivot element
		arr[i+1], arr[right] = arr[right], arr[i+1]
		return i + 1
	}

	var qs func([]int, int, int)
	qs = func(arr []int, left, right int) {
		// exit case
		if left >= right {
			return
		}
		pi := partition(arr, left, right)
		// element pi should be in place now
		qs(arr, left, pi-1)
		qs(arr, pi+1, right)
	}

	qs(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{32, 15, 65, 12, 64, 23}
	fmt.Printf("Before: \n%v\n", arr)
	quickSort(arr)
	fmt.Printf("After: \n%v\n", arr)
}
