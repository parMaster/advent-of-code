package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func solve(f string) (p1, p2 int) {

	in, _ := os.ReadFile(f)
	groups := strings.Split(strings.TrimSpace(string(in)), "\n\n")
	for _, group := range groups {
		m := map[rune]int{}
		lines := strings.Split(group, "\n")
		for _, line := range lines {
			for _, c := range line {
				m[c]++
			}
		}
		p1 += len(m)
		for _, ac := range m {
			if ac == len(lines) {
				p2++
			}
		}
	}

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 06: Custom Customs")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 6532
	fmt.Println("\tPart Two:", p2) // 3427
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
