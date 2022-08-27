package main

import (
	"fmt"
	"math"
)

type Tree5 struct {
	v           int
	left, right *Tree5
}

func (t *Tree5) isValidBST(min, max int) bool {
	/*
	   If it is a valid BST then it is a valid BST for all nodes,
	   so we'll use recurse to check it is valid while maintaining correct intervals
	*/

	return t.v > min && t.v <= max &&
		(t.left == nil || t.left.isValidBST(min, t.v)) &&
		(t.right == nil || t.right.isValidBST(t.v, max))
}

func main() {
	t := &Tree5{v: 5}
	t1 := &Tree5{v: 3}
	t2 := &Tree5{v: 7}
	t3 := &Tree5{v: 2}
	t4 := &Tree5{v: 8}
	t.left, t.right = t1, t2
	t1.left, t1.right = t3, t4
	fmt.Printf("isValidBST == %v\n", t.isValidBST(math.MinInt, math.MaxInt))
}
