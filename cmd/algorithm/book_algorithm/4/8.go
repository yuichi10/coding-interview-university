package main

import "fmt"

var memo []int

func fibo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	if memo[n] != -1 {
		return memo[n]
	}
	memo[n] = fibo(n-1) + fibo(n-2)
	return memo[n]
}

func main() {
	memo = make([]int, 50)
	for i := range memo {
		memo[i] = -1
	}
	fmt.Println(fibo(49))
}
