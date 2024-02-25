package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// solve reads the file and returns the number of times
// the sum of the numbers in the window is greater than the previous sum
func solve(filename string, window int) int {
	f, _ := os.ReadFile(filename)
	lines := strings.Split(string(f), "\n")

	pn := 0
	res := 0
	for i := 0; i < len(lines)-(window-1); i++ {
		n := 0
		for j := i; j < i+window; j++ {
			c, _ := strconv.Atoi(lines[j])
			n += c
		}
		if pn != 0 && n > pn {
			res++
		}
		pn = n
	}

	return res
}

func main() {
	start := time.Now()
	fmt.Println("Day 1: Sonar Sweep")
	fmt.Println("\tPart One:", solve("input.txt", 1)) // 1195
	fmt.Println("\tPart Two:", solve("input.txt", 3)) // 1235
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
