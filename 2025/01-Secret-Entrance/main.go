package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func solve(file string) (p1, p2 int) {
	dial := 50

	f, _ := os.ReadFile(file)
	for l := range strings.SplitSeq(string(f), "\n") {
		n, _ := strconv.Atoi(l[1:])

		for range n {
			if dial == 0 {
				p2++
			}

			switch l[0] {
			case 'L':
				dial -= 1
			case 'R':
				dial += 1
			}

			if dial < 0 {
				dial = 99
			}

			dial %= 100
		}

		if dial == 0 {
			p1++
		}
	}

	return p1, p2
}

func main() {
	start := time.Now()
	fmt.Println("Day 01: Secret Entrance")
	p1, p2 := solve("../aoc-inputs/2025/01/input.txt")
	// p1, p2 := solve("input-pub.txt")
	fmt.Println("\tPart One:", p1) // 1064
	fmt.Println("\tPart Two:", p2) // 6122
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
