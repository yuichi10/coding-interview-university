package main

import "fmt"

func main() {
	N, Y := 0, 0
	fmt.Scanf("%d %d", &N, &Y)

	const HIGH = 10000
	const MIDDLE = 5000
	const LOW = 1000

	if Y%1000 != 0 {
		fmt.Println("-1 -1 -1")
		return
	}

	for a := 0; a <= N; a++ {
		for b := 0; b <= N; b++ {
			c := N - (a + b)
			if c >= 0 {
				if HIGH*a+MIDDLE*b+LOW*c == Y {
					fmt.Printf("%d %d %d", a, b, c)
					return
				}
			}
		}
	}
	fmt.Println("-1 -1 -1")
}
