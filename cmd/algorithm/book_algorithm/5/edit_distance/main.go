package main

// https://atcoder.jp/contests/tessoku-book/tasks/tessoku_book_cs

import (
	"fmt"
	"math"
	"strings"
)

func chmin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func distance(S, T string) int {
	s := strings.Split(S, "")
	t := strings.Split(T, "")
	dp := make([][]int, len(S)+1)
	for i := range dp {
		dp[i] = make([]int, len(T)+1)
	}

	for i := 0; i < len(S)+1; i++ {
		for j := 0; j < len(T)+1; j++ {
			dp[i][j] = math.MaxInt
			if i == 0 && j == 0 {
				dp[i][j] = 0
			}
			if i > 0 && j > 0 {
				if s[i-1] == t[j-1] {
					dp[i][j] = chmin(dp[i][j], dp[i-1][j-1])
				} else {
					dp[i][j] = chmin(dp[i][j], dp[i-1][j-1]+1)
				}
			}
			if i > 0 {
				dp[i][j] = chmin(dp[i][j], dp[i-1][j]+1)
			} else if j > 0 {
				dp[i][j] = chmin(dp[i][j], dp[i][j-1]+1)
			}
		}
	}
	return dp[len(S)][len(T)]
}

func main() {
	S, T := "", ""

	fmt.Scanf("%s", &S)
	fmt.Scanf("%s", &T)
	ans := distance(S, T)
	fmt.Println(ans)
}
