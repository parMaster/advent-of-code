package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func read(f string) [][]int {
	res := [][]int{}
	s, _ := os.ReadFile(f)
	lines := strings.Split(string(s), "\n")
	for _, l := range lines {
		line := strings.Split(l, " ")
		levels := []int{}
		for i := range line {
			n, _ := strconv.ParseInt(line[i], 10, 32)
			levels = append(levels, int(n))
		}
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
			rProbe := make([]int, len(r)-1)
			copy(rProbe, r[:i])
			copy(rProbe[i:], r[i+1:])
			if safeReport(rProbe) {
				p2++
				break
			}
		}
	}

	return p1, p2
}

// func almost_linear_solution(f string) (p1, p2 int) {
// 	reports := read(f)
// 	for _, r := range reports {
// 		fmt.Println(r)

// 		fails := 0
// 		failed := true // presume
// 		for {

// 			if len(r) <= 2 || !failed || fails > 1 {
// 				break
// 			}
// 			failed = false
// 			for i := 1; i < len(r)-1; i++ {
// 				if r[i-1] == r[i] || r[i] == r[i+1] ||
// 					abs(r[i]-r[i-1]) > 3 || abs(r[i+1]-r[i]) > 3 ||
// 					sign(r[i]-r[i-1]) != sign(r[i+1]-r[i]) {
// 					r = slices.Delete(r, i, i+1)
// 					fails++
// 					failed = true
// 					break
// 				}
// 			}
// 		}
// 		if fails == 0 {
// 			p1++
// 		}
// 		if fails <= 1 {
// 			p2++
// 		}
// 		if fails >= 1 {
// 			fmt.Println(r, fails)
// 			fmt.Println()
// 		}

// 	}
// 	return p1, p2
// }

func main() {
	start := time.Now()
	fmt.Println("Day 02: Red-Nosed Reports")
	// p1, p2 := solve("input_pub.txt")
	p1, p2 := solve("../aoc-inputs/2024/02/input.txt")
	fmt.Println("\tPart One:", p1) // 680
	fmt.Println("\tPart Two:", p2) // 710
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
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
