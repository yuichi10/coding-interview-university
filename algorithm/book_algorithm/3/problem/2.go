package main

import "fmt"

func main() {
	N, v := 0, 0
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &v)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a[i])
	}

	count := 0
	for i := 0; i < N; i++ {
		if a[i] == v {
			count++
		}
	}
	fmt.Println(count)
}
