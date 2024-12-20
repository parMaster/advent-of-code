package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"encoding/json"
)

func read(f string) [][]int {
	res := [][]int{}
	s, _ := os.ReadFile(f)
	lines := strings.Split(string(s), "\n")
	for _, line := range lines {
		var levels []int
		json.Unmarshal([]byte("["+strings.ReplaceAll(line, " ", ",")+"]"), &levels)
		res = append(res, levels)
	}
	return res
}

func safeReport(r []int) bool {
	safe := true
	if r[0] == r[1] {
		return false
	}
	sign := (r[1] - r[0]) / abs(r[1]-r[0])
	for i := range len(r) - 1 {
		d := r[i+1] - r[i]
		if d == 0 || abs(d) > 3 || sign != (d/abs(d)) {
			safe = false
			break
		}
	}
	return safe
}

func solve(f string) (p1, p2 int) {
	reports := read(f)
	for _, r := range reports {
		safe := safeReport(r)
		if safe {
			p1++
		}

		for i := range r {
			if safeReport(slices.Delete(slices.Clone(r), i, i+1)) {
				p2++
				break
			}
		}
	}

	return p1, p2
}

func abs[T int | int64](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func sign[T int | int64](n T) T {
	return n / abs(n)
}

func main() {
	start := time.Now()
	fmt.Println("Day 02: Red-Nosed Reports")
	// p1, p2 := solve("input_pub.txt")
	p1, p2 := solve("../aoc-inputs/2024/02/input.txt")
	fmt.Println("\tPart One:", p1) // 680
	fmt.Println("\tPart Two:", p2) // 710
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
