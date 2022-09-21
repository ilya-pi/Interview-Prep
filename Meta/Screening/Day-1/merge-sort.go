package main

import "fmt"

func mergeSort(arr []int) []int {
	/*
	   Approach:
	   Split in two halves, sort them, then weave
	*/

	if len(arr) == 1 || len(arr) == 0 {
		return arr
	}

	p := len(arr) / 2
	a := mergeSort(arr[:p])
	b := mergeSort(arr[p:])

	merge := func(a, b []int) []int {
		ans := make([]int, len(a)+len(b))
		i := 0
		ja, jb := 0, 0
		for ja < len(a) && jb < len(b) {
			if a[ja] < b[jb] {
				ans[i] = a[ja]
				ja++
			} else {
				ans[i] = b[jb]
				jb++
			}
			i++
		}

		// drain the rest of a and b
		for ; ja < len(a); ja++ {
			ans[i] = a[ja]
			i++
		}

		for ; jb < len(b); jb++ {
			ans[i] = b[jb]
			i++
		}
		return ans
	}

	ans := merge(a, b)
	return ans
}

func main() {
	arr := []int{2, 61, 23, 63, 75, 23}
	fmt.Printf("Before: \n%v\n", arr)
	arr = mergeSort(arr)
	fmt.Printf("After: \n%v\n", arr)
}
