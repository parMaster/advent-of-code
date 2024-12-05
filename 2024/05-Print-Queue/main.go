package main

// https://adventofcode.com/2024/day/5

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func read(f string) (rules [][2]int, reportsSlices [][]int) {
	in, _ := os.ReadFile(f)
	parts := strings.Split(string(in), "\n\n")

	for _, order := range strings.Split(parts[0], "\n") {
		var rule [2]int
		fmt.Sscanf(order, "%d|%d", &rule[0], &rule[1])
		rules = append(rules, rule)
	}

	for _, reportStr := range strings.Split(parts[1], "\n") {
		var pages []int
		json.Unmarshal([]byte("["+reportStr+"]"), &pages)
		reportsSlices = append(reportsSlices, pages)
	}
	return rules, reportsSlices
}

func solve(f string) (p1, p2 int) {
	rules, reports := read(f)

	cmp := func(l, r int) int {
		for _, rule := range rules {
			if l == rule[0] && r == rule[1] {
				return -1
			}
			if r == rule[0] && l == rule[1] {
				return 1
			}
		}
		return 0
	}

	for ri := range reports {
		valid := slices.IsSortedFunc(reports[ri], cmp)
		if valid {
			p1 += reports[ri][len(reports[ri])/2]
		} else {
			slices.SortFunc(reports[ri], cmp)
			p2 += reports[ri][len(reports[ri])/2]
		}
	}

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 05: Print Queue")
	// p1, p2 := solve("input_pub.txt")
	p1, p2 := solve("../aoc-inputs/2024/05/input.txt")
	fmt.Println("\tPart One:", p1) // 5108
	fmt.Println("\tPart Two:", p2) // 7380
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
