package utils

import (
	"slices"
	"strconv"
)

func MustInt64(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

func MustInt(s string) int {
	n, _ := strconv.ParseInt(s, 10, 32)
	return int(n)
}

func ABS[T int | int64](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func AllEqual(a []int, s int) bool {
	for i := range a {
		if a[i] != s {
			return false
		}
	}
	return true
}

func AllSame(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}

func AllFunc[V any](a []V, f func(V) bool) bool {
	for i := range a {
		if !f(a[i]) {
			return false
		}
	}
	return true
}

func EatMemory(n int) (m []int64) {

	m = []int64{}
	for i := 0; i < n; i++ {
		m = append(m, int64(i))
	}

	return
}

// Returns prime numbers from 1 to n. Sieve of Eratosthenes
func Primes(n int) (result []int) {

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
func LCM(a []int) int {
	res := -1
	primes := Primes(slices.Max(a))[1:]
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

var vis = map[rune]string{
	'|': "│",
	'-': "─",
	'L': "└",
	'J': "┘",
	'7': "┐",
	'F': "┌",
	'.': " ",
	'S': "S",
}

// all recursive combinations of sum and multiplication of nums slice.
// always evaluated left-to-right, not according to precedence rules.
// e.g. for slice 1,2,3 input is all(1, [2,3]), which means:
// 1+2+3 => 3+3 => 6
// 1+2*3 => 3*3 => 9
// 1*2+3 => 2+3 => 5
// 1*2*3 => 2*3 => 6
func all(res uint64, nums []uint64) (out []uint64) {
	if len(nums) == 0 {
		return []uint64{res}
	}

	n := nums[0]
	rest := nums[1:]

	out = append(out, all(res+n, rest)...)
	out = append(out, all(res*n, rest)...)

	return out
}
