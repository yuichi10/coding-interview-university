package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	ss := strings.Split(sc.Text(), " ")

	N, _ := strconv.Atoi(ss[0])
	A, _ := strconv.Atoi(ss[1])
	B, _ := strconv.Atoi(ss[2])

	sum := 0
	for i := 0; i < N; i++ {
		n := 0
		num := i + 1

		for ; num > 0; num /= 10 {
			n += num % 10
		}
		if n >= A && n <= B {
			sum += i + 1
		}
	}
	fmt.Println(sum)
}
