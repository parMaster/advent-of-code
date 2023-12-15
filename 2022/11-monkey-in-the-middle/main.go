package main

import (
	"fmt"
)

func solve(file string) (int, string) {
	// input, _ := os.ReadFile(file)

	// var re = regexp.MustCompile(`(?m)\:((?:[\s,]+\d+)*)`)
	// var re = regexp.MustCompile(`(?m)(\d+)`)
	//
	//  for i, match := range re.FindAllString(str, -1) {
	//     fmt.Println(match, "found at index", i)
	// }

	return 0, ""
}

func main() {

	sum, _ := solve("../aoc-inputs/2022/11/input0.txt")
	fmt.Println("Day 11: Monkey in the Middle\n\tPart One:", sum)
}
