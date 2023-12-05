package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PartOne(file string) (res int) {

	input, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, l := range lines {
		l = strings.Replace(l, ",", "-", 1)
		p := strings.Split(l, "-")
		t := [4]int{0, 0, 0, 0}
		for pi, pp := range p {
			pp = strings.TrimSpace(pp)
			a, err := strconv.Atoi(pp)
			if err != nil {
				panic(err)
			}
			t[pi] = a
		}
		// left start, left finish, right start, right finish
		ls, lf, rs, rf := t[0], t[1], t[2], t[3]

		if (ls >= rs && lf <= rf) ||
			(rs >= ls && rf <= lf) {
			res++
		}

	}

	return
}

func PartTwo(file string) (res int) {

	input, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, l := range lines {
		l = strings.Replace(l, ",", "-", 1)
		p := strings.Split(l, "-")
		t := [4]int{0, 0, 0, 0}
		for pi, pp := range p {
			pp = strings.TrimSpace(pp)
			a, err := strconv.Atoi(pp)
			if err != nil {
				panic(err)
			}
			t[pi] = a
		}
		// left start, left finish, right start, right finish
		ls, lf, rs, rf := t[0], t[1], t[2], t[3]

		// either left start within the right boundaries of vice versa
		if (ls >= rs && ls <= rf) ||
			(lf >= rs && lf <= rf) ||
			(rs >= ls && rs <= lf) ||
			(rf >= ls && rf <= lf) {
			res++
		}

	}

	return
}

func main() {

	fmt.Println("Day 4: Camp Cleanup\n\tPart One:", PartOne("../aoc-inputs/2022/04/input1.txt"))
	fmt.Println("\tPart Two:", PartTwo("../aoc-inputs/2022/04/input1.txt"))

}
