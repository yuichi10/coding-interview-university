package main

import (
	"fmt"
	"strconv"
	"strings"
)

var maps []map[int]bool

func setmaps(A, B int) {
	if (len(maps)) < 0 {
		maps[0]
	}
	findid := ""
	for _, m := range maps {
		if _, ok := m[A]; ok {
			m[B] = true
			m[A] = true
		}
		if _, ok := m[B]; ok {
			m[B] = true
			m[A] = true
		}
	}
	if findid == "" {
		m := make(map[int]bool)
		m[A] = true
		m[B] = true
		maps = append(maps, m)
	}

	if len(findid) > 1 {
		ids := strings.Split(findid, "")

		first_id := ids[0]
		fid, _ := strconv.Atoi(first_id)
		for _, id := range ids {
			iid, _ := strconv.Atoi(id)
			if id == first_id {
				continue
			} else {
				for k, v := range maps[iid] {
					maps[fid][k] = v
				}
			}

		}
	}
}

func main() {
	N, M := 0, 0
	fmt.Scanf("%d %d", &N, &M)

	input := make(map[string]bool)

	output := make(map[int]int)

	check := make(map[int]bool)

	maps = make([]map[int]bool, 0)

	no := false

	for i := 0; i < M; i++ {
		A, B := 0, 0
		fmt.Scanf("%d %d", &A, &B)
		if _, ok := input[fmt.Sprintf("%d+%d", A, B)]; !ok && !no {
			if i == 0 {
				m := make(map[int]bool, 0)
				m[A] = true
				m[B] = true
				maps[0] = m
			}
			input[fmt.Sprintf("%d+%d", A, B)] = true
			input[fmt.Sprintf("%d+%d", B, A)] = true
			output[A] += 1
			output[B] += 1
			if output[A] > 2 || output[B] > 2 {
				fmt.Println("No")
				no = true
			}
			if output[A] == 2 {
				check[A] = true
			}
			if output[B] == 2 {
				check[B] = true
			}
		}
	}

	if len(check) == M {
		fmt.Println("Yes")
	} else if !no {
		fmt.Println("No")
	}
}
