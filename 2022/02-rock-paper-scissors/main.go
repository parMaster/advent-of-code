package main

import (
	"os"
	"strings"
)

// rules[oponent_move][our_move] gives a score
var rules = map[string]map[string]int{
	"A": {"X": 3, "Y": 6, "Z": 0},
	"B": {"X": 0, "Y": 3, "Z": 6},
	"C": {"X": 6, "Y": 0, "Z": 3},
}

// any move scores anyway
var move = map[string]int{
	"X": 1, "Y": 2, "Z": 3,
}

func PartOne(file string) int {
	input, _ := os.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	sum := 0
	for _, l := range lines {
		strat := strings.Split(strings.TrimSpace(l), " ")

		sum += rules[strat[0]][strat[1]] + move[strat[1]]
	}

	return sum
}

/** Part Two **/

var goals = map[string]int{
	"X": 0, // X means you need to loose
	"Y": 3, // draw
	"Z": 6, // win
}

// rules2[opponent_move][our_goal] gives a Move
var rules2 = map[string]map[string]string{
	"A": {"X": "Z", "Y": "X", "Z": "Y"},
	"B": {"X": "X", "Y": "Y", "Z": "Z"},
	"C": {"X": "Y", "Y": "Z", "Z": "X"},
}

func PartTwo(file string) int {
	input, _ := os.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	sum := 0
	for _, l := range lines {
		strat := strings.Split(strings.TrimSpace(l), " ")

		answer := rules2[strat[0]][strat[1]]

		sum += goals[strat[1]] + move[answer]
	}

	return sum
}
