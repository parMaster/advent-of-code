package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func solve(f string) (p1, p2 int) {
	s, _ := os.ReadFile(f)
	var left, right []int
	for _, line := range strings.Split(strings.TrimSpace(string(s)), "\n") {
		var l, r int
		fmt.Sscanf(line, "%d   %d", &l, &r)
		left = append(left, l)
		right = append(right, r)
	}

	// Part One
	slices.Sort(left)
	slices.Sort(right)
	for i := range left {
		if left[i] > right[i] {
			p1 += left[i] - right[i]
		} else {
			p1 += right[i] - left[i]
		}
	}

	// Part Two
	for _, l := range left {
		var cnt int
		for _, r := range right {
			if l == r {
				cnt++
			}
		}
		p2 += l * cnt
	}

	return p1, p2
}

func main() {
	start := time.Now()
	fmt.Println("Day 01: Historian Hysteria")
	p1, p2 := solve("../aoc-inputs/2024/01/input.txt")
	fmt.Println("\tPart One:", p1) // 1941353
	fmt.Println("\tPart Two:", p2) // 22539317
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
