package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
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
		root := Nums2Tree(nums)
		fmt.Println(isValidBST(root))
	}

}

func Nums2Tree(nums []int) *TreeNode {

	if len(nums) == 0 || nums[0] == -1 {
		return nil
	}

	nodes := make([]*TreeNode, len(nums))
	for i, v := range nums {

		if v == -1 {
			continue
		}
		nodes[i] = &TreeNode{
			Val: v,
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

func isValidBST(root *TreeNode) bool {

	inorder := make([]int, 0)
	var traversal func(cur *TreeNode)
	traversal = func(cur *TreeNode) {

		if cur == nil {
			return
		}
		traversal(cur.Left)
		inorder = append(inorder, cur.Val)
		traversal(cur.Right)
	}
	traversal(root)

	curMax := inorder[0]
	if len(inorder) == 1 {
		return true
	}

	for i := 1; i < len(inorder); i++ {
		if inorder[i] <= curMax {
			return false
		}
		curMax = inorder[i]
	}
	return true
}
