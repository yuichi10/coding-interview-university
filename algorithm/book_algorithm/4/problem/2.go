package main

import "fmt"

var memo []int

func fib(N int) int {
	if N == 0 {
		return 0
	} else if N == 1 {
		return 0
	} else if N == 2 {
		return 1
	}

	if memo[N] > -1 {
		return memo[N]
	}

	memo[N] = fib(N-1) + fib(N-2) + fib(N-3)
	return memo[N]
}

func main() {
	N := 0
	fmt.Scanf("%d", &N)

	memo = make([]int, N+1)

	for i := range memo {
		memo[i] = -1
	}

	fmt.Println(fib(N))
}
