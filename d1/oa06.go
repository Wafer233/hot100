package d1

//func main() {
//	var line string
//	_, _ = fmt.Scanln(&line)
//
//	part := strings.Split(line, ",")
//	nums := make([]int, len(part))
//	for index, value := range part {
//		nums[index], _ = strconv.Atoi(value)
//	}
//
//	fmt.Println(SubSet(nums))
//}

func SubSet(nums []int) [][]int {

	cur := make([]int, 0)
	ret := make([][]int, 0)

	var backtracking func(index int)
	backtracking = func(index int) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		ret = append(ret, tmp)
		if len(cur) == len(nums) {
			return
		}

		for i := index; i < len(nums); i++ {
			cur = append(cur, nums[i])
			backtracking(i + 1)
			cur = cur[:len(cur)-1]
		}
	}

	backtracking(0)
	return ret
}
