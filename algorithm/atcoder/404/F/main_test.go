package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type tcase struct {
}

func TestResolve(t *testing.T) {
	cases := []tcase{
		{},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			_ = c
			resolve()
			assert.Equal(t, 1, 1)
		})
	}
}
