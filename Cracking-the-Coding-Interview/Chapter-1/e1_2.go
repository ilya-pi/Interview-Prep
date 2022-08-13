package main

import (
	"fmt"
)

func isPermutation(s1, s2 string) bool {
	// 1. Record all characters of s1 in map with the amount of times we meet them and then count down deleting the records when it is at 0 — O(n) + O(n) space
	chars := map[rune]int{}
	for _, v := range s1 {
		chars[v]++
	}

	for _, v := range s2 {
		if _, ok := chars[v]; !ok {
			return false
		}
		chars[v]--
		if chars[v] == 0 {
			delete(chars, v)
		}
	}

	return len(chars) == 0

	// 2. Sort both strings and just compare them - O(nlogn) + O(1) space
	/*
		s1r := []rune(s1)
		s2r := []rune(s2)
		sort.Slice(s1r, func(i, j int) bool { return s1r[i] < s1r[j] })
		sort.Slice(s2r, func(i, j int) bool { return s2r[i] < s2r[j] })

		return string(s1r) == string(s2r)
	*/
}

func main() {
	s1 := "abcdef"
	s2 := "fedacb"
	fmt.Printf("Strings %s and %s are a permutations of each other == %v\n", s1, s2, isPermutation(s1, s2))
}
