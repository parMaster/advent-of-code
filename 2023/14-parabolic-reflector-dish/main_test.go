package main

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSort(t *testing.T) {

	s := []byte("O..O..")
	slices.Sort(s)
	slices.Reverse(s)

	sorted := []byte("OO....")
	require.Equal(t, sorted, s)

	reduced := []byte("OO")
	s = slices.DeleteFunc(s, func(e byte) bool { return e == byte('.') })
	require.Equal(t, reduced, s)
	// gotcha
	// len(s)

	s = []byte(".O.#O..O")
	sl := len(s)
	for _, part := range strings.Split(string(s), "#") {
		pp := []byte(part)
		slices.Sort(pp)
		slices.Reverse(pp)
		s = slices.DeleteFunc(pp, func(e byte) bool { return e == byte('.') })
		rocks := len(s)
		dots := len(pp) - rocks
		fmt.Println(s, sl, rocks, dots)
	}

}
