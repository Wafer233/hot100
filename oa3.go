package main

//func main() {
//
//	for {
//
//		var line string
//		_, _ = fmt.Scanln(&line)
//
//		part := strings.Split(line, ",")
//		candidate := make([]int, len(part))
//		for i, v := range part {
//			candidate[i], _ = strconv.Atoi(v)
//		}
//
//		var target int
//		_, _ = fmt.Scan(&target)
//
//		fmt.Println(CombinationSum(candidate, target))
//	}
//}

func CombinationSum(candidate []int, target int) [][]int {

	cur := make([]int, 0)
	ret := make([][]int, 0)
	curSum := 0

	var backtracking func(index int)
	backtracking = func(index int) {

		//结束条件
		if curSum >= target {
			if curSum == target {
				tmp := make([]int, 0)
				tmp = append(tmp, cur...)
				ret = append(ret, tmp)
			}
			return
		}

		for i := index; i < len(candidate); i++ {
			curSum += candidate[i]
			cur = append(cur, candidate[i])

			backtracking(i)

			curSum -= candidate[i]
			cur = cur[:len(cur)-1]
		}
	}

	backtracking(0)
	return ret
}
