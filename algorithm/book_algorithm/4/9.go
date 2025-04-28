package main

import "fmt"

func partialSum(i, w int, a []int) bool {
	if i == 0 {
		if w == 0 {
			return true
		}
		return false
	}

	if partialSum(i-1, w, a) {
		return true
	}

	if partialSum(i-1, w-a[i-1], a) {
		return true
	}

	return false
}

func main() {
	N, W := 0, 0
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &W)
	a := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a[i])
	}

	if partialSum(N, W, a) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
