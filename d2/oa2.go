package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	for {
		var line string
		_, _ = fmt.Scanln(&line)
		part := strings.Split(line, ",")
		nums := make([]int, len(part))
		for index, value := range part {
			nums[index], _ = strconv.Atoi(value)
		}
		var target int
		_, _ = fmt.Scan(&target)
		fmt.Println(SearchInsert(nums, target))
	}
}

func SearchInsert(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
