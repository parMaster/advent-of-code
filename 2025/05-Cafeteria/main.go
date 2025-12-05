package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func solve(file string) (p1, p2 int) {
	in, _ := os.ReadFile(file)
	pts := strings.Split(string(in), "\n\n")

	ranges := make([][2]int, len(strings.Split(pts[0], "\n")))
	for i, r := range strings.Fields(pts[0]) {
		var from, to int
		fmt.Sscanf(r, "%d-%d", &from, &to)
		ranges[i] = [2]int{from, to}
	}

	merging := true
	for merging {
		merging = false
		for i1, r1 := range ranges {
			for i2, r2 := range ranges {

				if i1 == i2 {
					continue
				}

				if r1[0] <= r2[0] && r1[1] >= r2[0] && r1[1] <= r2[1] {
					// r1 + r2, delete r2
					ranges[i1][1] = r2[1]
				} else if r1[0] <= r2[0] && r1[1] >= r2[1] {
					// r1 wins, delete r2
				} else {
					continue
				}

				ranges = slices.Delete(ranges, i2, i2+1)
				merging = true
				break
			}
			if merging {
				break
			}
		}
	}
	for _, r := range ranges {
		p2 += r[1] - r[0] + 1
	}

	// p1
	for idStr := range strings.SplitSeq(pts[1], "\n") {
		id, _ := strconv.Atoi(string(idStr))
		for _, r := range ranges {
			if r[0] <= id && r[1] >= id {
				p1++
				break
			}
		}
	}

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 05: Cafeteria")
	p1, p2 := solve("../aoc-inputs/2025/05/input.txt")
	// p1, p2 := solve("input-pub.txt")
	fmt.Println("\tPart One:", p1) // 720
	fmt.Println("\tPart Two:", p2) // 357608232770687
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
