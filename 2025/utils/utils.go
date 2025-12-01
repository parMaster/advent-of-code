package utils

import (
	"encoding/json"
	"os"
	"strings"
)

func ReadLinesOfNumbers(f string) [][]int {
	res := [][]int{}
	s, _ := os.ReadFile(f)
	lines := strings.Split(string(s), "\n")
	for _, line := range lines {
		var levels []int
		json.Unmarshal([]byte(line), &levels)
		res = append(res, levels)
	}
	return res
}

func ReadLines(file string) []string {
	input, _ := os.ReadFile(file)
	lines := strings.Fields(strings.TrimSpace(string(input)))

	ll := []string{}
	for _, l := range lines {
		ll = append(ll, l)
	}
	return ll
}
