/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	/*

	   pre-order is recursive I believe

	   Given node, you recur left, visit self, recur right

	   To do it in place, you could return list of node from self, then basically gluing it in place with


	   res := recur left
	   most right of res := self
	   self.right = recur right

	   if nil -> return nil

	      1
	     2 5
	*/

	var preOrder func(*TreeNode) *TreeNode
	preOrder = func(t *TreeNode) *TreeNode {
		if t == nil {
			return nil
		}
		if t.Left == nil && t.Right == nil {
			return t
		}
		// cut this node
		left := t.Left
		right := t.Right
		t.Left, t.Right = nil, nil
		t.Right = preOrder(left)
		mostRight := t
		for ; mostRight.Right != nil; mostRight = mostRight.Right {
		}
		mostRight.Right = preOrder(right)
		return t
	}
	preOrder(root)
}
