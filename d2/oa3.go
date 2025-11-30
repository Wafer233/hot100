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
		rows := strings.Split(line, ";")
		matrix := make([][]int, len(rows))
		for index, _ := range rows {
			cols := strings.Split(rows[index], ",")
			matrix[index] = make([]int, len(cols))
			for i, col := range cols {
				matrix[index][i], _ = strconv.Atoi(col)
			}
		}
		var target int
		_, _ = fmt.Scan(&target)
		fmt.Println(Search2DMatrix(matrix, target))
	}

}

func Search2DMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	left := 0
	right := m*n - 1

	for left <= right {
		mid := (left + right) / 2
		row := mid / n
		col := mid % n

		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
