package main

import (
	"fmt"
	"strings"
)

func isRotation(s1, s2 string) bool {
	return len(s1) == len(s2) && strings.Index(s2+s2, s1) > 0
}

func main() {
	s1 := "waterbottle"
	s2 := "erbottlewat"
	fmt.Printf("%q is rotation of %q == %v\n", s2, s1, isRotation(s1, s2))
}
