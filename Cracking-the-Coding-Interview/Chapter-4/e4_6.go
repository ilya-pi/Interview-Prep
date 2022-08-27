package main

import "fmt"

type Tree6 struct {
	v                   int
	parent, left, right *Tree6
}

func (t *Tree6) successor() *Tree6 {
	/*

	   There are three cases:
	   1) we have right node from current? -> most left in the right subtree
	   2) we came to this node from left -> go up until we can come from right
	   3) came from right -> that is the node

	   2)+3) are the same

	*/
	// 1 Right sub-tree is not nil
	if t.right != nil {
		k := t.right
		for ; k.left != nil; k = k.left {
		}
		return k
	}

	// 2 Came from left, no right subtree
	k := t
	for ; k.parent != nil && k.parent.left != k; k = k.parent {
	}
	if k.parent != nil {
		return k.parent
	}
	// This is the right-most node in the tree
	return nil
}

func (t *Tree6) String() string {
	if t == nil {
		return "nil"
	}
	return fmt.Sprintf("(%v)", t.v)
}

func main() {
	t := &Tree6{v: 10, parent: nil}
	t1 := &Tree6{v: 5, parent: nil}
	t2 := &Tree6{v: 15, parent: nil}
	t3 := &Tree6{v: 2, parent: nil}
	t4 := &Tree6{v: 4, parent: nil}
	t5 := &Tree6{v: 12, parent: nil}
	t.left, t.right, t.parent = t1, t2, nil
	t1.left, t1.right, t1.parent = t3, t4, t
	t2.left, t2.right, t2.parent = t5, nil, t
	t3.parent = t1
	t4.parent = t1
	t5.parent = t2

	fmt.Printf("Successor of %v is %v\n", t5, t5.successor())
}
