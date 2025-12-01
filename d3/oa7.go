package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {

	prefixSum := make(map[int]int)

	prefixSum[0] = 1

	curSum := 0
	count := 0

	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}

		curSum += root.Val
		//curSum - prefixSum = targetSum
		if cnt, exist := prefixSum[curSum-targetSum]; exist {
			count += cnt
		}
		prefixSum[curSum]++
		traversal(root.Left)
		traversal(root.Right)
		prefixSum[curSum]--
		curSum -= root.Val
	}

	traversal(root)
	return count
}
