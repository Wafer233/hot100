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
			} else {
				nums[i], _ = strconv.Atoi(p)
			}

		}
		root := BuildTree(nums)
		fmt.Println(inorderTraversal(root))
	}
}

func BuildTree(nums []int) *TreeNode {

	if len(nums) == 0 || nums[0] == -1 {
		return nil
	}

	nodes := make([]*TreeNode, len(nums))

	for i, v := range nums {
		if v == -1 {
			nodes[i] = nil
		} else {
			nodes[i] = &TreeNode{
				Val: v,
			}
		}

	}

	for i := 0; i < len(nodes); i++ {

		if nodes[i] == nil {
			continue
		}

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

func inorderTraversal(root *TreeNode) []int {

	ret := make([]int, 0)
	var traversal func(cur *TreeNode)
	traversal = func(cur *TreeNode) {
		if cur == nil {
			return
		}

		traversal(cur.Left)
		ret = append(ret, cur.Val)
		traversal(cur.Right)
	}
	traversal(root)
	return ret
}
