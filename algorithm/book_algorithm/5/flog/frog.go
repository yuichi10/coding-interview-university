package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func minDistance(N, step, cost, lowest int, hights []int) int {
// 	if cost > lowest {
// 		return lowest
// 	}
// 	if step == N-1 {
// 		if lowest > cost {
// 			lowest = cost
// 		}
// 		return lowest
// 	}

// 	c := hights[step+1] - hights[step]
// 	if c < 0 {
// 		c *= -1
// 	}
// 	lowest = minDistance(N, step+1, cost+c, lowest, hights)

// 	if step+2 <= N-1 {
// 		c = hights[step+2] - hights[step]
// 		if c < 0 {
// 			c *= -1
// 		}
// 		lowest = minDistance(N, step+2, cost+c, lowest, hights)
// 	}
// 	return lowest
// }

func minDistance(N int, heights []int) int {

	distance := make([]int, N)
	distance[0] = 0
	for i := 1; i < N; i++ {
		if i == 1 {
			c := heights[i-1] - heights[i]
			if c < 0 {
				c *= -1
			}
			distance[i] = c
			continue
		} else {
			c1 := heights[i-1] - heights[i]
			if c1 < 0 {
				c1 *= -1
			}
			c2 := heights[i-2] - heights[i]
			if c2 < 0 {
				c2 *= -1
			}

			d1 := distance[i-1] + c1
			d2 := distance[i-2] + c2
			if d1 > d2 {
				distance[i] = d2
			} else {
				distance[i] = d1
			}
		}
	}
	return distance[N-1]
}

func main() {
	N := 0
	fmt.Scanf("%d", &N)

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	heights := make([]int, N)

	for i := range inputs {
		heights[i], _ = strconv.Atoi(inputs[i])
	}

	// min := minDistance(N, 0, 0, math.MaxInt, heights)

	min := minDistance(N, heights)

	fmt.Println(min)
}
