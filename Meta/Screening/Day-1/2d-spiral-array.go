package main

import "fmt"

func spiral(n int) [][]int {
	/*

	   Approach: we know the end array's size and the pattern how we change the index direction.
	   Hence we could just fill it in in a while loop as long as there is anything to fill in.

	   We will need:

	   1. directions that we can iterate through, ideally with +1%4
	   2. for loop to fill it in

	*/

	dir := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	ans := make([][]int, n)
	for i, _ := range ans {
		ans[i] = make([]int, n)
	}

	// We have a field nxn and start indices i and j
	// Now we run a for loop to fill it in, until there
	// is no more space after direction switch
	cd := 0
	i, j := 0, 0
	v := 0
	for {
		v++
		ans[i][j] = v
		dx, dy := dir[cd][0], dir[cd][1]

		i, j = i+dx, j+dy
		if i < 0 || i > n-1 || j < 0 || j > n-1 || ans[i][j] != 0 {
			i, j = i-dx, j-dy
			cd = (cd + 1) % 4
			// now to see if we reached
			// the end in the center
			dx, dy = dir[cd][0], dir[cd][1]
			i, j = i+dx, j+dy
			if ans[i][j] != 0 {
				break
			}
		}
	}
	return ans
}

func main() {
	// 16:00
	// 16:25 working version
	// Other solutions could be to sit down and write a formula
	// of ans[i][j] - seems feasible, though may be partially annoying

	// Complexity: O(n^2), both space and time
	n := 20
	for _, v := range spiral(n) {
		fmt.Printf("%v\n", v)
	}
}
