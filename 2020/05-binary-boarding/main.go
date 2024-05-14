package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func decode(seat string) (row, col int) {
	in := [2]int{0, 127}

	for _, c := range seat[:7] {
		if c == 'F' {
			in[1] -= (in[1] - in[0]) / 2
			in[1]--
		} else {
			in[0] += (in[1] - in[0]) / 2
			in[0]++
		}
	}
	row = in[0]

	in = [2]int{0, 7}
	for _, c := range seat[7:] {
		if c == 'L' {
			in[1] -= (in[1] - in[0]) / 2
			in[1]--
		} else {
			in[0] += (in[1] - in[0]) / 2
			in[0]++
		}
	}
	col = in[0]

	return
}

func solve(f string) (p1, p2 int) {
	in, _ := os.ReadFile(f)
	lines := strings.Split(strings.TrimSpace(string(in)), "\n")

	seats := []int{}
	for _, line := range lines {
		row, col := decode(line)
		seat := row*8 + col
		seats = append(seats, seat)
	}

	for i := slices.Min(seats); i <= slices.Max(seats); i++ {
		if !slices.Contains(seats, i) {
			return slices.Max(seats), i
		}
	}

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 05: Binary Boarding")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 978
	fmt.Println("\tPart Two:", p2) // 727
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
