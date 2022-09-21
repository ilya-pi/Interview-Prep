package main

import "fmt"

func main() {
	// 17:10
	// 17:07 brute force solution
	// O(n)

	/*
		1
		11
		21
		1211
		111221
		312211
		13112221
		1113213211
		31131211131221
		13211311123113112211

	*/
	// Approach: it seems to be a solution of maintaining a state
	// We'll have a seed array
	// On each level we will run a loop to walk the state through seed sequence creating a new seed sequence
	seed := []int{2, 2}
	for n := 0; n < 20; n++ {
		current := seed[0]
		count := 1
		var ans []int
		for i := 1; i < len(seed); i++ {
			if seed[i] != current {
				ans = append(ans, count)
				ans = append(ans, current)
				current = seed[i]
				count = 1
			} else {
				count++
			}
		}
		// Add last seen
		ans = append(ans, count)
		ans = append(ans, current)
		// output and update seed
		fmt.Printf("%d's is %v\n", n, ans)
		seed = ans
	}
}
