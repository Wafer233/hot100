package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmall(root *TreeNode, k int) int {

	inorder := make([]int, 0)

	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}

		traversal(root.Left)
		inorder = append(inorder, root.Val)
		traversal(root.Right)
	}
	traversal(root)
	return inorder[k-1]
}
