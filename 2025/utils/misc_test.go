package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Nifty tricks that always same but different.
// Partially stolen, because I'm not this clever.
func Test_Commons(t *testing.T) {

	s := "50 98 2" // N=3
	trio := [3]int{}
	fmt.Sscanf(s, "%d %d %d", &trio[0], &trio[1], &trio[2])
	require.Equal(t, [3]int{50, 98, 2}, trio)

	s = "seeds: 0 14 55 13 4 55 6 777" //..N
	seeds := []int{}
	json.Unmarshal([]byte("["+strings.Join(strings.Fields(strings.Split(s, ": ")[1]), ",")+"]"), &seeds)
	require.EqualValues(t, []int{0, 14, 55, 13, 4, 55, 6, 777}, seeds)

	var re = regexp.MustCompile(`(?m)(\d+)`)
	s = "0 14 www 13 $$$ 4 55 *** 6 777" //..N
	require.EqualValues(t, []string{"0", "14", "13", "4", "55", "6", "777"}, re.FindAllString(s, -1))
}

func Test_All(t *testing.T) {
	t.Run("AllEqual", func(t *testing.T) {
		require.True(t, AllEqual([]int{0, 0, 0, 0, 0}, 0))
		require.False(t, AllEqual([]int{1, 0, 0, 0, 0}, 0))
		require.False(t, AllEqual([]int{0, 1, 0, 0, 0}, 0))
		require.False(t, AllEqual([]int{1}, 0))
		require.True(t, AllEqual([]int{0}, 0))
	})
	t.Run("AllSame", func(t *testing.T) {
		require.True(t, AllSame([]int{}))
		require.True(t, AllSame([]int{0}))
		require.True(t, AllSame([]int{0, 0}))
		require.False(t, AllSame([]int{0, 1}))
	})

	t.Run("AllFunc", func(t *testing.T) {
		require.True(t, AllFunc([]int{5, 5, 5}, func(a int) bool { return a == 5 }), "every element of slice equals 5")
		require.False(t, AllFunc([]int{5, 5, 1}, func(a int) bool { return a == 5 }), "every element of slice equals 5")

		require.True(t, AllFunc([]int{2, 4, 6, 224}, func(a int) bool { return a%2 == 0 }), "Multiple of 2")
		require.False(t, AllFunc([]int{2, 4, 6, 224}, func(a int) bool { return a%3 == 0 }), "but not all of them are multiple of 3")
		require.True(t, AllFunc([]int{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}, func(a int) bool { return a-0x30 >= 0 && a-0x30 <= 9 }), "digits (0-9) in ASCII?")

		require.True(t, AllFunc([][2]string{{"oh", "yes"}, {"oh", "yes"}}, func(e [2]string) bool { return e[1] == "yes" }), "generic AllFunc")
		require.False(t, AllFunc([][2]string{{"oh", "yes"}, {"oh", "no"}}, func(e [2]string) bool { return e[1] == "yes" }), "generic AllFunc")
	})
}

func Test_EatMemory(t *testing.T) {
	n := 300
	// n := 300 * 1000 * 1000
	// n = 3603973818
	n = 360397381
	require.Equal(t, n, len(EatMemory(n)))
}

// test "all" func
func Test_all(t *testing.T) {

	tests := []struct {
		nums []uint64
		want []uint64
	}{
		{nums: []uint64{1, 2}, want: []uint64{3, 2}},
		{nums: []uint64{1, 2, 0}, want: []uint64{3, 0, 2, 0}},
		{nums: []uint64{1, 2, 3}, want: []uint64{6, 9, 5, 6}},
		{nums: []uint64{10, 2, 3}, want: []uint64{15, 36, 23, 60}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("all(%v)", tt.nums), func(t *testing.T) {
			got := all(tt.nums[0], tt.nums[1:])
			slices.Sort(got)
			slices.Sort(tt.want)
			require.EqualValues(t, tt.want, got)
		})
	}
}
