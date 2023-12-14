package main

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValid(t *testing.T) {
	// #.#.### 1,1,3
	// .#...#....###. 1,1,3
	// .#.###.#.###### 1,3,1,6
	// ####.#...#... 4,1,1
	// #....######..#####. 1,6,5
	// .###.##....# 3,2,1

	var re = regexp.MustCompile(`(?m)(#+)`)
	s := "..#...#...###."
	require.EqualValues(t, []string{"#", "#", "###"}, re.FindAllString(s, -1))

	require.True(t, valid("#.#.###", []int{1, 1, 3}))
	require.False(t, valid("##.#.###", []int{1, 1, 3}))
	require.False(t, valid(".#.###", []int{1, 1, 3}))
	require.False(t, valid(".#.###", []int{}))
	require.False(t, valid("", []int{1, 1, 3}))
	require.True(t, valid("..#...#...###.", []int{1, 1, 3}))
	require.False(t, valid("..#?..#...###.", []int{1, 1, 3}))
}

func TestBruteforce(t *testing.T) {
	s := ".###.##.?.?#"
	require.Equal(t, ".###.##.*.?#", s[:strings.Index(s, "?")]+"*"+s[strings.Index(s, "?")+1:])

	require.Equal(t, 1, bf("???.###", []int{1, 1, 3}))
	require.Equal(t, 4, bf(".??..??...?##.", []int{1, 1, 3}))
	require.Equal(t, 10, bf("?###????????", []int{3, 2, 1}))
	require.Equal(t, 1, bf("???.###????.###????.###????.###????.###", []int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3}))
	// require.Equal(t, 1, bf(".??..??...?##.?.??..??...?##.?.??..??...?##.?.??..??...?##.?.??..??...?##.", []int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3}))

	require.Equal(t, 21, p1("input0.txt", false))
	require.Equal(t, 7047, p1("input.txt", false)) // 7047s

}
