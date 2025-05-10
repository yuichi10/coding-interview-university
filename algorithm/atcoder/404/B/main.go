package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func findMissingCount(N int, S, T [][]string) int {
	counter := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if S[i][j] != T[i][j] {
				counter++
			}
		}
	}
	return counter
}

func rotateRight(N int, S [][]string) [][]string {
	newS := make([][]string, N)

	for i := 0; i < N; i++ {
		newS[i] = make([]string, N)
		for j := 0; j < N; j++ {
			newS[i][j] = S[N-1-j][i]
		}
	}

	return newS
}

func calculate(N int, S, T [][]string) int {
	counters := make([]int, 4)

	for i := 0; i < 4; i++ {
		counters[i] = i
		if i != 0 {
			S = rotateRight(N, S)
		}
		counters[i] += findMissingCount(N, S, T)
	}

	min := math.MaxInt
	for _, c := range counters {
		if c < min {
			min = c
		}
	}
	return min
}

func main() {

	N := 0
	fmt.Scanf("%d", &N)

	S := make([][]string, N)
	T := make([][]string, N)
	for i := range S {
		S[i] = make([]string, N)
		T[i] = make([]string, N)
	}

	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < N; i++ {
		sc.Scan()

		for j, char := range sc.Text() {
			S[i][j] = string(char)
		}
	}
	for i := 0; i < N; i++ {
		sc.Scan()
		for j, char := range sc.Text() {
			T[i][j] = string(char)
		}
	}
	fmt.Println(calculate(N, S, T))
}
