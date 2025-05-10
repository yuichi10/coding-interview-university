package main

import "fmt"

func main() {
	N, W := 0, 0
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &W)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a[i])
	}
	exist := false
	for bit := 0; bit < (1 << N); bit++ {
		sum := 0
		fmt.Println("bit", bit)
		for i := 0; i < N; i++ {
			if bit&(1<<i) > 0 {
				fmt.Println("flag on", i)
				sum += a[i]
			}
		}
		if sum == W {
			exist = true
			break
		}
	}
	if exist {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
