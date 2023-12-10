package main

import (
	"encoding/json"
	"fmt"
	"image"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// directions   ↖ ↑ ↗ → ↘ ↓ ↙ ←
// directions	0 1 2 3 4 5 6 7
var dir = []image.Point{{-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}}

var xyDir = [][]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}

var diagDir = [][]int{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}}

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

func absPoint(d image.Point) image.Point {
	if d.X < 0 {
		d.X = -d.X
	}

	if d.Y < 0 {
		d.Y = -d.Y
	}
	return d
}

func mustInt64(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func Test_EatMemory(t *testing.T) {
	n := 300 * 1000 * 1000
	// n = 3603973818
	require.Equal(t, n, len(EatMemory(n)))
}

func EatMemory(n int) (m []int64) {

	m = []int64{}
	for i := 0; i < n; i++ {
		m = append(m, int64(i))
	}

	return
}

// Returns prime numbers from 1 to n. Sieve of Eratosthenes
func primes(n int) (result []int) {

	a := []bool{}
	for i := 0; i <= n; i++ {
		a = append(a, true)
	}

	for i := 2; i <= n; i++ {
		if a[i] {
			j := i * i
			for j <= n {
				a[j] = false
				j += i
			}
		}
	}

	for i := 1; i <= n; i++ {
		if a[i] {
			result = append(result, i)
		}
	}
	return
}

// Returns Least Common Multiple of the integers slice
// LCM with division method:
// Divide numbers by prime numbers as long as at least one of the
// numbers is evenly divisible by a prime number.
func lcm(a []int) int {
	res := -1
	primes := primes(slices.Max(a))[1:]
	dividers := []int{}
	for {
		nodiv := true
		for _, p := range primes {
			nodiv := true
			for i := range a {
				if a[i]%p == 0 && a[i] > 1 {
					a[i] /= p
					nodiv = false
				}
			}
			if !nodiv {
				dividers = append(dividers, p)
			}
		}
		if nodiv {
			if slices.Max(a) == 1 {
				for i := range dividers {
					res *= dividers[i]
				}
				return -res
			}
			return res
		}
	}
}