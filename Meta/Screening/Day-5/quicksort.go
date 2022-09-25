package main

import "fmt"

func quicksort(arr []int) {
	/*

	   exit case
	   partition
	   qs left qs right

	*/

	partition := func(arr []int, left, right int) int {
		pivot := arr[right]
		i := left - 1
		for j := left; j < right; j++ {
			if arr[j] < pivot {
				i++
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
		// partition the pivot element
		arr[right], arr[i+1] = arr[i+1], arr[right]
		return i + 1
	}

	var qs func([]int, int, int)
	qs = func(arr []int, left, right int) {
		if left >= right {
			return
		}
		pi := partition(arr, left, right)
		qs(arr, left, pi-1)
		qs(arr, pi+1, right)
	}
	qs(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{32, 15, 6, 34, 293, 1, 4, 5, 2}
	fmt.Printf("Before: \n%v\n", arr)
	quicksort(arr)
	fmt.Printf("After: \n%v\n", arr)
}
