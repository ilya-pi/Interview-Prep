package main

import "fmt"

type BST struct {
	v           int
	left, right *BST
}

func (b *BST) add(v int) {
	t := b
	for {
		if v < t.v {
			if t.left == nil {
				t.left = &BST{v: v}
				return
			}
			t = t.left
			continue
		}
		if v >= t.v {
			if t.right == nil {
				t.right = &BST{v: v}
				return
			}
			t = t.right
			continue
		}
	}
}

func (b BST) search(v int) int {
	t := &b
	for {
		if v == t.v {
			return t.v
		}
		if v < t.v {
			if t.left == nil {
				return -1
			}
			t = t.left
			continue
		}
		if v >= t.v {
			if t.right == nil {
				return -1
			}
			t = t.right
			continue
		}
	}
	// not possible
	return -2
}

func main() {
	t := BST{v: 3}
	t.add(2)
	t.add(23)
	t.add(10)
	t.add(5)
	fmt.Printf("BST is %v\n", t)
	fmt.Printf("Found %v\n", t.search(5))
	fmt.Printf("Not Found %v\n", t.search(7))
}
