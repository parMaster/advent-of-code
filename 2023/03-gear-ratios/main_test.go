package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PartOne(t *testing.T) {

	s := ".0*"

	require.EqualValues(t, 46, '.')
	require.EqualValues(t, 0, s[1]-0x30)

	require.Equal(t, true, isNum(s[1]))
	require.Equal(t, false, isNum(s[0]))
	require.Equal(t, false, isNum(s[2]))

	require.Equal(t, false, isSym(s[0]))
	require.Equal(t, false, isSym(s[1]))
	require.Equal(t, true, isSym(s[2]))

	require.Equal(t, 4361, PartOne("input0.txt"))
	require.Equal(t, 550064, PartOne("input1.txt"))
}

func Test_PartTwo(t *testing.T) {

	require.EqualValues(t, 46, '.')

	s := "..35&..*633.3.3"

	require.Equal(t, -1, catchNumber(s, 0))
	require.Equal(t, 35, catchNumber(s, 2))
	require.Equal(t, 35, catchNumber(s, 3))
	require.Equal(t, -1, catchNumber(s, 7))
	require.Equal(t, 633, catchNumber(s, 8))
	require.Equal(t, 633, catchNumber(s, 9))
	require.Equal(t, 633, catchNumber(s, 10))
	require.Equal(t, -1, catchNumber(s, 11))
	require.Equal(t, 3, catchNumber(s, 12))
	require.Equal(t, -1, catchNumber(s, 13))
	require.Equal(t, 3, catchNumber(s, 14))

	require.Equal(t, 467835, PartTwo("input0.txt"))
	require.Equal(t, 85010461, PartTwo("input1.txt"))
}
