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

type uf struct {
	per []int
	siz []int
}

func NewUF(n int) *uf {
	per := make([]int, n)
	siz := make([]int, n)
	for i := range per {
		per[i] = -1
		siz[i] = 1
	}
	return &uf{
		per, siz,
	}
}

func (u *uf) root(x int) int {
	if u.per[x] == -1 {
		return x
	}
	u.per[x] = u.root(u.per[x])
	return u.per[x]
}

func (u *uf) union(x, y int) {
	x = u.root(x)
	y = u.root(y)

	if x == y {
		return
	}

	if u.siz[x] < u.siz[y] {
		a := x
		x = y
		y = a
	}

	u.per[y] = x
	u.siz[x] += u.siz[y]
}

func (u *uf) size(x int) int {
	return u.siz[u.root(x)]
}

func init() {
	sc = bufio.NewScanner(os.Stdin)
}

func main() {
	N, M := I2(" ")

	vec := make(map[int]int)

	uf := NewUF(N + 1)
	input := -1
	for i := 0; i < M; i++ {
		a, b := I2(" ")
		vec[a]++
		vec[b]++
		uf.union(a, b)
		input = a
	}

	for _, v := range vec {
		if v != 2 {
			fmt.Println("No")
			return
		}
	}
	if uf.size(input) == N {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
