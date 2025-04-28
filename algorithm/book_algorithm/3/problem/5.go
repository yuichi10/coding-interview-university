package main

import (
	"fmt"
)

func main() {
	N := 0
	fmt.Scanf("%d", &N)
	// fmt.Scanf("%d", &v)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a[i])
	}

	or := 0
	for i := 1; i < N; i++ {
		or = or | a[i]
	}
	fmt.Println(or)
	counter := 0
	for i := or; i%2 == 0; i = i / 2 {
		counter++
	}
	fmt.Println(counter)
}
