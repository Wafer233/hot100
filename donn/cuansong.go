package main

import "fmt"

func main() {
	fmt.Println(CuanSongDai(5, 24))
}

func CuanSongDai(a, b int) int {

	queue := []Node{Node{x: a}}
	visited := map[int]bool{}
	visited[a] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.x == b {
			return cur.step
		}

		dirs := []int{
			cur.x * 2,
			cur.x + 1,
			cur.x / 2,
		}

		for _, next := range dirs {
			if next < 0 || next > 2000 || visited[next] {
				continue
			}
			visited[next] = true
			queue = append(queue, Node{x: next, step: cur.step + 1})
		}

	}
	return -1
}

type Node struct {
	x    int
	step int
}
