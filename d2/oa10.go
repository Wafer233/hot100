package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var line1 string
	var line2 string

	_, _ = fmt.Scanln(&line1)
	parts1 := strings.Split(line1, ",")
	_, _ = fmt.Scanln(&line2)
	parts2 := strings.Split(line2, ",")
	nums1 := make([]int, len(parts1))
	nums2 := make([]int, len(parts2))
	for i := 0; i < len(parts1); i++ {
		nums1[i], _ = strconv.Atoi(parts1[i])
	}
	for i := 0; i < len(parts2); i++ {
		nums2[i], _ = strconv.Atoi(parts2[i])
	}
	fmt.Println(buildTree(nums1, nums2))
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	targetIndex := slices.Index(inorder, rootVal)

	inorderLeft := inorder[:targetIndex]
	inorderRight := inorder[targetIndex+1:]
	preorderLeft := preorder[1 : 1+len(inorderLeft)]
	preorderRight := preorder[1+len(inorderLeft):]

	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(preorderLeft, inorderLeft),
		Right: buildTree(preorderRight, inorderRight),
	}
}
