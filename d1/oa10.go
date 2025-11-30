package d1

//func main() {
//	for {
//
//		var line string
//		var target int
//		fmt.Scanln(&line)
//
//		part := strings.Split(line, ",")
//		nums := make([]int, len(part))
//		for i := 0; i < len(part); i++ {
//			nums[i], _ = strconv.Atoi(part[i])
//		}
//		fmt.Scan(&target)
//
//		fmt.Println(Search34(nums, target))
//	}
//}

func Search34(nums []int, target int) [2]int {

	firstFind := search33(nums, target, true)
	if firstFind == -1 {
		return [2]int{-1, -1}
	}

	lastFind := search33(nums, target, false)
	return [2]int{firstFind, lastFind}
}

func search33(nums []int, target int, isFirst bool) int {

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			if isFirst {

				if mid == left || nums[mid-1] != target {
					return mid
				} else {
					right = mid - 1
				}

			} else {

				if mid == right || nums[mid+1] != target {
					return mid
				} else {
					left = mid + 1
				}
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
