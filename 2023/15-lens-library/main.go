package main

import (
	"fmt"
	"os"
	"strings"
)

func hash(s string) (h int) {
	for _, c := range s {
		h += int(c)
		h *= 17
		h %= 256
	}
	return
}

func hashSeq(seq string) int {
	sum := 0
	for _, s := range strings.Split(strings.TrimSpace(seq), ",") {
		sum += hash(s)
	}
	return sum
}

func p1(f string) int {
	in, _ := os.ReadFile(f)
	return hashSeq(strings.TrimSpace(string(in)))
}

func main() {
	fmt.Println("Day 15: Lens Library")
	fmt.Println("\tPart One:", p1("../aoc-inputs/2023/15/input.txt")) //
	// fmt.Println("\tPart Two:", p2("../aoc-inputs/2023/15/input.txt")) //
}
