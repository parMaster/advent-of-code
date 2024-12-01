package main

import (
	"fmt"
	"slices"
	"time"

	"github.com/parMaster/advent-of-code/2024/utils"
)

func solve(f string) (p1, p2 int64) {
	numbers := utils.ReadLines(f)

	var left, right []int64
	for i := 0; i < len(numbers); i++ {
		left = append(left, utils.MustInt64(numbers[i]))
		i++
		right = append(right, utils.MustInt64(numbers[i]))
	}

	// Part One
	slices.Sort(left)
	slices.Sort(right)
	for i := range left {
		p1 += utils.ABS(int64(left[i] - right[i]))
	}

	// Part Two
	for _, l := range left {
		var cnt int64
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
