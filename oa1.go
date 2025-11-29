package main

//func main() {
//
//	digit := ""
//
//	for {
//		num, _ := fmt.Scan(&digit)
//		if num == 0 {
//			break
//		} else {
//			fmt.Println(LetterCombination(digit))
//		}
//	}
//}

func LetterCombination(digits string) []string {

	ret := make([]string, 0)

	cur := make([]byte, 0)

	num2char := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var backtracking func(index int)
	backtracking = func(index int) {

		// end condition
		if index == len(digits) {
			tmp := make([]byte, 0)
			tmp = append(tmp, cur...)
			ret = append(ret, string(tmp))
			return
		}

		curCharacters := num2char[digits[index]]

		for i := 0; i < len(curCharacters); i++ {
			cur = append(cur, curCharacters[i])
			backtracking(index + 1)
			cur = cur[:len(cur)-1]
		}
	}

	backtracking(0)
	return ret
}
