package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	for {
		group := 0
		_, err := fmt.Fscan(in, &group)
		if err != nil {
			break
		}
		//ret := make([]string,0)

		for i := 0; i < group; i++ {
			m := 0
			n := 0
			_, _ = fmt.Fscan(in, &m, &n)

			str := ""
			_, _ = fmt.Fscan(in, &str)
			as := make([]int, n)
			for j := 0; j < n; j++ {
				_, _ = fmt.Fscan(in, &as[j])
			}

			randoms := rand.Int31()
			if randoms%2 == 0 {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}

		}

	}
}
