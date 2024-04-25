package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func solve(f string) (p1, p2 int) {
	in, _ := os.ReadFile(f)
	lines := strings.Split(strings.TrimSpace(string(in)), "\n")
	m := []int{0}
	for _, l := range lines {
		n, _ := strconv.Atoi(l)
		m = append(m, n)
	}

	slices.Sort(m)

	m = append(m, m[len(m)-1]+3)

	steps := []int{}
	diff1, diff3 := 0, 0
	for i := 1; i < len(m); i++ {
		steps = append(steps, m[i]-m[i-1])
		if m[i]-m[i-1] == 1 {
			diff1++
		}
		if m[i]-m[i-1] == 3 {
			diff3++
		}
	}

	p1 = diff1 * diff3

	// counting sizes of intervals of 1-s (1,1,1..)
	intervals := []int{}
	interval := 0
	for _, n := range steps {
		if n == 3 {
			intervals = append(intervals, interval)
			interval = 0
			continue
		}
		interval++
	}

	p2 = 1
	for _, in := range intervals {
		p2 *= Trib(in + 3)
	}

	return
}

// Tribonacci
func Trib(n int) int {
	switch n {
	case 1:
		return 0
	case 2:
		return 0
	case 3:
		return 1
	}
	return Trib(n-1) + Trib(n-2) + Trib(n-3)
}

func main() {
	start := time.Now()
	fmt.Println("Day 10: Adapter Array")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 2048
	fmt.Println("\tPart Two:", p2) // 1322306994176
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
