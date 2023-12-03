package main

import (
	"fmt"
	"os"
	"strings"
)

func Solve(file string) (count int, maxScore int) {
	input, _ := os.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	m := map[int]map[int]int{}
	for j, l := range lines {
		l = strings.TrimSpace(l)
		m[j] = make(map[int]int)
		for i := 0; i < len(l); i++ {
			m[j][i] = int(l[i] - 0x30)
		}

	}

	for i := 1; i < len(m)-1; i++ {
		for j := 1; j < len(m[i])-1; j++ {
			visible, score := lookout(m, i, j)
			if visible {
				count++
			}
			maxScore = max(maxScore, score)
		}
	}

	return count + len(m)*2 + (len(m[0])-2)*2, maxScore
}

// look in each side, count the sides with taller trees, count how many trees are visible
// repetitive but readable
func lookout(m map[int]map[int]int, x, y int) (visible bool, score int) {

	score = 1

	sideScore := 0
	sidesClosed := 0
	for i := x - 1; i >= 0; i-- {
		sideScore++
		if m[i][y] >= m[x][y] {
			sidesClosed++
			break
		}
	}
	score *= sideScore

	sideScore = 0
	for i := x + 1; i < len(m); i++ {
		sideScore++
		if m[i][y] >= m[x][y] {
			sidesClosed++
			break
		}
	}
	score *= sideScore

	sideScore = 0
	for j := y - 1; j >= 0; j-- {
		sideScore++
		if m[x][j] >= m[x][y] {
			sidesClosed++
			break
		}
	}
	score *= sideScore

	sideScore = 0
	for j := y + 1; j < len(m[0]); j++ {
		sideScore++
		if m[x][j] >= m[x][y] {
			sidesClosed++
			break
		}
	}
	score *= sideScore

	return sidesClosed < 4, score
}

func printmatrix(m map[int]map[int]int) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Print(m[i][j])
		}
		fmt.Println()
	}
}
