package main

import "fmt"

func main() {
	a, b := 0, 0
	fmt.Scanf("%d %d", &a, &b)

	if (a*b)%2 == 0 {
		fmt.Println("Even")
		return
	}
	fmt.Println("Odd")
}
