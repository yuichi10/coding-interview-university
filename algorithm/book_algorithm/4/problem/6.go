package main

import (
	"fmt"
)

var memo map[int]bool

func part(i, w int, a []int) bool {
	if i == 0 {
		if w == 0 {
			return true
		} else {
			return false
		}
	}

	if memo[w] {
		return true
	}

	if part(i-1, w, a) {
		memo[w] = true
		return true
	}

	if part(i-1, w-a[i-1], a) {
		memo[w] = true
		return true
	}

	return false
}

func main() {
	N, W := 0, 0
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &W)
	memo = make(map[int]bool)
	a := make([]int, N)
	for i := range a {
		fmt.Scanf("%d", &a[i])
	}

	if part(N, W, a) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
