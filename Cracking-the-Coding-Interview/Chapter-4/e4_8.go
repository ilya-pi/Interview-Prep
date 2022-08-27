package main

import "fmt"

type Tree8 struct {
	v                   int
	parent, left, right *Tree8
}

func (t *Tree8) String() string {
	return fmt.Sprintf("%v", t.v)
}

func firstCommonAncestor(n1, n2 *Tree8) *Tree8 {
	/*
		We'll calculate height to the top of the tree
		and then start going up with offset on one pointer
	*/

	h1, h2 := 0, 0
	for k := n1; k.parent != nil; k = k.parent {
		h1++
	}
	for k := n2; k.parent != nil; k = k.parent {
		h2++
	}

	if h1 < h2 {
		h1, n1, h2, n2 = h2, n2, h1, n1
	}

	k1 := n1
	for i := 0; i < h1-h2; i++ {
		k1 = k1.parent
	}

	for k2 := n2; k1 != k2; k1, k2 = k1.parent, k2.parent {
	}

	return k1
}

func main() {
	t := &Tree8{v: 1}
	t1 := &Tree8{v: 2}
	t2 := &Tree8{v: 3}
	t3 := &Tree8{v: 4}
	t4 := &Tree8{v: 5}
	t5 := &Tree8{v: 6}
	t6 := &Tree8{v: 7}
	t.parent, t.left, t.right = nil, t1, t2
	t1.parent, t1.left, t1.right = t, t3, t4
	t3.parent = t1
	t4.parent = t1
	t2.parent, t2.left, t2.right = t, t5, t6
	t5.parent, t6.parent = t2, t2

	fmt.Printf("First common ancestor %v and %v is %v\n", t4, t6, firstCommonAncestor(t4, t6))
}
