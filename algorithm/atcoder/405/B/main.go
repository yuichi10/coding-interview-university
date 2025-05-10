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

func checkFill(nums map[int]int) bool {
	for _, v := range nums {
		if v == 0 {
			return false
		}
	}
	return true
}

func main() {
	N, M := I2(" ")

	A := IA(" ")

	nums := make(map[int]int)
	for i := 1; i <= M; i++ {
		nums[i] = 0
	}
	for _, a := range A {
		if a <= M {
			nums[a]++
		}
	}

	count := 0
	for i := 0; i < N; i++ {
		if !checkFill(nums) {
			break
		}
		n := A[len(A)-1]
		A = A[:len(A)-1]
		if n <= M {
			nums[n]--
		}
		count++
	}
	fmt.Println(count)
}
