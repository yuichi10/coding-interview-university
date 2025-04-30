package main

import "fmt"

func main() {
	a, b, c := 0, 0, 0
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d %d", &b, &c)

	s := ""
	fmt.Scanf("%s", &s)

	fmt.Printf("%d %s", (a + b + c), s)
}
