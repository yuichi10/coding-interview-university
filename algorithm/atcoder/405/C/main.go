package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc *bufio.Scanner

var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

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
	sc.Buffer(make([]byte, 10000001), 10000000)
}

func main() {
	// N := I1()
	N := 0
	fmt.Scanf("%d", &N)

	sc.Split(bufio.ScanWords)

	as := make([]int64, 0)
	// asSum := make([]int, 0)
	sumNum := int64(0)
	sum := int64(0)
	for i := 0; i < N; i++ {
		sc.Scan()
		numa, _ := strconv.Atoi(sc.Text())
		n := int64(numa)
		sumNum += n
		as = append(as, n)
	}

	for i := 0; i < N-1; i++ {
		sumNum = sumNum - int64(as[i])
		// fmt.Println(as[i])
		// fmt.Println(int64(sumNum) - int64(as[i]))
		// fmt.Printf("%d * %d\n", (as[i]), sumNum)
		sum += int64(as[i]) * sumNum
	}
	// A := SA(" ")

	// // fmt.Println(len(A))
	// for i := 0; i < N-1; i++ {
	// 	for j := i + 1; j < N; j++ {
	// 		a1, _ := strconv.Atoi(A[i])
	// 		a2, _ := strconv.Atoi(A[j])
	// 		sum += int64(a1) * int64(a2)
	// 	}
	// }
	fmt.Println(sum)
}
