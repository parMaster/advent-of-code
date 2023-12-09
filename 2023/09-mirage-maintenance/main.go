package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
)

func predict(a []int) int {
	if m := slices.Max(a); m == 0 && slices.Max(a) == slices.Min(a) {
		return 0
	}

	nexta := []int{}
	for i := 0; i < len(a)-1; i++ {
		nexta = append(nexta, a[i+1]-a[i])
	}

	return a[len(a)-1] + predict(nexta)
}

func solve(f string, reverse bool) (sum int) {
	in, _ := os.ReadFile(f)

	for _, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		a := []int{}
		json.Unmarshal([]byte("["+strings.Join(strings.Fields(l), ",")+"]"), &a)
		if reverse {
			slices.Reverse(a)
		}
		sum += predict(a)
	}
	return
}

func main() {
	fmt.Println("Day 9: Mirage Maintenance")
	fmt.Println("\tPart One:", solve("../aoc-inputs/2023/09/input.txt", false)) // 1921197370
	fmt.Println("\tPart Two:", solve("../aoc-inputs/2023/09/input.txt", true))  // 1124
}
