package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"time"
)

func maxJolt(line []int) (nMax, max int) {
	nMax, max = 0, line[0]
	for n := range line {
		if line[n] > max {
			nMax, max = n, line[n]
		}
	}
	return
}

func maxSet(line []int, setSize int) int {
	lineJolt := 0
	for n := setSize; n > 0; n-- {
		// fmt.Println("line: ", line)
		maxn, maxv := maxJolt(line[:len(line)-(n-1)])
		line = line[maxn+1:]
		lineJolt += maxv * (int(math.Pow10(n - 1)))
		// fmt.Println("maxn, maxv, lineJolt", maxn, maxv, lineJolt)
	}
	return lineJolt

}

func solve(file string) (p1, p2 int) {
	in, _ := os.ReadFile(file)
	for bLine := range bytes.FieldsSeq(in) {
		line := []int{}
		for _, b := range bLine {
			line = append(line, int(b-0x30))
		}
		p1 += maxSet(line, 2)
		p2 += maxSet(line, 12)
	}

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 03: Lobby")
	p1, p2 := solve("../aoc-inputs/2025/03/input.txt")
	// p1, p2 := solve("input-pub.txt")
	fmt.Println("\tPart One:", p1) // 17405
	fmt.Println("\tPart Two:", p2) // 171990312704598
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
