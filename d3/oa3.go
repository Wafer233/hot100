package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {

	levelOrder := make([][]int, 0)

	var tarversal func(root *TreeNode, level int)
	tarversal = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(levelOrder) == level {
			levelOrder = append(levelOrder, []int{})
		}
		tarversal(root.Left, level+1)
		tarversal(root.Right, level+1)
		levelOrder[level] = append(levelOrder[level], root.Val)
	}
	tarversal(root, 0)

	ret := make([]int, len(levelOrder))

	for index, value := range levelOrder {
		ret[index] = value[len(value)-1]
	}
	return ret
}
