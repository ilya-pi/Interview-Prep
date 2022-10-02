/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	/*

	   We essentially need to walk through all nodes of the tree and pick the biggest left path + right path node

	   This should be possible in recursion, I think, let's see:

	   We are in node n, and we get longestPathLeft and we get longestPathRight
	   Then we will need to return the longer one of them +1
	   If at this node we happen to see lPL + lPR + 2 > current best, then update it

	   Exit cases: if nil -> then 0

	     1
	    / \
	   2   3
	        \
	         4

	*/

	if root == nil {
		return 0
	}

	var ans int
	var walk func(*TreeNode) int
	walk = func(t *TreeNode) int {
		// Exit case
		if t.Left == nil && t.Right == nil {
			return 0
		}
		var longestPathLeft int
		if t.Left != nil {
			longestPathLeft = walk(t.Left) + 1
		}
		var longestPathRight int
		if t.Right != nil {
			longestPathRight = walk(t.Right) + 1
		}
		// Try as best diameter
		diameter := longestPathLeft + longestPathRight
		if diameter > ans {
			ans = diameter
		}
		// Return longest path
		longestPath := longestPathLeft
		if longestPathRight > longestPathLeft {
			longestPath = longestPathRight
		}
		return longestPath
	}
	walk(root)
	return ans
}
