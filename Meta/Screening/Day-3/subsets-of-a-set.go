package main

import "fmt"

func subsets(arr []int) [][]int {
	/*
	   Recursive (accurate)

	   Exit -> empty
	*/
	if len(arr) == 0 {
		return [][]int{{}}
	}

	a := arr[0]
	prev := subsets(arr[1:])
	ans := make([][]int, len(prev)*2)
	copy(ans, prev)
	copy(ans[len(prev):], prev)
	for i := len(prev); i < len(ans); i++ {
		ans[i] = append(ans[i], a)
	}

	// O(2^n-1 + 2^n-2 + .. ) == O(2^n)
	// space - 2^n-1 * n + 2^n * n = 2^n * n
	return ans
}

func subsets2(arr []int) [][]int {
	/*
	   Recursive with backtracking


	   On each step, we look at next element and create subset with and without it
	*/

	var acc [][]int

	var subs func([]int, []int, int)
	subs = func(r []int, arr []int, ind int) {
		if ind == len(arr) {
			acc = append(acc, r)
			return
		}

		// not take it
		subs(r, arr, ind+1)
		// take it
		r = append(r, arr[ind])
		subs(r, arr, ind+1)
	}
	// O(2^n) :-/
	// space - 2^n * n

	subs(nil, arr, 0)
	return acc

}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	arr = []int{1, 2, 3, 4, 5}
	for _, v := range subsets(arr) {
		fmt.Printf("%v\n", v)
	}
	fmt.Printf("----\n")
	for _, v := range subsets2(arr) {
		fmt.Printf("%v\n", v)
	}
}
