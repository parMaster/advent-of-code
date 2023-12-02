package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Movable Stacks
type MStack Stack

type MStacks struct {
	stacks map[int]*Stack
}

func NewMStacks(n int) *MStacks {
	m := &MStacks{}
	m.stacks = make(map[int]*Stack)

	for i := 0; i < n; i++ {
		m.stacks[len(m.stacks)] = &Stack{}
	}

	return m
}

func (m *MStacks) move(n, from, to int) {
	for i := 0; i < n; i++ {
		v := m.stacks[from].pop()
		m.stacks[to].push(v)
	}
}

func readInput(file string) (stacks *MStacks, movesLines []string) {
	input, _ := os.ReadFile(file)

	parts := strings.Split(string(input), "\n\n")

	crateLines := strings.Split(parts[0], "\n")
	movesLines = strings.Split(parts[1], "\n")

	n := (len(crateLines[0]) + 1) / 4
	stacks = NewMStacks(n)

	for ci := len(crateLines) - 2; ci >= 0; ci-- {
		l := crateLines[ci]
		for i := 0; i < n; i++ {
			item := strings.Trim(string(l[i*4:min(i*4+3, len(l))]), " []")
			if len(item) > 0 {
				stacks.stacks[i].push(item)
			}
		}
	}

	return
}

func PartOne(file string) (res string) {

	stacks, moves := readInput(file)

	for _, l := range moves {
		if l == "" {
			break
		}
		m := strings.Split(l, " ")

		n, _ := strconv.Atoi(m[1])
		from, _ := strconv.Atoi(m[3])
		to, _ := strconv.Atoi(m[5])

		stacks.move(n, from-1, to-1)
	}

	for i := 0; i <= len(stacks.stacks)-1; i++ {
		res = res + stacks.stacks[i].pop()
	}

	return
}

func (m *MStacks) moveTogether(n, from, to int) {
	t := Stack{}
	for i := 0; i < n; i++ {
		v := m.stacks[from].pop()
		t.push(v)
	}
	for i := 0; i < n; i++ {
		m.stacks[to].push(t.pop())
	}
}

func PartTwo(file string) (res string) {

	stacks, moves := readInput(file)

	for _, l := range moves {
		if l == "" {
			break
		}
		m := strings.Split(l, " ")

		n, _ := strconv.Atoi(m[1])
		from, _ := strconv.Atoi(m[3])
		to, _ := strconv.Atoi(m[5])

		stacks.moveTogether(n, from-1, to-1)
	}

	for i := 0; i <= len(stacks.stacks)-1; i++ {
		res = res + stacks.stacks[i].pop()
	}

	return
}

func main() {
	fmt.Println("Day 5: Supply Stacks\n\tPart One:", PartOne("input1.txt"))
	fmt.Println("\tPart Two:", PartTwo("input1.txt"))
}
