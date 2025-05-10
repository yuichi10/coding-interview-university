package main

import (
	"fmt"
	"strings"
)

func resolve(T, U string) string {
	u := strings.Split(U, "")

	pattern := make([][]string, len(T))

	for p := range pattern {
		pattern[p] = make([]string, len(U))
	}

	for i := 0; i < len(U); i++ {
		pattern := T
		iPattern := strings.Replace(pattern, "?", u[i], 1)
		for j := i; j < len(U); j++ {
			jPattern := strings.Replace(iPattern, "?", u[j], 1)
			for k := j; k < len(U); k++ {
				kPattern := strings.Replace(jPattern, "?", u[k], 1)
				for l := k; l < len(U); l++ {
					lPattern := strings.Replace(kPattern, "?", u[l], 1)
					if strings.Contains(lPattern, U) {
						return "Yes"
					}
				}
			}
		}
	}
	return "No"
}

func main() {
	T, U := "", ""
	fmt.Scanf("%s", &T)
	fmt.Scanf("%s", &U)

	fmt.Println(resolve(T, U))
}
