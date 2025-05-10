package main

import "fmt"

func fib(N int) int {
	if N == 0 {
		return 0
	} else if N == 1 {
		return 0
	} else if N == 2 {
		return 1
	}

	sum := fib(N-1) + fib(N-2) + fib(N-3)
	fmt.Println("N: ", N, " SUM: ", sum)
	return sum
}

func main() {
	N := 0
	fmt.Scanf("%d", &N)

	fmt.Println(fib(N))
}
