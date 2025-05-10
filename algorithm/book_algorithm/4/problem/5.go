package main

import (
	"fmt"
)

var counter int

func f753(K, cur, use int, counter *int) {
	if cur > K {
		return
	}
	if use == 0b111 {
		*counter++
	}
	f753(K, cur*10+7, use|0b100, counter)
	f753(K, cur*10+5, use|0b010, counter)
	f753(K, cur*10+3, use|0b001, counter)
}

func main() {
	K := 0
	fmt.Scanf("%d", &K)

	counter := 0
	f753(K, 0, 0, &counter)
	fmt.Println(counter)
}
