package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"encoding/json"
)

// xVelocity, yVelocity, targetX[from, to], targetY[from, to]
func simulate(xv, yv int, tx, ty []int) (int, bool) {
	x, y := 0, 0
	// fmt.Println("Starting simulation", xv, yv, tx, ty)
	maxY := y

	for x < tx[1] && y > ty[0] {

		x += xv
		if xv > 0 {
			xv--
		} else if xv < 0 {
			xv++
		}
		y += yv
		yv--
		maxY = max(maxY, y)

		// fmt.Println("Step", x, y)

		if x >= tx[0] && x <= tx[1] && y >= ty[0] && y <= ty[1] {
			// fmt.Println("Found a solution", x, y)
			return maxY, true
		}
	}

	return -1, false
}

func solve(filename string) (p1 int, p2 int) {
	txt, _ := os.ReadFile(filename)
	//target area: x=20..30, y=-10..-5
	input := strings.TrimSpace(string(txt))
	tx, ty := []int{}, []int{}
	json.Unmarshal([]byte("["+strings.Replace(strings.Fields(input)[2][2:len(strings.Fields(input)[2])-1], "..", ",", -1)+"]"), &tx)
	json.Unmarshal([]byte("["+strings.Replace(strings.Fields(input)[3][2:], "..", ",", -1)+"]"), &ty)

	maxY := 0
	p2 = 0
	for x := -1000; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			res, found := simulate(x, y, tx, ty)
			if found {
				p2++
			}
			maxY = max(maxY, res)
		}
	}

	return maxY, p2
}

func main() {
	start := time.Now()
	fmt.Println("Day 17: Trick Shot")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 33670
	fmt.Println("\tPart Two:", p2) // 4903
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
