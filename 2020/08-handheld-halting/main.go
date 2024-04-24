package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func findLoop(lines []string) (acc int, loopFound bool) {
	pos := 0
	execLines := []int{}
	for {
		if slices.Contains(execLines, pos) {
			return acc, true
		}
		if pos == len(lines) {
			return acc, false
		}
		execLines = append(execLines, pos)

		cmd := strings.Split(lines[pos], " ")
		sign := 1
		if cmd[1][0] == '-' {
			sign = -1
		}
		arg, _ := strconv.Atoi(cmd[1][1:])
		switch cmd[0] {
		case "acc":
			acc = acc + sign*arg
			pos++
		case "jmp":
			pos += sign * arg
		case "nop":
			pos++
		}
	}
}

func solve(f string) (p1, p2 int) {
	in, _ := os.ReadFile(f)
	lines := strings.Split(strings.TrimSpace(string(in)), "\n")

	p1, _ = findLoop(lines)

	for i, line := range lines {

		if line[:3] == "jmp" {
			lines[i] = strings.Replace(line, "jmp", "nop", 1)
			if p2res, p2ok := findLoop(lines); !p2ok {
				p2 = p2res
				break
			}
		}

		if line[:3] == "nop" {
			lines[i] = strings.Replace(line, "nop", "jmp", 1)
			if p2res, p2ok := findLoop(lines); !p2ok {
				p2 = p2res
				break
			}
		}

		lines[i] = line
	}

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 08: Handheld Halting")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 1709
	fmt.Println("\tPart Two:", p2) // 1976
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
