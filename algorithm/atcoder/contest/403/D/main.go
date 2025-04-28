package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func resolve(N, D int, AS []string) {
	deleteList := make([]int, 0)
}

func main() {
	N, D := 0, 0

	fmt.Scanf("%d %d", &N, &D)

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	as := strings.Split(sc.Text(), " ")

	resolve(N, D, as)
}
