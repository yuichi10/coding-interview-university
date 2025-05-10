package main

import "fmt"

func sumTotal(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumTotal(n-1)
}

func main() {
	fmt.Println(sumTotal(5))
}
