package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func solve(f string) (p1, p2 int) {
	m := read(f)
	p2 = 1

	slopes := [5][2]int{
		{3, 1},
		{1, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for islope, slope := range slopes {
		x, y := 0, 0
		p2trees := 0
		for {
			x += slope[0]
			y += slope[1]

			if y >= len(m) {
				break
			}
			if x >= len(m[0]) {
				x -= len(m[0])
			}

			if m[y][x] == '#' && islope == 0 {
				p1++
			}

			if m[y][x] == '#' {
				p2trees++
			}
		}
		p2 *= p2trees
	}

	return
}

func read(f string) [][]rune {
	in, _ := os.ReadFile(f)
	var m [][]rune
	for _, line := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		m = append(m, []rune(line))
	}
	return m
}

func main() {
	start := time.Now()
	fmt.Println("Day 03: Toboggan Trajectory")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 240
	fmt.Println("\tPart Two:", p2) // 2832009600
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
