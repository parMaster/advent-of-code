package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func solve(filename string) (int, int) {
	f, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	h, d, d2, aim := 0, 0, 0, 0
	for _, l := range lines {
		cmd, val := l[:strings.Index(l, " ")], l[strings.Index(l, " ")+1:]
		v, _ := strconv.Atoi(val)
		switch cmd {
		case "forward":
			h += v
			d2 += aim * v
		case "down":
			d += v
			aim += v
		case "up":
			d -= v
			aim -= v
		}
	}

	return h * d, h * d2
}

func main() {
	start := time.Now()
	fmt.Println("Day 2: Dive!")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 2272262
	fmt.Println("\tPart Two:", p2) // 2134882034
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
