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
		root := nums2tree(nums)
		fmt.Println(isSymmetric(root))
	}
}

func nums2tree(nums []int) *TreeNode {

	if len(nums) == 0 || nums[0] == -1 {
		return nil
	}

	nodes := make([]*TreeNode, len(nums))

	for i, n := range nums {
		if n == -1 {
			continue
		}
		nodes[i] = &TreeNode{
			Val: n,
		}
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

func isSymmetric(root *TreeNode) bool {

	var compare func(a, b *TreeNode) bool
	compare = func(a, b *TreeNode) bool {

		switch {
		case a == nil && b == nil:
			return true
		case a == nil && b != nil:
			return false
		case a != nil && b == nil:
			return false
		case a != nil && b != nil && a.Val != b.Val:
			return false
		}

		inner := compare(a.Right, b.Left)
		outer := compare(a.Left, b.Right)
		return inner && outer
	}
	return compare(root.Left, root.Right)
}
