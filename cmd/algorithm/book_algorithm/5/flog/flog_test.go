package main

import (
	"fmt"
	"testing"
)

type minDistanceCase struct {
	N       int
	heights []int
	answer  int
}

func TestMain(t *testing.T) {
	cases := []minDistanceCase{
		{
			6,
			[]int{30, 10, 60, 10, 60, 50},
			40,
		},
		{
			2,
			[]int{10, 10},
			0,
		},
		{
			4,
			[]int{10, 30, 40, 20},
			30,
		},
		{
			5,
			[]int{10, 10, 10, 10, 10},
			0,
		},
		{
			5,
			[]int{10, 10, 9, 10, 10},
			0,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%d", c.answer), func(t *testing.T) {
			result := minDistance(c.N, c.heights)
			if result != c.answer {
				t.Errorf("wrong %d %d", c.answer, result)
			}
		})
	}
}
