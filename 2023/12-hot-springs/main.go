package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`(?m)(#+)`)

func valid(conditions string, groups []int) bool {

	if (conditions == "" && len(groups) != 0) ||
		(strings.Index(conditions, "#") != -1 && len(groups) == 0) ||
		(strings.Index(conditions, "?") != -1) {
		return false
	}

	foundGroups := re.FindAllString(conditions, -1)
	if len(foundGroups) != len(groups) {
		return false
	}

	for i, cg := range foundGroups {
		if len(cg) != groups[i] {
			return false
		}
	}

	return true
}

func bf(s string, groups []int) (sum int) {

	if strings.Index(s, "?") == -1 {
		if valid(s, groups) {
			return 1
		}
		return 0
	}

	sum += bf(s[:strings.Index(s, "?")]+"."+s[strings.Index(s, "?")+1:], groups)
	sum += bf(s[:strings.Index(s, "?")]+"#"+s[strings.Index(s, "?")+1:], groups)

	return sum
}

func p1(f string, verbose bool) int {
	in, _ := os.ReadFile(f)

	sum := 0
	for i, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		s := strings.Split(strings.TrimSpace(string(l)), " ")[0]
		groups := []int{}
		json.Unmarshal([]byte("["+strings.Split(strings.TrimSpace(string(l)), " ")[1]+"]"), &groups)
		arrangements := bf(s, groups)
		sum += arrangements
		if verbose {
			fmt.Println(i, l, arrangements)
		}
	}
	return sum
}

func main() {
	verbose := false
	fmt.Println("Day 12: Hot Springs")
	fmt.Println("\tPart One:", p1("../aoc-inputs/2023/12/input.txt", verbose)) // 7047
}
