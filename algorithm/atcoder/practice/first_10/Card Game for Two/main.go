package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	N := 0
	fmt.Scanf("%d", &N)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	as := strings.Split(sc.Text(), " ")
	a := make([]int, len(as))

	for i, s := range as {
		n, _ := strconv.Atoi(s)
		a[i] = n
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	diff := 0
	for i := 0; i < N-1; i += 2 {
		diff += a[i] - a[i+1]
	}

	if N%2 == 1 {
		diff += a[N-1]
	}
	fmt.Println(diff)
}
