package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

func Solve(file string) (sum int, delete int) {

	input, _ := os.ReadFile(file)

	path := []string{}
	dirs := make(map[string]int)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, l := range lines {

		if len(l) >= 7 && l[:7] == "$ cd .." {
			path = path[:len(path)-1]
		} else if len(l) >= 5 && l[:5] == "$ cd " {
			path = append(path, l[5:])
		} else if l[:4] != "dir " && l != "$ ls" {

			// it's a file
			ll := strings.Split(l, " ")
			size, _ := strconv.Atoi(ll[0])

			for i := range path {

				if _, ok := dirs[strings.Join(path[:i+1], "-")]; !ok {
					dirs[strings.Join(path[:i+1], "-")] = size
				} else {
					dirs[strings.Join(path[:i+1], "-")] += size
				}

			}

		}
	}

	// Part One
	for _, d := range dirs {
		if d < 100000 {
			sum += d
		}
	}

	// Part Two
	diskSize := 70000000
	required := 30000000

	dirSizes := maps.Values(dirs)
	slices.Sort(dirSizes)
	total := dirSizes[len(dirSizes)-1]
	free := diskSize - total

	for _, delete = range dirSizes {
		if free+delete >= required {
			break
		}
	}

	return
}

func main() {

	one, two := Solve("input1.txt")

	fmt.Println("Day 5: Supply Stacks\n\tPart One:", one)
	fmt.Println("\tPart Two:", two)
}
