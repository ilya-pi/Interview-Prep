package main

import "fmt"

type Tree10 struct {
	v           int
	left, right *Tree10
}

func isSubSlice(t1 []*int, t2 []*int) bool {

	k := 0
	for _, v := range t1 {
		if (t2[k] == nil && v == nil) || (t2[k] != nil && v != nil && *t2[k] == *v) {
			k++
		} else {
			k = 0
		}
		if k == len(t2) {
			return true
		}
	}

	return false
}

func (t *Tree10) isSubtree(t2 *Tree10) bool {
	/*
	   Approach: I believe if we have a regular
	   in-order representation of the tree (with
	   nil values!) then it shows that one is subtree of the other

	   1. Get slice rep of in-order t2
	   2. Start in order of the second
	   tree while trying to match to the
	   slice from 1 with a pointer, with backtracking

	*/

	var inOrder func(*Tree10, *[]*int)
	inOrder = func(n *Tree10, r *[]*int) {
		if n == nil {
			(*r) = append((*r), nil)
			return
		}
		(*r) = append((*r), &n.v)
		inOrder(n.left, r)
		inOrder(n.right, r)
	}

	var t2Slice []*int
	inOrder(t2, &t2Slice)
	var tSlice []*int
	inOrder(t, &tSlice)

	printSlice(tSlice)
	printSlice(t2Slice)

	return isSubSlice(tSlice, t2Slice)
}

func printSlice(t []*int) {
	for _, v := range t {
		if v != nil {
			fmt.Printf("%v ", *v)
		} else {
			fmt.Printf("nil ")
		}

	}
	fmt.Printf("\n")
}

func main() {
	t1 := &Tree10{v: 1}
	t11, t12 := &Tree10{v: 4}, &Tree10{v: 10}
	t1.left, t1.right = t11, t12
	t111, t112 := &Tree10{v: 2}, &Tree10{v: 5}
	t11.left, t11.right = t111, t112

	t2 := &Tree10{v: 4}
	t21, t22 := &Tree10{v: 2}, &Tree10{v: 5}
	t2.left, t2.right = t21, t22

	fmt.Printf("t2 is sub-tree of t1 == %v\n", t1.isSubtree(t2))
}
