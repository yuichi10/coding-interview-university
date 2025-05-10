package main

import (
	"fmt"
	"strings"
)

func main() {
	S := ""

	fmt.Scanf("%s", &S)
	fmt.Println(strings.Count(S, "1"))
}
