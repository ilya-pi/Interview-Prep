package main

import "fmt"

func nextNumber(v uint32) uint32 {
	/*
	   1001101110110111

	   To keep the same amount of 1, we need to find first 0
	   that we can flip to 1, but we will need to flip another
	   bit to 0 then of those already seen 1s

	   1000

	   For that we need to find the first sequence of 1-s that we can move left 1 bit

	   1) find position of 1st 1
	   2) find position of 1st 0 after first 1
	   3) move all ones met before that to right
	*/

	// 1 Find position of 1st 1
	i := 0
	for ; i < 32 && (v&(uint32(1)<<i) == 0); i++ {
	}
	// 2 Find position of 1st 0 after 1
	j := i
	for ; j < 32 && (v&(uint32(1)<<j) > 0); j++ {
	}
	// 3 Find amount of 1s we have before that number
	i = j - 1
	ones := 0
	for ; i >= 0; i-- {
		if v&(uint32(1)<<i) > 0 {
			ones++
		}
	}
	fmt.Printf("j == %d, ones == %d\n", j, ones)
	// Flip those bits
	r := set1(v, j)
	ones-- // set 1 one
	for i = 0; i < j; i++ {
		if ones > 0 { // We already set 1 one forward
			r = set1(r, i)
			ones--
		} else {
			r = set0(r, i)
		}
	}
	return r
}

func set1(v uint32, i int) uint32 {
	return v | (uint32(1) << i)
}

func set0(v uint32, i int) uint32 {
	return v & ^(uint32(1) << i)
}

func prevNumber(v uint32) uint32 {
	// Edgecase are pretty wild here and require some thinking
	if v < 3 {
		return 0
	}
	/*
	   To find prev number we need to find
	   the first 1 after 0 and then shift it right,
	   otherwise the "same amount of 0" number won't be
	   smaller or otherwise it won't have the same amount of 1's
	*/
	// Find first 0
	i := 0
	for ; i < 32 && (v&(uint32(1)<<i) > 0); i++ {
	}
	// Find first 1 after first 0 after ones
	j := i
	for ; j < 32 && (v&(uint32(1)<<j) == 0); j++ {
	}
	fmt.Printf("i == %d, j == %d\n", i, j)
	r := set0(v, j)
	r = set1(r, j-1)

	return r
}

func main() {
	n := uint32(876423)
	fmt.Printf("Next to \n%b\nis\n%b\n\n", n, nextNumber(n))
	fmt.Printf("Prev to \n%b\nis\n%b\n\n", n, prevNumber(n))
}
