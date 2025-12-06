package d1

//func main() {
//	var s string
//	_, _ = fmt.Scan(&s)
//	fmt.Println(Palindrome(s))
//}

func Palindrome(s string) [][]string {

	cur := make([]string, 0)
	ret := make([][]string, 0)

	var backtracking func(index int)
	backtracking = func(index int) {
		if index == len(s) {
			tmp := make([]string, len(cur))
			copy(tmp, cur)
			ret = append(ret, tmp)
			return
		}

		for i := index; i < len(s); i++ {
			curSub := s[index : i+1]
			if isPali(curSub) {
				cur = append(cur, curSub)
				backtracking(i + 1)
				cur = cur[:len(cur)-1]
			}
		}
	}
	backtracking(0)
	return ret
}

func isPali(s string) bool {

	l := 0
	r := len(s) - 1

	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
