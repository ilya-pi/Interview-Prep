package main

import "fmt"

func setBit(n int32, i int) int32 {
	return n | 1<<i
}

func fitBit(i, j int, m, n int32) int32 {
	/*
		  i   j
		00000000000
		00000000000

		mask
		  i   j
		00111110000

	*/
	var mask int32
	for k := j; k <= i; k++ {
		mask = setBit(mask, k)
	}
	fmt.Printf("mask\n%0b\n", mask)

	//mask := (-1 >>>j)^(-1>>>i)
	// empty mask bits to 0
	fmt.Printf("n is \n%b\n", n)
	n = n & ^mask
	fmt.Printf("Emptied mask bits\n%b\n", n)
	fmt.Printf("m \n%b\n", m)
	fmt.Printf("^mask \n%b\n", ^mask)
	m = (m << j) | ^mask
	fmt.Printf("Masked m\n%b\n", m)
	res := n & m
	fmt.Printf("\n\n\nn %b\nm %b\nr %b\n\n\n", n, m, res)
	return n & m
}

func main() {
	n := int32(12321233)
	m := int32(123)
	i := 20
	j := 10
	fmt.Printf("--\n%b\n%b\nfitbit %d to %d\n%b\n", n, m, i, j, fitBit(i, j, m, n))
}
