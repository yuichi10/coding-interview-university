package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func resolve(N, M, Q int, query []string) {
// 	data := make([][]int, N+1)
// 	for i := 0; i < N; i++ {
// 		data[i] = make([]int, M+1)
// 	}

// 	memo := make(map[string]bool)

// 	for _, q := range query {
// 		qs := strings.Split(q, " ")
// 		q1, _ := strconv.Atoi(qs[0])
// 		q2, _ := strconv.Atoi(qs[1])
// 		q3 := -1
// 		if q1 != 2 {
// 			q3, _ = strconv.Atoi(qs[2])
// 		}
// 		if q1 != 3 {
// 			_, ok := memo[q]
// 			if ok {
// 				continue
// 			}
// 		}

// 		if q1 == 1 {
// 			memo[q] = true
// 			data[q2-1][q3-1] = 1
// 		}
// 		if q1 == 2 {
// 			memo[q] = true
// 			for i := 0; i < M; i++ {
// 				data[q2-1][i] = 1
// 			}
// 		}
// 		if q1 == 3 {
// 			if data[q2-1][q3-1] == 1 {
// 				fmt.Println("Yes")
// 			} else {
// 				fmt.Println("No")
// 			}
// 		}
// 	}
// }

func resolve2(N, M, Q int, query []string) {
	data := make(map[string]bool)

	memo := make(map[string]bool)

	for _, q := range query {
		qs := strings.Split(q, " ")
		q1, _ := strconv.Atoi(qs[0])
		q2, _ := strconv.Atoi(qs[1])
		q3 := -1
		if q1 != 2 {
			q3, _ = strconv.Atoi(qs[2])
		}
		if q1 != 3 {
			_, ok := memo[q]
			if ok {
				continue
			}
		}

		if q1 == 1 {
			memo[q] = true
			data[fmt.Sprintf("%d-%d", q2-1, q3-1)] = true
		}
		if q1 == 2 {
			memo[q] = true
			data[fmt.Sprintf("%d-all", q2-1)] = true
		}
		if q1 == 3 {
			_, ok := data[fmt.Sprintf("%d-all", q2-1)]
			if ok {
				fmt.Println("Yes")
				continue
			}
			_, ok = data[fmt.Sprintf("%d-%d", q2-1, q3-1)]
			if ok {
				fmt.Println("Yes")
				continue
			}
			fmt.Println("No")
		}
	}
}

func main() {
	N, M, Q := 0, 0, 0
	fmt.Scanf("%d %d %d", &N, &M, &Q)

	sc := bufio.NewScanner(os.Stdin)
	query := make([]string, Q)
	for i := 0; i < Q; i++ {
		sc.Scan()
		query[i] = sc.Text()
	}

	resolve2(N, M, Q, query)
}
