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
//		fmt.Println(Search33(nums, target))
//	}
//}

func Search33(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[right] < nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}

		}
	}
	return -1
}
