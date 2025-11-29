package main

//func main() {
//
//	var n int
//	for {
//
//		num, _ := fmt.Scan(&n)
//		if num == 0 {
//			break
//		} else {
//			fmt.Println(generate(n))
//		}
//	}
//}

func generate(n int) []string {

	cur := make([]byte, 0)
	ret := make([]string, 0)

	var backtracking func(left, right int)
	backtracking = func(left, right int) {
		//结束条件
		if len(cur) == 2*n {
			tmp := make([]byte, 0)
			tmp = append(tmp, cur...)
			ret = append(ret, string(tmp))
			return
		}

		if left < n {
			cur = append(cur, '(')
			backtracking(left+1, right)
			cur = cur[:len(cur)-1]
		}

		if left > right {
			cur = append(cur, ')')
			backtracking(left, right+1)
			cur = cur[:len(cur)-1]
		}

	}

	backtracking(0, 0)
	return ret

}
