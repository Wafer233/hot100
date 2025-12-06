package main

import "fmt"

func main() {

	for {
		var s string

		_, _ = fmt.Scan(&s)
		print(longestPalindrome(s))
	}
}

func longestPalindrome(s string) string {

	//dp[i][j] 以i开头，j结尾的index的string
	//dp[i][j] = true
	//dp[i+1][j-1] = true && s[i] = s[j] (i+1 <= j-1)
	//

	var ret string

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for i := range dp {
		dp[i][i] = true
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {

			if s[i] == s[j] {
				if i+1 > j-1 || dp[i+1][j-1] {
					dp[i][j] = true

					curLen := j - i + 1
					if curLen > len(ret) {
						ret = s[i : j+1]
					}
				}
			}
		}
	}

	return ret
}
