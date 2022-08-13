package main

import (
	"fmt"
)

func urlify(s string) string {
	/*
		parts := strings.Split(s, " ")
		return strings.Join(parts, "%20")
	*/

	parts := []string{}
	var prev int
	for i, v := range s {
		if v == ' ' {
			parts = append(parts, s[prev:i])
			prev = i + 1
		}
	}
	if prev+1 < len(s) {
		parts = append(parts, s[prev:])
	}

	var res string
	for i := 0; i < len(parts)-1; i++ {
		res += parts[i]
		res += "%20"
	}
	res += parts[len(parts)-1]
	return res
}

func main() {
	fmt.Printf("%s\n", urlify("Mr John Smith"))
}
