package main

import (
	"fmt"
)

func main() {
	S := ""
	fmt.Scanf("%s", &S)

	ans := ""
	for _, s := range S {
		if s >= 'A' && s <= 'Z' {
			ans += string(s)
		}
	}

	fmt.Println(ans)
}
