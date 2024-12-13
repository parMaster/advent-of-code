package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

// Cramer's Rule
func solveEq(a1, a2, b1, b2, c1, c2 float64) (x, y float64) {
	x = (c1*b2 - c2*b1) / (a1*b2 - a2*b1)
	y = (a1*c2 - a2*c1) / (a1*b2 - a2*b1)
	return
}

func solve(file string) (p1, p2 int) {
	input, _ := os.ReadFile(file)
	for _, ins := range strings.Split(string(input), "\n\n") {
		in := strings.Split(ins, "\n")
		var a1, a2, b1, b2, c1, c2 float64
		fmt.Sscanf(in[0], "Button A: X+%f, Y+%f", &a1, &a2)
		fmt.Sscanf(in[1], "Button B: X+%f, Y+%f", &b1, &b2)
		fmt.Sscanf(in[2], "Prize: X=%f, Y=%f", &c1, &c2)

		x, y := solveEq(a1, a2, b1, b2, c1, c2)

		if math.Abs(x-float64(int(x))) < 1e-6 && math.Abs(y-float64(int(y))) < 1e-6 {
			p1 += int(x*3 + y)
		}

		x, y = solveEq(a1, a2, b1, b2, c1+10000000000000, c2+10000000000000)

		if math.Abs(x-float64(int(x))) < 1e-6 && math.Abs(y-float64(int(y))) < 1e-6 {
			p2 += int(x*3 + y)
		}
	}

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 13: Claw Contraption")
	// p1, p2 := solve("input.txt")
	p1, p2 := solve("../aoc-inputs/2024/13/input.txt")
	fmt.Println("\tPart One:", p1) // 29711
	fmt.Println("\tPart Two:", p2) // 94955433618919
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
