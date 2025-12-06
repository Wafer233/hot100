package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	for {
		n := 0

		_, err := fmt.Fscan(in, &n)
		if err != nil {
			break
		}

		xi := make([]int, n)
		ai := make([]int, n)
		bi := make([]int, n)
		for i := 0; i < n; i++ {
			_, _ = fmt.Fscan(in, &xi[i])
		}
		for i := 0; i < n; i++ {
			_, _ = fmt.Fscan(in, &ai[i])
		}
		for i := 0; i < n; i++ {
			_, _ = fmt.Fscan(in, &bi[i])
		}

		m := 0
		_, _ = fmt.Fscan(in, &m)
		ci := make([]int, m)
		for i := 0; i < m; i++ {
			_, _ = fmt.Fscan(in, &ci[i])
		}

		out := lingyang(xi, ai, bi, ci)

		for index, value := range out {
			if index == 0 {
				fmt.Print(value)
			} else {
				fmt.Print(" ")
				fmt.Print(value)
			}
		}
	}
}

func lingyang(xs, as, bs, cs []int) []int {

	findByA := make(map[int][]*pet)
	findByB := make(map[int][]*pet)

	// 初始化宠物
	for i := 0; i < len(xs); i++ {
		findByA[as[i]] = append(findByA[as[i]], &pet{
			a:      as[i],
			b:      bs[i],
			x:      xs[i],
			status: 1,
		})

		findByB[bs[i]] = append(findByB[bs[i]], &pet{
			a:      as[i],
			b:      bs[i],
			x:      xs[i],
			status: 1,
		})
	}

	ret := make([]int, 0)
	for _, c := range cs {
		exist := false
		petsA := make([]*pet, 0)
		petsB := make([]*pet, 0)
		x1 := int(10e9) + 1
		a1 := -1
		b1 := -1
		a2 := -1
		b2 := -1
		x2 := int(10e9) + 1
		if petsA, exist = findByA[c]; exist {
			sort.Slice(petsA, func(i, j int) bool {
				return petsA[i].x < petsA[j].x
			})

			for j := 0; j < len(petsA); j++ {
				if petsA[j].status == 0 {
					continue
				}
				x1 = min2(x1, petsA[j].x)
				a1 = petsA[j].a
				b1 = petsA[j].b
				break
			}

		}

		if petsB, exist = findByB[c]; exist {
			sort.Slice(petsB, func(i, j int) bool {
				return petsB[i].x < petsB[j].x
			})

			for j := 0; j < len(petsB); j++ {
				if petsB[j].status == 0 {
					continue
				}
				x2 = min2(x2, petsB[j].x)
				a2 = petsB[j].a
				b2 = petsB[j].b
				break
			}
		}

		if x1 == int(10e9)+1 && x2 == int(10e9)+1 {
			ret = append(ret, -1)
			continue
		}

		if x1 < x2 {
			pets := findByA[a1]
			for index, _ := range pets {
				if pets[index].x == x1 {
					pets[index].status = 0
				}
			}
			findByA[a1] = pets

			pets = findByB[b1]
			for index, _ := range pets {
				if pets[index].x == x1 {
					pets[index].status = 0
				}
			}
			findByB[b1] = pets
			ret = append(ret, x1)
		} else {
			pets := findByA[a2]
			for index, _ := range pets {
				if pets[index].x == x2 {
					pets[index].status = 0
				}
			}
			findByA[a2] = pets

			pets = findByB[b2]
			for index, _ := range pets {
				if pets[index].x == x2 {
					pets[index].status = 0
				}
			}
			findByB[b2] = pets
			ret = append(ret, x2)

		}

	}
	return ret

}

func min2(a, b int) int {
	if a < b {
		return a
	}

	return b
}

type pet struct {
	a      int
	b      int
	x      int
	status int
}
