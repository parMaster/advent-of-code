package main

import (
	"fmt"
	"os"
	"strings"
)

var scores string = "0abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func PartOne(file string) (sum int) {

	input, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, l := range lines {
		items := make(map[byte]bool)

		// O(N) solution
		for i := 0; i <= len(l)/2-1; i++ {
			if _, ok := items[l[i]]; !ok {
				items[l[i]] = true
			}
		}
		for i := len(l) / 2; i <= len(l)-1; i++ {
			if _, ok := items[l[i]]; ok {
				sum += strings.Index(scores, string(l[i]))
				break
			}
		}
	}

	return
}

// Repetitive, required refactoring to be scaleable
// but faster than O(N)
func PartTwo(file string) (sum int) {

	input, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for j := 0; j <= (len(lines)-1)/3; j++ {
		li := j * 3

		// hash lines[i]
		items0 := make(map[byte]bool)

		l := lines[li]
		for i := 0; i <= len(l)-1; i++ {
			if _, ok := items0[l[i]]; !ok {
				items0[l[i]] = true
			}
		}

		// hash lines [i+1], only items that present in items0
		items1 := make(map[byte]bool)
		l = lines[li+1]
		for i := 0; i <= len(l)-1; i++ {
			if _, ok := items0[l[i]]; ok {
				items1[l[i]] = true
			}
		}

		// third line, finding any item that in items1 and items0
		l = lines[li+2]
		for i := 0; i <= len(l)-1; i++ {
			if _, ok := items1[l[i]]; ok {
				if _, ok := items0[l[i]]; ok {
					sum += strings.Index(scores, string(l[i]))
					break
				}
			}
		}
	}

	return
}

func main() {
	fmt.Println("Day 3: Rucksack Reorganization\n\tPart One:", PartOne("../aoc-inputs/2022/03/input1.txt"))
	fmt.Println("\tPart Two:", PartTwo("../aoc-inputs/2022/03/input1.txt"))
}
