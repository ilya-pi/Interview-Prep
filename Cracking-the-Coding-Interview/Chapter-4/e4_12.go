package main

import "fmt"

type Tree12 struct {
	v           int
	left, right *Tree12
}

func (t *Tree12) numPaths(v int) int {
	/*
	   Expand out of the top of the tree, maintaining
	   a slice with all sums so far and it a sum == v,
	   we have one more solution. Each layer creates
	   new sum and adjusts current ones
	*/

	var countSums func(*Tree12, int, []int) int
	countSums = func(n *Tree12, target int, sums []int) int {
		if n == nil {
			return 0
		}
		fmt.Printf("n == %v Sums %v \n", n.v, sums)
		// 1. Update sums and see if we have a match
		for i, _ := range sums {
			sums[i] += n.v
		}
		// 2. Add current nodes val to the sums
		sums = append(sums, n.v)
		var r int
		for _, v := range sums {
			if v == target {
				r++
			}
		}
		// as slices are passed by link
		sums2 := make([]int, len(sums))
		copy(sums2, sums)
		// 3. Return recursive result with sub-children
		return r + countSums(n.left, target, sums) + countSums(n.right, target, sums2)
	}

	return countSums(t, v, nil)
}

func main() {
	t := &Tree12{v: 1}
	t1 := &Tree12{v: 2}
	t2 := &Tree12{v: 3}
	t3 := &Tree12{v: 4}
	t4 := &Tree12{v: 5}
	t5 := &Tree12{v: 3}
	t6 := &Tree12{v: 3}
	t.left, t.right = t1, t2
	t1.left, t1.right = t3, t4
	t2.left, t2.right = t5, t6
	fmt.Printf("Got %v ways to make %v", t.numPaths(7), 7)
}
