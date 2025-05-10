package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	N := 0
	fmt.Scanf("%d", &N)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	AS := strings.Split(sc.Text(), " ")

	num := 0b0
	for i := 0; i < N; i++ {
		n, _ := strconv.Atoi(AS[i])
		num = num | n
	}

	counter := 0
	for ; num%2 == 0; num /= 2 {
		counter++
	}
	fmt.Println(counter)
}
