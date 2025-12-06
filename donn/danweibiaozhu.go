package main

import "fmt"

func main() {
	nums := []int{4, 4, 1, 1, 1, 4, 4, 4}
	fmt.Println(danweibiaozhu(nums))
}

func danweibiaozhu(nums []int) int {

	dp := make([][]int, len(nums))
	for index, _ := range nums {
		dp[index] = make([]int, len(nums))
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				dp[i][j] = 1
			} else {
				dp[i][j] = 400
			}
		}
	}

	for length := 2; length <= len(nums); length++ {

		for l := 0; l+length-1 < len(nums); l++ {
			r := l + length - 1

			for k := l; k < r; k++ {

				dp[l][r] = min(dp[l][k]+dp[k+1][r], dp[l][r])
			}

			if nums[l] == nums[r] {
				dp[l][r] = dp[l][r-1]
			}
		}
	}

	return dp[0][len(nums)-1]
}
