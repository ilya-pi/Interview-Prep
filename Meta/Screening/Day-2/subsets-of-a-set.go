package main

import "fmt"

func subsets(arr []int) [][]int {

	var ans [][]int
	var subs func([]int, []int, int)
	subs = func(arr []int, curr []int, ind int) {
		// record and exit
		if ind == len(arr) {
			ans = append(ans, curr)
			return
		}

		// Add current element
		curr = append(curr, arr[ind])
		subs(arr, curr, ind+1)
		// Backtrack adding current element
		curr = curr[:len(curr)-1]
		subs(arr, curr, ind+1)
	}
	subs(arr, nil, 0)
	return ans
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("Subsets of %v:\n", arr)
	for _, v := range subsets(arr) {
		fmt.Printf("%v\n", v)
	}
}
