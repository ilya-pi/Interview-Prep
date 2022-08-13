package main

import "fmt"

func isPalindromePermutation(s string) bool {
	/*
	   Keep counting all elements in a hashmap
	   and then removing them, if at the end len
	   of hashmap is 0 or 1 then it is a palindrome
	   permutation
	*/
	elems := map[rune]int{}

	for _, v := range s {
		if _, ok := elems[v]; ok {
			elems[v]--
			if elems[v] == 0 {
				delete(elems, v)
			}
		} else {
			elems[v]++
		}
	}

	return len(elems) == 1 || len(elems) == 0

}

func main() {
	s := "abab"
	fmt.Printf("%s is a palindrome == %v\n ", s, isPalindromePermutation(s))
}
