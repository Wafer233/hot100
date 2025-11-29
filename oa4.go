package main

//
//func main() {
//
//	for {
//		var line string
//		_, _ = fmt.Scanln(&line)
//
//		part := strings.Split(line, ",")
//		nums := make([]int, len(part))
//
//		for i, s := range part {
//			nums[i], _ = strconv.Atoi(s)
//		}
//
//		fmt.Println(Permutation(nums))
//	}
//}

func Permutation(nums []int) [][]int {

	cur := make([]int, 0)
	ret := make([][]int, 0)
	used := make(map[int]bool)

	var backtracking func()
	backtracking = func() {

		if len(cur) == len(nums) {
			tmp := make([]int, 0)
			tmp = append(tmp, cur...)
			ret = append(ret, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {

			if !used[nums[i]] {
				used[nums[i]] = true
				cur = append(cur, nums[i])
				backtracking()
				used[nums[i]] = false
				cur = cur[:len(cur)-1]
			}
		}

	}

	backtracking()
	return ret
}
