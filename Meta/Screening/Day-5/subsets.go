package main

import "fmt"

func subsets(arr []int) [][]int {
	/*
	   Approach:

	   Recusrively take and not take every element with backtracking, if element in question is > len -> add it

	*/

	var ans [][]int

	var subst func([]int, []int, int)
	subst = func(acc []int, src []int, ind int) {
		if ind == len(src) {
			ans = append(ans, acc)
			return
		}

		// not take ind
		subst(acc, src, ind+1)
		// takeind
		acc = append(acc, src[ind])
		subst(acc, src, ind+1)
	}
	// call the inner function
	subst(nil, arr, 0)

	return ans
	// O(2^n) - n levels with two branches on each
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	for _, v := range subsets(arr) {
		fmt.Printf("%v\n", v)
	}
}
