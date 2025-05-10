package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func resolve(N int, order []string) int {
	sum := 0
	for i := 0; i < N; i += 2 {
		num, _ := strconv.Atoi(order[i])
		sum += num
	}
	return sum
}

func main() {
	N := 0
	fmt.Scanf("%d", &N)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	order := strings.Split(scanner.Text(), " ")

	ans := resolve(N, order)

	fmt.Println(ans)
}
