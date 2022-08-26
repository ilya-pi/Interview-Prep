package main

import (
	"fmt"
	"math"
)

type Tree4 struct {
	v           int
	left, right *Tree4
}

func (t *Tree4) isBalanced() bool {
	/*
	   We can run DFS and record the first found depth,
	   while comparing it to other found depths,
	   Once we captured all depths, we can check that
	   we have no more then 2 values there (otherwise they diff more then 1
	*/

	depths := map[int32]bool{}

	var dfs func(*Tree4, int32)
	dfs = func(n *Tree4, d int32) {
		if n.left != nil {
			dfs(n.left, d+1)
		}
		if n.right != nil {
			dfs(n.right, d+1)
		}
		if n.left == nil && n.right == nil {
			// Leaf node, record depth
			depths[d] = true
		}
	}
	dfs(t, 0)

	fmt.Printf("Found depths: %v\n", depths)
	var ar []int32
	for k, _ := range depths {
		ar = append(ar, k)
	}
	return len(ar) == 1 || (len(ar) == 2 && math.Abs(float64(ar[0]-ar[1])) < 2)
}

func main() {
	t := &Tree4{v: 1}
	t1 := &Tree4{v: 2}
	t2 := &Tree4{v: 3}
	t3 := &Tree4{v: 4}
	t4 := &Tree4{v: 5}
	t5 := &Tree4{v: 6}
	t6 := &Tree4{v: 7}
	t7 := &Tree4{v: 8}
	//t8 := &Tree4{v: 9}
	t.left = t1
	t.right = t2
	t1.left = t3
	t1.right = t4
	t2.left = t5
	t2.right = t6
	t4.left = t7
	//t7.right = t8
	fmt.Printf("Is Balanced == %v\n", t.isBalanced())
}
