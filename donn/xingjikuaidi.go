package main

import "fmt"

func main() {
	ranliao := 20
	tongxingzheng := 1
	kuai := [][]int{
		{8, 5},
		{7, 4},
		{10, 6},
	}
	fmt.Println(kuaidi(ranliao, tongxingzheng, kuai))
}
func kuaidi(ranliao int, tongxingzheng int, kuaidi [][]int) (int, int) {

	dp := make([][]int, ranliao+1)

	for i := range dp {
		dp[i] = make([]int, tongxingzheng+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	dp[0][0] = 0

	for _, item := range kuaidi {

		a := item[0]
		b := item[1]

		ndp := make([][]int, ranliao+1)
		for i := range ndp {
			ndp[i] = make([]int, tongxingzheng+1)
			copy(ndp[i], dp[i])
		}

		for i := ranliao; i >= 0; i-- {
			for j := tongxingzheng; j >= 0; j-- {

				if dp[i][j] == -1 {
					continue
				}

				// 一般配送
				if i+a <= ranliao {
					ndp[i+a][j] = max(ndp[i+a][j], dp[i][j]+1)
				}

				// 虫洞配送
				if i+b <= ranliao && j+1 <= tongxingzheng {
					ndp[i+b][j+1] = max(ndp[i+b][j+1], dp[i][j]+1)
				}
			}
		}

		dp = ndp

	}

	baoguo := 0
	xiaohao := ranliao

	for i := 0; i <= ranliao; i++ {
		for j := 0; j <= tongxingzheng; j++ {

			if dp[i][j] > baoguo {
				baoguo = dp[i][j]
				xiaohao = i
			} else if dp[i][j] == baoguo && dp[i][j] != -1 {
				xiaohao = min(xiaohao, i)
			}
		}
	}
	return baoguo, xiaohao

}
