package main

import (
	"fmt"
)

func main() {
	S := ""
	fmt.Scanf("%s", &S)

	alphabet := make(map[string]string)
	for i := 'a'; i <= 'z'; i++ {
		alphabet[string(i)] = string(i)
	}

	for _, s := range S {
		delete(alphabet, string(s))
	}

	for _, s := range alphabet {
		fmt.Println(s)
		return
	}
}
