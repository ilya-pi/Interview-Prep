package main

import "fmt"

func pairwiseSwap(a uint32) uint32 {
	/*
	   We want to mask both and then shift and & them


	   A in binary:

	   1010
	*/

	mask := uint32(0xAAAAAAAA)
	fmt.Printf("Mask is \n%b\n", mask)
	a1 := (a & mask) >> 1
	fmt.Printf("a1 is \n%b\n", a1)
	a2 := (a & (mask >> 1)) << 1
	fmt.Printf("a2 is\n%b\n", a2)
	merged := a1 | a2
	fmt.Printf("Merged is \n%b\n", merged)
	return merged
}

func main() {
	a := uint32(1022)
	fmt.Printf("Pairwise swap\n%b = \n%b\n", a, pairwiseSwap(a))
}
