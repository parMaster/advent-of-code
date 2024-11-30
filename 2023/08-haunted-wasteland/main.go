package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var dir = map[rune]int{'L': 0, 'R': 1}

func input(f string) (string, map[string][2]string) {
	input, _ := os.ReadFile(f)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	route := lines[0]
	m := map[string][2]string{}
	for _, l := range lines[2:] {
		m[l[0:3]] = [2]string{l[7:10], l[12:15]}
	}
	return route, m
}

func PartOne(f string) int {
	route, m := input(f)

	i := 0
	c := "AAA" // current, start
	for c != "ZZZ" {
		for _, move := range []rune(route) {
			c = m[c][dir[move]]
			i++
			if c == "ZZZ" {
				break
			}
		}
	}
	return i
}

func PartTwo(f string, lcm func([]int) int) int {
	results := []int{}
	route, m := input(f)
	for c := range m {
		if c[2] == 'A' {
			i := 0
			// c := m k // current, start
			for c[2] != 'Z' {
				for _, move := range []rune(route) {
					c = m[c][dir[move]]
					i++
					if c[2] == 'Z' {
						break
					}
				}
			}
			results = append(results, i)
		}
	}

	return lcm(results)
}

func main() {
	fmt.Println("Day 8: Haunted Wasteland")
	fmt.Println("\tPart One:", PartOne("../aoc-inputs/2023/08/input1.txt"))      // 12737
	fmt.Println("\tPart Two:", PartTwo("../aoc-inputs/2023/08/input1.txt", lcm)) // 9064949303801
	if slices.Index(os.Args[1:], "--bruteforce") != -1 {
		fmt.Println("\tPart Two (brute forcing LCM):", PartTwo("../aoc-inputs/2023/08/input1.txt", lcm_smartbf)) // 9064949303801
	}
}

// Sieve of Eratosthenes
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

// bruteforcing LCM, just for fun... ~1B/10sec
func lcm_bf(a []int) int {
	mr := slices.Max(a)
	for {
		d := true
		for i := range a {
			if mr%a[i] != 0 {
				d = false
				break
			}
		}

		if d == true {
			return mr
		}
		mr++
		if mr%1000000000 == 0 {
			fmt.Println(mr)
		}
	}
}

// smart bruteforcing LCM
func lcm_smartbf(a []int) int {

	n := 1
	for i := range a {
		n *= a[i]
	}
	mx := slices.Max(a)

	i := 0
	l := len(a)
	for i = 1; i <= n; i++ {
		div := 0
		for j := range a {
			if (mx*i)%a[j] == 0 {
				div++
			}
		}
		if div == l {
			return mx * i
		}
		// debug print
		// if i%100000000 == 0 {
		// 	fmt.Println(i*mx, "of", n, "checked")
		// }
	}

	return -0
}
