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

		root := nums2tree(nums)
		fmt.Println(levelOrder(root))
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
		nodes[i] = &TreeNode{Val: n}
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

func levelOrder(root *TreeNode) [][]int {

	ret := make([][]int, 0)

	var traversal func(cur *TreeNode, level int)
	traversal = func(cur *TreeNode, level int) {
		if cur == nil {
			return
		}

		if level >= len(ret) {
			ret = append(ret, []int{})
		}

		traversal(cur.Left, level+1)
		traversal(cur.Right, level+1)
		ret[level] = append(ret[level], cur.Val)
	}
	traversal(root, 0)
	return ret
}
