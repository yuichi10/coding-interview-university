package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var sc *bufio.Scanner

func S() string {
	sc.Scan()
	return sc.Text()
}

func SA(sep string) []string {
	return strings.Split(S(), sep)
}

func IA(sep string) []int {
	sa := SA(sep)

	var ia []int
	for _, i := range sa {
		n, _ := strconv.Atoi(i)
		ia = append(ia, n)
	}
	return ia
}

func I1() int {
	str := S()
	n, _ := strconv.Atoi(str)
	return n
}

func I2(sep string) (int, int) {
	nums := IA(sep)
	return nums[0], nums[1]
}

func init() {
	sc = bufio.NewScanner(os.Stdin)
}

func main() {
}
