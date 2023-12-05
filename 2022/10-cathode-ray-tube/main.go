package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(file string) (int, string) {
	input, _ := os.ReadFile(file)

	m := map[int]int{}
	cycle := 1
	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		l := strings.Split(strings.TrimSpace(line), " ")
		if l[0] == "addx" {
			n, _ := strconv.Atoi(l[1])
			cycle += 2
			m[cycle] = n
		} else {
			cycle++
		}
	}

	x, sum, s := 1, 0, ""
	for i := 1; i < cycle; i++ {
		if n, ok := m[i]; ok {
			x += n
		}
		if i == 20 || (i > 55 && (i-20)%40 == 0) {
			sum += i * x
		}

		if i%40 >= x && i%40 < x+3 {
			s += "#"
		} else {
			s += "."
		}

		if i%40 == 0 {
			s += "\n"
		}
	}

	return sum, s
}

func main() {

	sum, letters := solve("../aoc-inputs/2022/10/input1.txt")
	fmt.Println("Day 10: Cathode ray tube\n\tPart One:", sum) // 14720
	fmt.Println("\tPart Two:")                                // FZBPBFZF
	fmt.Println(letters)
}
