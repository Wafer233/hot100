package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBT(root *TreeNode) int {

	diameter := 0

	var traversal func(root *TreeNode) int
	traversal = func(root *TreeNode) int {

		if root == nil {
			return 0
		}
		left := traversal(root.Left)
		right := traversal(root.Right)

		diameter = max(diameter, left+right)
		return max(left, right) + 1
	}
	traversal(root)
	return diameter
}
