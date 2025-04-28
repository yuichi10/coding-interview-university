package main

import "fmt"

func main() {
	K, S := 0, 0
	fmt.Scanf("%d %d", &K, &S)

	counter := 0
	for i := 0; i <= K; i++ {
		for j := 0; j <= K; j++ {
			if i+j > S {
				break
			}
			if S-(i+j) <= K {
				counter++
			}
		}
	}
	fmt.Println(counter)
}
