package main

import "fmt"

func allUnique(s string) bool {
	// Use hashmap to track characters we already met
	// If cannot use - then compare every character to every — O(n^2)
	chars := map[rune]bool{}
	for _, v := range s {
		if _, ok := chars[v]; ok {
			return false
		}
		chars[v] = true
	}
	return true
}

func main() {
	s := "aslkda"
	fmt.Printf("String %s has all unique characters == %v\n", s, allUnique(s))
}
