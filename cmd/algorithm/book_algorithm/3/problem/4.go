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

	b1 := -1
	b2 := -1
	for i := 0; i < N; i++ {
		if a[i] > b1 {
			b2 = b1
			b1 = a[i]
			continue
		}
		if a[i] > b2 {
			b2 = a[i]
		}
	}
	fmt.Println(b2)
}
