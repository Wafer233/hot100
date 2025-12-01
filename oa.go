package goland_oa

func Solution(A []int) int {
	// Implement your solution here

	mapping := make(map[int]bool)

	for _, value := range A {
		if value <= 0 {
			continue
		}
		mapping[value] = true
	}

	for i := 1; i < len(A); i++ {
		if _, exist := mapping[i]; !exist {
			return i
		}
	}
	return -1
}
