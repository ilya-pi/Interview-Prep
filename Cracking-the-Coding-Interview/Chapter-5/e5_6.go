package main

import "fmt"

func bitsToConvert(a, b uint32) uint32 {
	/*
	   Approach:

	   1) xor a and b, we will get only the "different" bit positions
	   2) count amount of 1s in the result
	*/

	diff := a ^ b

	var r uint32
	for i := 0; i < 32; i++ {
		if diff&(uint32(1)<<i) > 0 {
			r++
		}
	}
	return r
}

func main() {
	a := uint32(3178)
	b := uint32(9832)
	fmt.Printf("Need %d bit flips to convert \n%b to \n%b\n", bitsToConvert(a, b), a, b)
}
