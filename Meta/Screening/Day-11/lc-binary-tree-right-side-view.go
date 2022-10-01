/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	/*
	   I think we will need some form of traversal that exposes right side first, where we populate a slice with values for matching level if not filled yet

	   The traversal should be: right, self, left and it will submit value bottom up
	*/
	if root == nil {
		return []int{}
	}
	acc := make(map[int]*TreeNode)
	var maxDepth int
	var reverseInOrder func(*TreeNode, int)
	reverseInOrder = func(t *TreeNode, depth int) {
		if t == nil {
			return
		}
		//right
		reverseInOrder(t.Right, depth+1)
		//center
		if _, ok := acc[depth]; !ok {
			acc[depth] = t
		}
		if depth > maxDepth {
			maxDepth = depth
		}
		//left
		reverseInOrder(t.Left, depth+1)
	}
	reverseInOrder(root, 0)
	ans := make([]int, maxDepth+1)
	for i := range ans {
		ans[i] = acc[i].Val
	}
	return ans
}
