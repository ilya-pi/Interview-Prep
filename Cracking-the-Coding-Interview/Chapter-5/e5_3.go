package main

import "fmt"

func bitToWin(v int32) int32 {
	/*

	   Approach:

	   Left to right
	     Count amount of 1s
	     If bit is 0 -> flip it and continue counting move "start" to that bit and record that was flipped
	     If bit is 0 and flipped, make the new sequence longer from "new" 1

	*/

	lastFlipped := int32(-1)
	longest := int32(0)
	length := int32(0)
	flipped := false
	for i := int32(31); i >= 0; i-- {
		c := int32(1 << i)
		is1 := c&v > 0
		fmt.Printf("i: %v %v length: %v longest: %v\n", i, is1, length, longest)
		if is1 {
			length++
			continue
		}
		if !is1 && !flipped {
			flipped = true
			lastFlipped = i
		} else if !is1 && flipped {
			if length > longest {
				longest = length
			}
			flipped = false
			start := lastFlipped - 1
			length = start - i
		}
	}
	if length > longest {
		return length
	}

	return longest
}

func main() {
	v := int32(1234295)
	fmt.Printf("Longest 1 of %b is %v\n", v, bitToWin(v))
}
