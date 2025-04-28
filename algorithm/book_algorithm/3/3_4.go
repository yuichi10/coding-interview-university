package main

import (
	"fmt"
	"math"
)

func main() {
	N, K := 0, 0
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &K)
	a, b := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &b[i])
	}
	min_value := math.MaxInt
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if a[i]+b[j] < K {
				continue
			}
			if a[i]+b[j] < min_value {
				min_value = a[i] + b[j]
			}
		}
	}
	fmt.Println(min_value)
}
