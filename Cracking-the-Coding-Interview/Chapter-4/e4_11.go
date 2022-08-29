package main

import (
	"fmt"
	"math/rand"
)

type Tree11 struct {
	v           int
	parent      *Tree11
	left, right *Link
}

type Link struct {
	count int
	n     *Tree11
}

func (t *Tree11) insert(v int) {
	n := &Tree1{v: v}
	if t == nil {
		(*t) = n
	}
	k := t
	for (v < k.v && k.left != nil) || (v >= k.v && k.right != nil) {
		if v < k.v {
			k.left.count++
			k = k.left.n
		} else {
			k.right.count++
			k = k.right.n
		}
	}
	n.parent = k
	// Technically the null check here is obsolete
	if v < k.v && k.left == nil {
		k.left = &Link{count: 1, n: n}
	} else if v >= k.v && k.right == nil {
		k.right = &Link{count: 1, n: n}
	}
}

func (t *Tree11) find(v int) {
	k := t
	for k.v != v {
		if v < k.v {
			if k.left == nil {
				return nil
			}
			k = k.left.n
		} else {
			if k.right == nil {
				return nil
			}
			k = r.right.n
		}
	}
	return k
}

func (t *Tree11) nodes() []*Tree11 {
	var res []*Tree16

	var walk func(n *Tree11)
	walk = func(n *Tree11) {
		if n == nil {
			return
		}
		res = append(res, n)
		if n.left != nil {
			walk(n.left.n)
		}
		if n.right != nil {
			walk(n.right.n)
		}
	}

	return res
}

func (t *Tree11) leftChild() bool {
	if t != nil && t.parent.left == t {
		return true
	} else {
		return false
	}
}

func (t *Tree11) delete(v int) {
	/*
		1. if it is a leaf node — just delete it
		2. if this node has one child — just move it up
		3. otherwise - find an in-order successor and
		change value with that node, then delete it
	*/
	n := t.find(v)
	// not found
	if n == nil {
		return
	}

	// leaf node
	if n.isLeaf() {
		if n.parent == nil {
			(*t) = nil
		}

		if n.parent.left == n {
			n.parent.left = nil
		}
		if n.parent.right == n {
			n.parent.right = nil
		}
		// Correct the left and right amount in the node
		for k := n; k != nil; {
			if k.leftChild() {
				k.parent.left.count--
				k = k.parent
			} else {
				k.parent.right.count--
				k = k.parent
			}
		}
		return
	}

	// single node
	if n.left == nil {
		n.v = n.right.v
		n.left, n.right = n.right.left, n.right.right
		return
	}
	if n.right == nil {
		n.v = n.left.v
		n.left, n.right = n.left.left, n.left.right
		return
	}

	// find in-order successor and replace with that node
	// It is only the case when n.right != nil => hence it
	// is the left most node or right sub-tree
	succ := n.right
	for ; succ.left != nil; succ = succ.left {
	}
	n.v = succ.v
	if succ.isLeaf() {
		succ.parent.left = nil
	} else {
		succ.parent.left, succ.right.parent = succ.right, succ.parent
	}
	// Correct amounts in the tree
	for k := succ.parent; k != nil; {
		if k.leftChild() {
			k.parent.left.count--
			k = k.parent
		} else {
			k.parent.right.count--
			k = k.parent
		}
	}
}

func (t *Tree11) isLeaf() bool {
	return t != nil && t.left == nil && t.right == nil
}

func (t *Tree11) getRandomNode() *Tree11 {
	/*

		Say we have a well-distributed rand function

		Approach that is easy to understand:
		Store elements in a linked list and select random one on function

		Approach 2:
		Evenly select at each level whether to go down.
		If there are 50 nodes in this tree (20 left + 1 parent + 29 right)
		the numbers 1-50 would tell us whether we go left, right or pick this node

	*/
	for k := t; k.left != nil || k.right != nil; {
		tot := 1 + k.left.count + k.right.count
		ran := 1 + rand.Intn(tot)
		switch {
		case ran == k.left.count+1:
			return k
		case ran < k.left.count+1:
			k = k.left.n
		case ran > k.left.count+1:
			k = k.right.n
		}
	}
	// Genuinely speaking unreachable
	return nil
}

func (t *Tree11) String() string {
	return fmt.Sprintf("%v", t.v)
}

func main() {
	t := &Tree11{v: 1}
	for i := 0; i < 10; i++ {
		t.insert(i)
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("Random node is %v\n", t.getRandomNode())
	}
}
