package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func valid(n int, m []int) bool {

	for i := 0; i < len(m)-1; i++ {
		for j := i + 1; j < len(m); j++ {
			if m[i]+m[j] == n {
				return true
			}
		}
	}

	return false
}

func findSet(n int, m []int) (found bool, set []int) {
	for i := 0; i < len(m); i++ {
		sum := 0
		set = []int{}
		found = false
		for j := i; j < len(m); j++ {
			sum += m[j]
			set = append(set, m[j])
			if sum == n {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	return
}

func solve(f string) (p1, p2 int) {
	in, _ := os.ReadFile(f)
	lines := strings.Split(strings.TrimSpace(string(in)), "\n")
	m := []int{}
	for _, l := range lines {
		n, _ := strconv.Atoi(l)
		m = append(m, n)
	}

	preamble := 25

	for i := preamble; i < len(m); i++ {
		if !valid(m[i], m[i-preamble:i]) {
			p1 = m[i]
			break
		}
	}

	for i := 0; i < len(m); i++ {
		if found, set := findSet(p1, m[i:]); found {
			p2 = slices.Min(set) + slices.Max(set)
			break
		}
	}

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 09: Encoding Error")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 50047984
	fmt.Println("\tPart Two:", p2) // 5407707
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
