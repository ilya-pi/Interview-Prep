package main

import "fmt"

type BST struct {
	v                   int
	parent, left, right *BST
}

func (b *BST) insert(v int) {
	if b == nil {
		panic("alloc tree first")
	}

	if v < b.v {
		if b.left == nil {
			b.left = &BST{v: v, parent: b}
		} else {
			b.left.insert(v)
		}
	}
	if v > b.v {
		if b.right == nil {
			b.right = &BST{v: v, parent: b}
		} else {
			b.right.insert(v)
		}
	}
	// if == , then we already have it
}

func (b *BST) String() string {
	res := ""

	var pr func(b *BST)
	pr = func(b *BST) {
		if b == nil {
			return
		}

		pr(b.left)
		res += fmt.Sprintf("%d ", b.v)
		pr(b.right)
	}
	pr(b)
	return res
}

func (b *BST) find(v int) *BST {
	if b == nil {
		return nil
	}

	if b.v == v {
		return b
	}

	if b.v > v {
		return b.left.find(v)
	}

	if b.v < v {
		return b.right.find(v)
	}

	return nil
}

func (b *BST) next(v int) *BST {
	n := b.find(v)
	if n == nil {
		return nil
	}

	/*
		The bigger value must be on the right side of current node
		If current node has right node, then it will be most left in the right subtree
		If there is no right child, it will be first parent on the "right" (where we come through left child)

	*/
	// If current node has right node, then it will be most left in the right subtree
	if n.right != nil {
		n = n.right
		for ; n.left != nil; n = n.left {
		}
		return n
	}
	// If there is no right child, it will be first parent on the "right" (where we come through left child)
	for ; n.parent != nil && n.parent.left != n; n = n.parent {
	}
	return n.parent
}

func main() {
	t := BST{v: 10}
	t.insert(5)
	t.insert(9)
	t.insert(2)
	t.insert(19)
	t.insert(29)
	t.insert(1)
	fmt.Printf("BST - %v\n", &t)
	fmt.Printf("Next to 5 is %d\n", t.next(5).v)
	fmt.Printf("Next to 9 is %d\n", t.next(9).v)
	fmt.Printf("Next to 10 is %d\n", t.next(10).v)
}
