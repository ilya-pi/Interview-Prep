/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalOrder(root *TreeNode) [][]int {
	/*
	    We can count column number from root as zero, every step left decreases column number
	       1
	      / \
	     2.  3
	    /\.   \
	   5   7
	        \
	         19

	         0: 1, 7
	         -1: 2
	         -2: 5
	         ...
	    And further. If we keep track of the lowest column, then we can map the results to a slice.

	    Order seem to be naturally top to bottom and left to right, if we do self-left-right

	    So the preOrder would be -
	    func(*Node, column)

	    and we will accumulate results to map[int][]int

	    acc[0] = {1, 7}
	    acc[-1] = {2}
	    acc[-2] = {5}
	    acc[0] = {}
	    ...
	*/

	if root == nil {
		return [][]int{}
	}

	acc := make(map[int][]int)
	var minColumn int

	type nodeAndColumn struct {
		node   *TreeNode
		column int
	}
	q := []nodeAndColumn{{root, 0}}
	for len(q) > 0 {
		node, column := q[0].node, q[0].column
		q = q[1:]

		if column < minColumn {
			minColumn = column
		}
		acc[column] = append(acc[column], node.Val)

		if node.Left != nil {
			q = append(q, nodeAndColumn{node.Left, column - 1})
		}
		if node.Right != nil {
			q = append(q, nodeAndColumn{node.Right, column + 1})
		}
	}
	/*
	   var preOrder func(*TreeNode, int)
	   preOrder = func(t *TreeNode, column int) {
	       if t == nil {
	           return
	       }
	       if column < minColumn {
	           minColumn = column
	       }
	       acc[column] = append(acc[column], t.Val)
	       preOrder(t.Left, column - 1)
	       preOrder(t.Right, column + 1)
	   }
	   preOrder(root, 0)
	*/

	var ans [][]int
	for i := minColumn; acc[i] != nil; i++ {
		ans = append(ans, acc[i])
	}
	return ans
}
