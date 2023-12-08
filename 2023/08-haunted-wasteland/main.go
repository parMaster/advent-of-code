package main

import (
	"fmt"
	"os"
	"strings"
)

var route string
var m map[string][2]string

func main() {
	input, _ := os.ReadFile("../aoc-inputs/2023/08/input1.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	route = lines[0]
	m := map[string][2]string{}
	for _, l := range lines[2:] {
		m[l[0:3]] = [2]string{l[7:10], l[12:15]}
	}

	i := 0
	c := "AAA" // current, start
	dir := map[rune]int{'L': 0, 'R': 1}
	for c != "ZZZ" {
		for _, move := range []rune(route) {
			c = m[c][dir[move]]
			i++
			if c == "ZZZ" {
				break
			}
		}
	}
	fmt.Println("Part One:", i) // 12737
}
