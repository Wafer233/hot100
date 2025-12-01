package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	for {
		var line string
		_, _ = fmt.Scanln(&line)
		parts := strings.Split(line, ",")
		nums := make([]int, len(parts))
		for i, p := range parts {
			nums[i], _ = strconv.Atoi(p)
		}
		fmt.Println(leveOrder(sortedArrayToBST(nums)))
	}
}

func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	rootIndex := len(nums) / 2
	rootVal := nums[rootIndex]

	left := nums[:rootIndex]
	right := nums[rootIndex+1:]

	return &TreeNode{
		Val:   rootVal,
		Left:  sortedArrayToBST(left),
		Right: sortedArrayToBST(right),
	}
}

func leveOrder(root *TreeNode) [][]int {

	ret := make([][]int, 0)
	var traverse func(root *TreeNode, level int)
	traverse = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(ret) == level {
			ret = append(ret, []int{})
		}
		traverse(root.Left, level+1)
		traverse(root.Right, level+1)
		ret[level] = append(ret[level], root.Val)
	}

	traverse(root, 0)
	return ret
}
