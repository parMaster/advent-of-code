package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var re = regexp.MustCompile(`(?m)mul\((\d+),(\d+)\)`)

func mul(s string) (res int) {
	for _, match := range re.FindAllStringSubmatch(string(s), -1) {
		v1, _ := strconv.Atoi(string(match[1]))
		v2, _ := strconv.Atoi(string(match[2]))
		res += v1 * v2
	}
	return res
}

func solve(f string) (p1, p2 int) {
	s, _ := os.ReadFile(f)
	mem := strings.ReplaceAll(string(s), "\n", "")

	p1 = mul(mem)

	donts := strings.Split(mem, "don't()")
	p2 += mul(donts[0])
	for _, dont := range donts[1:] {
		if strings.Contains(dont, "do()") {
			p2 += mul(dont[strings.Index(dont, "do()"):])
		}
	}

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 03: Mull It Over")
	// p1, p2 := solve("input_pub.txt")
	p1, p2 := solve("../aoc-inputs/2024/03/input.txt")
	fmt.Println("\tPart One:", p1) // 187194524
	fmt.Println("\tPart Two:", p2) // 127092535
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
