package main

import (
	"fmt"
	"time"

	"github.com/parMaster/advent-of-code/2024/utils"
)

func solve(f string) (p1, p2 int) {
	lines := utils.ReadLines(f)

	return len(lines), 0
}

func main() {
	start := time.Now()
	fmt.Println("Day 12: Rain Risk")
	p1, p2 := solve("input_pub.txt")
	// p1, p2 := solve("../aoc-inputs/2024/01/input.txt")
	fmt.Println("\tPart One:", p1) //
	fmt.Println("\tPart Two:", p2) //
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
