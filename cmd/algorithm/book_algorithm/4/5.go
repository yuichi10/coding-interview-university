package main

import "fmt"

func fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

func fibo_for(n int) int {
	f := make([]int, 50)
	f[0] = 0
	f[1] = 1
	for n := 2; n < 50; n++ {
		f[n] = f[n-1] + f[n-2]
	}
	return f[len(f)-1]
}

func main() {
	fmt.Println(fibo(6))
}
