package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	Q := 0
	fmt.Scanf("%d", &Q)

	sc := bufio.NewScanner(os.Stdin)

	queue := make([]string, 0)

	for i := 0; i < Q; i++ {
		sc.Scan()
		s := sc.Text()
		ss := strings.Split(s, " ")

		if ss[0] == "2" {
			if len(queue) > 0 {
				fmt.Println(queue[0])
				queue = queue[1:]
			}
		}
		if ss[0] == "1" {
			queue = append(queue, ss[1])
		}
	}
}
