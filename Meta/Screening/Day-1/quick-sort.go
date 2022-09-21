package main

import "fmt"

func quickSort(arr []int) {

	partition := func(arr []int, l, r int) int {
		pivot := arr[r]
		i := l - 1
		for j := l; j < r; j++ {
			if arr[j] < pivot {
				i++
				if i != j {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
		arr[i+1], arr[r] = arr[r], arr[i+1]
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
	arr := []int{1, 4, 2, 51, 56, 12, 23}
	fmt.Printf("Before \n%v\n", arr)
	quickSort(arr)

	fmt.Printf("After \n%v\n", arr)
}
