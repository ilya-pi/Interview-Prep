package main

import "fmt"

func oneAway(s1, s2 string) bool {
	/*
		Split into either replace scenario or one extra char and see if comparison can be completed in one diff
	*/
	switch {
	case len(s1) != len(s2):
		if len(s2) > len(s1) {
			s1, s2 = s2, s1
		}
		if len(s1) != len(s2)+1 {
			return false
		}
		var changes int
		var i, j int
		for i < len(s1) && j < len(s2) {
			if s1[i] != s2[j] {
				changes++
				if changes > 1 {
					return false
				}
				i++
			} else {
				i++
				j++
			}
		}
	default: // 1 replacement
		var changes int
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				changes++
				if changes > 1 {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	s1 := "pale"
	s2 := "bake"
	fmt.Printf("%s and %s are one away == %v\n", s1, s2, oneAway(s1, s2))
}
