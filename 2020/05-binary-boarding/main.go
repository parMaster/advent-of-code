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
	maxSeat := 0
	for _, line := range lines {
		row, col := decode(line)
		seat := row*8 + col
		seats = append(seats, seat)
		maxSeat = max(maxSeat, seat)
	}
	slices.Sort(seats)
	for i := 1; i < len(seats); i++ {
		if seats[i] == seats[i-1]+2 {
			p2 = seats[i] - 1
			break
		}
	}

	return maxSeat, p2
}

func main() {
	start := time.Now()
	fmt.Println("Day 05: Binary Boarding")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 978
	fmt.Println("\tPart Two:", p2) // 727
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
