package main

import "fmt"

func isColourful(n int) bool {
	/*
	   1. Split into array of digits

	   2. For every pair O(n^2) calculate product

	   3. See if we have it already, otherwise bail
	*/

	// 1 Split into array of digits
	var digits []int
	for n != 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	products := make(map[int]bool)
	// 2 For every pair O(n^2)
	for i := 0; i < len(digits); i++ {

		product := 1
		for j := i; j < len(digits); j++ {
			product *= digits[j]
			if _, ok := products[product]; ok {
				return false
			} else {
				products[product] = true
			}
		}
	}
	return true
	// O(log^2n)
}

func main() {
	fmt.Printf("%v\n", isColourful(3245))
	fmt.Printf("%v\n", isColourful(326))
}
