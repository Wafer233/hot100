package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LCA(root *TreeNode, p, q *TreeNode) *TreeNode {

	var traversal func(root *TreeNode) *TreeNode
	traversal = func(root *TreeNode) *TreeNode {
		switch root {
		case nil:
			return nil
		case p:
			return p
		case q:
			return q
		}

		left := traversal(root.Left)
		right := traversal(root.Right)

		switch {
		case left == nil:
			return right
		case right == nil:
			return left
		default:
			return root
		}
	}

	return traversal(root)

}
