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
		n = n % 100

		switch l[0] {
		case 'L':
			dial -= n
			if dial < 0 {
				dial = 100 + dial
			}
		case 'R':
			dial += n
			dial = dial % 100
		}

		if dial == 0 {
			p1++
		}

		// fmt.Println(n, dial)
	}

	return p1, p2
}

func main() {
	start := time.Now()
	fmt.Println("Day 01: Secret Entrance")
	p1, p2 := solve("../aoc-inputs/2025/01/input.txt")
	// p1, p2 := solve("input-pub.txt")
	fmt.Println("\tPart One:", p1) // 1064
	fmt.Println("\tPart Two:", p2) //
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
