package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func chmax[T int | int64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func knapsack(N, W int, vs []int64, ws []int) int64 {
	dmap := make([][]int64, N+1)
	for i := range dmap {
		dmap[i] = make([]int64, W+1)
	}

	for n := 0; n < N; n++ {
		for w := 0; w <= W; w++ {
			if w-ws[n] >= 0 {
				// 選ぶ場合
				dmap[n+1][w] = chmax(dmap[n+1][w], dmap[n][w-ws[n]]+vs[n])
			}
			dmap[n+1][w] = chmax(dmap[n+1][w], dmap[n][w])
		}
	}
	return dmap[N][W]
}

func main() {
	N, W := 0, 0

	fmt.Scanf("%d %d", &N, &W)

	values := make([]int64, N)
	weight := make([]int, N)

	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < N; i++ {
		sc.Scan()
		s := strings.Split(sc.Text(), " ")
		iv, _ := strconv.Atoi(s[0])
		iw, _ := strconv.Atoi(s[1])
		values[i] = int64(iv)
		weight[i] = iw
	}
	fmt.Println(knapsack(N, W, values, weight))
}
