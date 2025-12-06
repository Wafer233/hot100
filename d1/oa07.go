package d1

//func main() {
//
//	for {
//		var line string
//		_, _ = fmt.Scan(&line)
//
//		part := strings.Split(line, ",")
//		board := make([][]byte, len(part))
//		//abcd
//		for index, _ := range part {
//			bytes := []byte(part[index])
//			board[index] = bytes
//		}
//
//		var word string
//		_, _ = fmt.Scan(&word)
//		fmt.Println(WordSearch(board, word))
//
//	}
//}

func WordSearch(board [][]byte, word string) bool {

	used := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		used[i] = make([]bool, len(board[i]))
	}

	var backtracking func(row, col, index int) bool
	backtracking = func(row, col, index int) bool {

		// 推出条件
		if index == len(word) {
			return true
		}

		if row < 0 ||
			row >= len(board) ||
			col < 0 ||
			col >= len(board[0]) ||
			used[row][col] ||
			word[index] != board[row][col] {
			return false
		}

		used[row][col] = true

		if backtracking(row-1, col, index+1) ||
			backtracking(row+1, col, index+1) ||
			backtracking(row, col-1, index+1) ||
			backtracking(row, col+1, index+1) {
			return true
		}
		used[row][col] = false
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if word[0] == board[i][j] && backtracking(i, j, 0) {
				return true
			}
		}
	}
	return false
}
