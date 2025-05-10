package main

import (
	"bufio"
	"fmt"
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

func S1() string {
	return S()
}

func S2(sep string) (string, string) {
	sa := SA(sep)
	return sa[0], sa[1]
}

func S3(sep string) (string, string, string) {
	sa := SA(sep)
	return sa[0], sa[1], sa[2]
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

func I3(sep string) (int, int, int) {
	nums := IA(sep)
	return nums[0], nums[1], nums[2]
}

func init() {
	sc = bufio.NewScanner(os.Stdin)
}

func main() {
	R, X := I2(" ")

	if X == 1 {
		if R >= 1600 && R <= 2999 {
			fmt.Println("Yes")
			return
		}
	} else if X == 2 {
		if R >= 1200 && R <= 2399 {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
