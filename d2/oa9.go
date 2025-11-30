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
			if p == "null" {
				nums[i] = -1
			}
			nums[i], _ = strconv.Atoi(p)
		}

		root := nums2Tree(nums)
		fmt.Println(maxDepth(root))
	}
}

func nums2Tree(nums []int) *TreeNode {
	if len(nums) == 0 || nums[0] == -1 {
		return nil
	}

	nodes := make([]*TreeNode, len(nums))
	for i, v := range nums {
		if v == -1 {
			continue
		}
		nodes[i] = &TreeNode{Val: v, Left: nil, Right: nil}
	}

	for i := 0; i < len(nodes); i++ {
		left := 2*i + 1
		right := 2*i + 2
		if left < len(nums) {
			nodes[i].Left = nodes[left]
		}
		if right < len(nums) {
			nodes[i].Right = nodes[right]
		}
	}
	return nodes[0]
}

func maxDepth(root *TreeNode) int {

	var traversal func(cur *TreeNode) int
	traversal = func(cur *TreeNode) int {

		if cur == nil {
			return 0
		}
		left := traversal(cur.Left)
		right := traversal(cur.Right)
		return max(left, right) + 1
	}
	return traversal(root)

}
