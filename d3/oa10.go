package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {

	maxSum := root.Val

	var traversal func(root *TreeNode) int
	traversal = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := max(traversal(root.Left), 0)
		right := max(traversal(root.Right), 0)

		tmpMax := left + right + root.Val
		maxSum = max(maxSum, tmpMax)
		return max(left, right) + root.Val

	}

	traversal(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
