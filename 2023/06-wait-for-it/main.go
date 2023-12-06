package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	file := "../aoc-inputs/2023/06/input0.txt"
	input, _ := os.ReadFile(file)
	l := strings.Split(strings.TrimSpace(string(input)), "\n")

	times := []int{}
	json.Unmarshal([]byte("["+strings.Join(strings.Fields(strings.Split(l[0], ": ")[1]), ",")+"]"), &times)

	distances := []int{}
	json.Unmarshal([]byte("["+strings.Join(strings.Fields(strings.Split(l[1], ": ")[1]), ",")+"]"), &distances)

	total := 1

	for race, t := range times {

		wins := 0
		for hold := 1; hold < t; hold++ {

			if hold*(t-hold) > distances[race] {
				wins++
			}

		}
		total = total * wins
	}

	fmt.Println(total)

}
