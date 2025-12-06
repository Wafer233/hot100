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
		parts := strings.Split(line, ",")
		nums := make([]int, len(parts))
		for index, value := range parts {
			nums[index], _ = strconv.Atoi(value)
		}
		fmt.Println(FindMin(nums))
	}
}

func FindMin(nums []int) int {

	left := 0
	right := len(nums) - 1
	ret := nums[0]

	for left <= right {

		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {

			if nums[mid] < ret {
				ret = nums[mid]
			}
			right = mid - 1
		}
	}

	return ret
}
