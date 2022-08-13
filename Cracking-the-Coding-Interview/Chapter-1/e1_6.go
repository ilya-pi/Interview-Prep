package main

import "fmt"

func compress(s string) string {
	/*
	   Go char by char, maintaining current char and and the count, on the new char - change; in the end compare lengths and return the shorter or original if equal
	*/

	if len(s) < 3 {
		return s
	}

	compressRune := func(r rune, count int) string {
		return fmt.Sprintf("%c%d", r, count)
	}

	arr := []rune(s)

	prev := arr[0]
	count := 1
	var res string
	for i := 1; i < len(arr); i++ {
		switch {
		case arr[i] == prev:
			count++
		case arr[i] != prev:
			res += compressRune(prev, count)
			prev = arr[i]
			count = 1
		}
	}
	res += compressRune(prev, count)
	return res
}

func main() {
	s := "aabcccccaaa"
	fmt.Printf("compressed(%s) == %s\n", s, compress(s))
}
