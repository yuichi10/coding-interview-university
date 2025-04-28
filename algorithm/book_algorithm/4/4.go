package main

import "fmt"

func GCD(m, n int) int {
	if n == 0 {
		return m
	}
	return GCD(n, m%n)
}
func main() {
	fmt.Println(GCD(51, 15))
	fmt.Println(GCD(15, 51))

}
