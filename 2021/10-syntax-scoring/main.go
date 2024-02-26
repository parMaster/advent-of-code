package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

var pairs = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

var antiPairs = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}
var scores = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}
var closingScores = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func solve(filename string) (int, int) {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	p1 := 0
	p2 := []int{}
	for _, line := range lines {
		chars := strings.Split(line, "")
		stack := []string{}
		stop := false
		for _, c := range chars {
			switch c {
			case "(", "{", "<", "[": // open
				stack = append(stack, c)
			case ")", "}", ">", "]": // close
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if pop != pairs[c] {
					// fmt.Println("Illegal char:", c)
					stop = true
					p1 += scores[c]
				}
			}
			if stop {
				break
			}
		}
		if len(stack) > 0 && !stop {
			// fmt.Println("Stack not empty:", stack)
			score := 0
			for i := len(stack) - 1; i >= 0; i-- {
				score *= 5
				score += closingScores[antiPairs[stack[i]]]
				// fmt.Println(c, antiPairs[c], closingScores[antiPairs[stack[i]]], score)
			}
			// fmt.Println(score)
			p2 = append(p2, score)
		}
	}
	slices.Sort(p2)

	return p1, p2[int(len(p2)/2)]
}

func main() {
	start := time.Now()
	fmt.Println("Day 10: Syntax Scoring")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 268845
	fmt.Println("\tPart Two:", p2) // 4038824534
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
