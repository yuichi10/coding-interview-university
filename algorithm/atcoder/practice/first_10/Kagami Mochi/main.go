package main

import "fmt"

func main() {
	N := 0
	fmt.Scanf("%d", &N)

	d := make(map[int]bool)
	for i := 0; i < N; i++ {
		n := 0
		fmt.Scanf("%d", &n)
		if _, ok := d[n]; !ok {
			d[n] = true
		}
	}

	fmt.Println(len(d))
}
