package main

import "fmt"

func oneEditAway(s1, s2 string) bool {
	/*
	   Approach: insert or remove is the same, but the longer string is different
	   replace is one different character away

	   Split those cases and check them separately
	*/

	switch {
	case len(s1) == len(s2):
		var diff int
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				diff++
				if diff > 1 {
					return false
				}
			}
		}
		return true
	default:
		// make s1 the longer string
		if len(s1) < len(s2) {
			s1, s2 = s2, s1
		}
		// compare rune by rune with skipping 1 rune
		var d int
		for i := 0; i < len(s2); i++ {
			if s1[i+d] != s2[i] {
				if s1[i+1] == s2[i] && d == 0 {
					d = 1
				} else {
					return false
				}
			}
		}
		return true
	}
}

func main() {
	// 18:06
	// 18:15
	s1 := "caer"
	s2 := "cat"
	fmt.Printf("%v\n", oneEditAway(s1, s2))
}
