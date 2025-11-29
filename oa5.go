package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	fmt.Println(NQueens(n))
}

func NQueens(n int) [][]string {

	cur := make([][]byte, 0)
	ret := make([][]string, 0)

	//初始化棋盘
	for i := 0; i < n; i++ {
		cur = append(cur, []byte{})
		for j := 0; j < n; j++ {
			cur[i] = append(cur[i], '.')
		}
	}

	var backtracking func(row int)
	backtracking = func(row int) {

		//结束条件
		if row == n {
			curStr := make([]string, 0)
			for i := 0; i < n; i++ {
				tmp := make([]byte, n)
				copy(tmp, cur[i])
				curStr = append(curStr, string(tmp))
			}
			ret = append(ret, curStr)
		}

		for col := 0; col < n; col++ {
			if isValid(row, col, n, cur) {
				cur[row][col] = 'Q'
				backtracking(row + 1)
				cur[row][col] = '.'
			}

		}
	}

	backtracking(0)
	return ret

}

func isValid(row, col, n int, cur [][]byte) bool {

	for i := 0; i < row; i++ {
		if cur[i][col] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if cur[i][j] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if cur[i][j] == 'Q' {
			return false
		}
	}
	return true
}
