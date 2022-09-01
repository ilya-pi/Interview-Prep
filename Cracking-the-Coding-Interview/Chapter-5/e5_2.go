package main

import "fmt"

func binaryToString(v float64) string {
	/*
		Approach:

		If I have x that is < x < 1, the leading bit will be 1 if I multiply it by
		2 as that is a shift of value left in binary. That way I can construct the while number
	*/

	res := "."
	for {
		if v < 0.000001 || len(res) >= 32 {
			break
		}
		v *= 2
		if v >= 1.0 {
			res += "1"
			v -= 1
		} else {
			res += "0"
		}
	}
	return res
}

func main() {
	v := 0.72
	fmt.Printf("Binary of %v is %v\n", v, binaryToString(v))
}
