package main

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

	// n := 400000000
	// require.Equal(t, n, len(EatMemory(n)))
}

func EatMemory(n int) (m []int64) {

	m = []int64{}
	for i := 0; i < n; i++ {
		m = append(m, int64(i))
	}

	return
}
