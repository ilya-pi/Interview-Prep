package main

import "fmt"

/*
Approach:

Limitations:
• Is integer division oke?
• How big are the values?
• How big is the tree?

1. Have Tree holding structure
2. With BFS get values on every level and average them once we hit a new level
  • So we keep track of last see level of the node
  • Once we see new one -> average current values and move to new level
  • Can start writing to the target array right away (but need to use a slice as we don't know the depth of the tree yet)
3. Output result
*/

type Tree struct {
	v           int
	left, right *Tree
}

func (t *Tree) levelsAverage() []int {
	// In the queue we need node and depth of the node
	type NodeAndDepth struct {
		n     *Tree
		depth int
	}

	// Queue for BFS
	q := []NodeAndDepth{{t, 0}}
	var ans []int
	var level int
	var sum int
	var count int
	for len(q) > 0 {
		// Pop
		e := q[0]
		q = q[1:]

		if e.depth == level {
			// Still exploring current level
			sum += e.n.v
			count++
		} else {
			// Need to average out and reset sum and count
			avg := sum / count
			// Got to a new level
			sum = e.n.v
			count = 1
			level = e.depth
			// Record
			ans = append(ans, avg)
		}
		if e.n.left != nil {
			q = append(q, NodeAndDepth{e.n.left, e.depth + 1})
		}
		if e.n.right != nil {
			q = append(q, NodeAndDepth{e.n.right, e.depth + 1})
		}
	}
	// Process last level
	ans = append(ans, sum/count)
	return ans
}

func main() {
	/*
		Given a binary tree, get the average value at each level of the tree
		```
		Input:

		    4
		   / \
		  7   9
		/ \    \
		10  2    6
		    /
		   6
		  /
		 2

		[4, 8 , 6, 6, 2]
		```
	*/
	t := &Tree{v: 4}
	t1 := &Tree{v: 7}
	t2 := &Tree{v: 9}
	t3 := &Tree{v: 6}
	t.left, t.right = t1, t2
	t2.right = t3
	t11 := &Tree{v: 10}
	t12 := &Tree{v: 2}
	t21 := &Tree{v: 6}
	t31 := &Tree{v: 2}
	t1.left, t1.right = t11, t12
	t12.left = t21
	t21.left = t31
	fmt.Printf("Average per level is %v\n", t.levelsAverage())
}
