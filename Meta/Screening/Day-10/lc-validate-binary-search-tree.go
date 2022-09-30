import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	/*
	   Valid BST will have all nodes on left smaller then current value and all on right bigger then current value. And that will contonuously narrow down the interval of valid values

	   So we'll have a helper function "valid" that takes interval and adjusts it on the current node's value when recursing
	*/

	var valid func(*TreeNode, int, int) bool
	valid = func(t *TreeNode, left, right int) bool {
		if t == nil {
			return true
		}
		if t.Val > left && t.Val < right {
			return valid(t.Left, left, t.Val) && valid(t.Right, t.Val, right)
		}
		return false
	}
	return valid(root, math.MinInt, math.MaxInt)
}
