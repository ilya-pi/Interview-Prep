import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	/*

	   It seems like I can calculate the "best" way to come into the node I am looking at and compare it with max. It is, I guess, somewhat post-order walk trough?

	   Then at each "visit" I should see if I want to record the max
	*/

	ans := math.MinInt
	var postOrder func(*TreeNode) int
	postOrder = func(t *TreeNode) int {
		if t == nil {
			return 0
		}
		if t.Left == nil && t.Right == nil {
			if ans < t.Val {
				ans = t.Val
			}
			return t.Val
		}
		maxLeft := postOrder(t.Left)
		if maxLeft < 0 {
			maxLeft = 0
		}
		maxRight := postOrder(t.Right)
		if maxRight < 0 {
			maxRight = 0
		}
		// try Max value
		if ans < maxLeft+maxRight+t.Val {
			ans = maxLeft + maxRight + t.Val
		}
		// return best option
		max := maxLeft
		if maxRight > max {
			max = maxRight
		}
		return max + t.Val
	}
	postOrder(root)
	return ans

}
