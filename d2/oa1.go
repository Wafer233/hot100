package d2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	for {
		var line string
		_, _ = fmt.Scanln(&line)
		part := strings.Split(line, ",")
		nums1 := make([]int, len(part))
		for index, value := range part {
			nums1[index], _ = strconv.Atoi(value)
		}
		var line2 string
		_, _ = fmt.Scanln(&line2)
		part2 := strings.Split(line2, ",")
		nums2 := make([]int, len(part2))
		for index, value := range part2 {
			nums2[index], _ = strconv.Atoi(value)
		}

		fmt.Println(findMediumSortedArray(nums1, nums2))
	}
}

func findMediumSortedArray(nums1 []int, nums2 []int) float64 {
	// 找到最小的当成nums1
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m := len(nums1)
	n := len(nums2)
	left := 0
	right := m

	for left <= right {

		//找分割点
		partitionA := (left + right) / 2
		partitionB := (m+n+1)/2 - partitionA

		maxLeftA := math.MinInt
		if partitionA != 0 {
			maxLeftA = nums1[partitionA-1]
		}
		minRightA := math.MaxInt
		if partitionA != m {
			minRightA = nums1[partitionA]
		}

		maxLeftB := math.MinInt
		if partitionB != 0 {
			maxLeftB = nums2[partitionB-1]
		}
		minRightB := math.MaxInt
		if partitionB != n {
			minRightB = nums2[partitionB]
		}

		if maxLeftA <= minRightB && maxLeftB <= minRightA {

			if (m+n)%2 == 0 {
				maxLeft := math.Max(float64(maxLeftA), float64(maxLeftB))
				minRight := math.Min(float64(minRightA), float64(minRightB))
				return (maxLeft + minRight) / 2.0
			} else {
				return math.Max(float64(maxLeftA), float64(maxLeftB))
			}

		} else if maxLeftA > minRightB {
			right = partitionA - 1
		} else {
			left = partitionA + 1
		}
	}
	return 0.0
}

//
//nums1[i-1] < nums2[j]
//nums2[j-1] < nums1[i]
// i+j = (len1+len2+1)/2
