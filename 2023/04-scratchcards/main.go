package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strings"
)

func Solve(file string) (int, int) {
	input, _ := os.ReadFile(file)
	cards := strings.Split(strings.TrimSpace(string(input)), "\n")
	memo := make([]int, len(cards))
	for i := 0; i < len(cards); i++ {
		memo[i] = 1
	}
	var re = regexp.MustCompile(`(?m)(\d+)`)

	sum, two := float64(0), 0
	for ic, c := range cards {
		numbers := strings.Split(strings.TrimSpace(strings.Split(c, ":")[1]), "|")
		winning := re.FindAllString(numbers[0], -1)
		checking := re.FindAllString(numbers[1], -1)

		w := 0
		for _, check := range checking {
			if slices.Contains(winning, check) {
				w++
			}
		}

		if w > 0 {
			// next W cards won MEMO[ic] copies each
			for i := 1; i <= w && i+ic < len(memo); i++ {
				memo[i+ic] += memo[ic]
			}
			sum += math.Pow(2, float64(w-1))
		}
		two += memo[ic]
	}

	return int(sum), two
}

func main() {
	one, two := Solve("input1.txt")
	fmt.Println("Day 4: Scratchcards\n\tPart One:", one) // 23441
	fmt.Println("\tPart Two:", two)                      // 5923918
}
