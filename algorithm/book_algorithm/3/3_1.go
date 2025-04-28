package main

import "fmt"

func main() {
	N := 0
	v := 0
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &v)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a[i])
	}

	found_id := -1
	for i := 0; i < N; i++ {
		if a[i] == v {
			found_id = i
			break
		}
	}
	fmt.Println(found_id)
}
