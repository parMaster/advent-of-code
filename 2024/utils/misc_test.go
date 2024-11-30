package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Nifty tricks that always same but different
// partially stolen, because I'm not this clever
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
	require.True(t, AllEqual([]int{0, 0, 0, 0, 0}, 0))
	require.False(t, AllEqual([]int{1, 0, 0, 0, 0}, 0))
	require.False(t, AllEqual([]int{0, 1, 0, 0, 0}, 0))
	require.False(t, AllEqual([]int{1}, 0))
	require.True(t, AllEqual([]int{0}, 0))

	require.True(t, AllSame([]int{0}))
	require.True(t, AllSame([]int{0, 0}))
	require.False(t, AllSame([]int{0, 1}))

	require.True(t, AllFunc([]int{5, 5, 5}, func(a int) bool { return a == 5 }), "every element of slice equals 5")
	require.False(t, AllFunc([]int{5, 5, 1}, func(a int) bool { return a == 5 }), "every element of slice equals 5")

	require.True(t, AllFunc([]int{2, 4, 6, 224}, func(a int) bool { return a%2 == 0 }), "Multiple of 2")
	require.False(t, AllFunc([]int{2, 4, 6, 224}, func(a int) bool { return a%3 == 0 }), "but not all of them are multiple of 3")
	require.True(t, AllFunc([]int{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}, func(a int) bool { return a-0x30 >= 0 && a-0x30 <= 9 }), "digits (0-9) in ASCII?")

	require.True(t, AllFunc([][2]string{{"oh", "yes"}, {"oh", "yes"}}, func(e [2]string) bool { return e[1] == "yes" }), "generic AllFunc")
	require.False(t, AllFunc([][2]string{{"oh", "yes"}, {"oh", "no"}}, func(e [2]string) bool { return e[1] == "yes" }), "generic AllFunc")
}

func Test_EatMemory(t *testing.T) {
	n := 300
	// n := 300 * 1000 * 1000
	// n = 3603973818
	require.Equal(t, n, len(EatMemory(n)))
}
