package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/maps"
)

type set struct {
	lines []string
	cols  []string
}

func read(f string) (sets []set) {
	in, _ := os.ReadFile(f)
	for _, m := range strings.Split(string(in), "\n\n") {
		lines := []string{}
		cols := []string{}

		lines = strings.Split(strings.TrimSpace(string(m)), "\n")

		for i := range lines {
			for j, c := range lines[i] {
				if len(cols) == 0 {
					cols = make([]string, len(lines[i]))
				}
				cols[j] += string(c)
			}
		}

		sets = append(sets, set{lines: lines, cols: cols})
	}
	return sets
}

func findMirrors(s []string, smudgesAllowed int) map[int]bool {
	mirrors := map[int]bool{}
	for i := range s {
		if reflected(s[0:i], smudgesAllowed) {
			mirrors[(i+1)/2] = true
			// fmt.Println(0, "x", i)
		}

		if reflected(s[i:], smudgesAllowed) {
			mirrors[len(s)-(len(s)-i)/2] = true
			// fmt.Println(i, "x", len(s))
		}

	}
	return mirrors
}

// returns true if a slice of lines can be folded in the middle
// considering number of smudges that allowed for the whole pattern (O..n)
func reflected(lines []string, smudgesAllowed int) bool {
	if len(lines)%2 != 0 || len(lines) == 0 {
		return false
	}

	smudges := 0
	for i := range lines[:len(lines)/2] {
		l1 := lines[i]
		l2 := lines[len(lines)-1-i]
		for li := range l1 {
			if l1[li] != l2[li] {
				smudges++
			}
		}
	}

	return smudgesAllowed == smudges
}

func solve(f string, smudgesAllowed int) int {
	sets := read(f)
	sum := 0
	for _, s := range sets {
		m := findMirrors(s.lines, smudgesAllowed)
		for _, r := range maps.Keys(m) {
			sum += r * 100
		}
		m = findMirrors(s.cols, smudgesAllowed)
		for _, r := range maps.Keys(m) {
			sum += r
		}
	}
	return sum
}

func main() {
	fmt.Println("Day 13: Point of Incidence")
	fmt.Println("\tPart One:", solve("../aoc-inputs/2023/13/input.txt", 0)) // 36041
	fmt.Println("\tPart Two:", solve("../aoc-inputs/2023/13/input.txt", 1)) // 35915
}
