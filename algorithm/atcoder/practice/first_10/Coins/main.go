package main

import "fmt"

func main() {
	A, B, C, X := 0, 0, 0, 0

	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	fmt.Scanf("%d", &C)
	fmt.Scanf("%d", &X)

	counter := 0
	for i := 0; i <= A; i++ {
		for j := 0; j <= B; j++ {
			for k := 0; k <= C; k++ {
				ans := 500*i + 100*j + 50*k
				if ans == X {
					counter++
				}
			}
		}
	}
	fmt.Println(counter)
}
