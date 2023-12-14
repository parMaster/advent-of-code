package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSort(t *testing.T) {

	s := []byte("O..O..")

	slices.Sort(s)
	slices.Reverse(s)
	sorted := []byte("OO....")
	require.Equal(t, sorted, s)

	s = slices.DeleteFunc(s, func(e byte) bool { return e == byte('.') })
	reduced := []byte("OO")
	require.Equal(t, reduced, s)

	s = []byte(".O.#O..O")
	//          87654321
	//.O.#O..O->O..#OO.. = 8+4+3 = 15
	require.Equal(t, 15, shift(s))

}
