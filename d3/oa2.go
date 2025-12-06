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
	var line string
	_, _ = fmt.Scanln(&line)
	parts := strings.Split(line, ",")
	nums := make([]int, len(parts))
	for i, p := range parts {
		if p == "null" {
			nums[i] = 101

		} else {
			nums[i], _ = strconv.Atoi(p)
		}

	}

	root := nums2tree(nums)
	flatten(root)

	show(root)
}

func nums2tree(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(nums))

	for i, n := range nums {
		if n == 101 {
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

func show(root *TreeNode) [][]int {
	levelorder := make([][]int, 0)
	var traverse func(root *TreeNode, level int)
	traverse = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(levelorder) == level {
			levelorder = append(levelorder, []int{})
		}
		traverse(root.Left, level+1)
		traverse(root.Right, level+1)
		levelorder[level] = append(levelorder[level], root.Val)
	}
	traverse(root, 0)
	return levelorder

}

func flatten(root *TreeNode) {

	preorder := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		preorder = append(preorder, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}

	traversal(root)

	var build func(nums []int) *TreeNode
	build = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}

		rootVal := nums[0]
		right := nums[1:]

		return &TreeNode{
			Val:   rootVal,
			Right: build(right),
		}
	}

	newTree := build(preorder)

	if root == nil {
		return
	}
	root.Val = newTree.Val
	root.Left = newTree.Left
	root.Right = newTree.Right
}
