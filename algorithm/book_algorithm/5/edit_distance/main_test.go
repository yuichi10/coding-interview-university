package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type tcase struct {
	S      string
	T      string
	answer int
}

func TestDistance(t *testing.T) {

	cases := []tcase{
		{
			"tokyo",
			"kyoto",
			4,
		},
		{
			"competitive",
			"programming",
			10,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := distance(c.S, c.T)
			assert.Equal(t, c.answer, result)
		})
	}
}
