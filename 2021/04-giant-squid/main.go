package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Board struct {
	board [5][5]int
	cols  [5]int
	rows  [5]int
}

func readInput(filename string) ([]Board, []int) {
	b := []Board{}
	draws := []int{}

	content, _ := os.ReadFile(filename)
	groups := strings.Split(strings.TrimSpace(string(content)), "\n\n")
	drawStrings := strings.Split(strings.TrimSpace(groups[0]), ",")
	for _, d := range drawStrings {
		draw, _ := strconv.Atoi(d)
		draws = append(draws, draw)
	}

	for _, g := range groups[1:] {
		lines := strings.Split(strings.TrimSpace(g), "\n")
		board := [5][5]int{}
		for i, l := range lines {
			numbers := strings.Split(strings.Replace(strings.TrimSpace(l), "  ", " ", -1), " ")
			for j, n := range numbers {
				board[i][j], _ = strconv.Atoi(n)
			}
		}
		b = append(b, Board{board: board, cols: [5]int{0, 0, 0, 0, 0}, rows: [5]int{0, 0, 0, 0, 0}})
	}

	return b, draws
}

func solve(filename string) (int, int) {
	boards, draws := readInput(filename)
	boardsWon := []int{}
	firstToWin, lastToWin := 0, 0

	for id, d := range draws {
		for ib, b := range boards {

			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if b.board[i][j] == d {
						b.rows[i]++
						b.cols[j]++
					}
				}
			}

			for i := 0; i < 5; i++ {
				if b.rows[i] == 5 || b.cols[i] == 5 {
					// fmt.Println("Bingo!", ib, d)
					if !slices.Contains(boardsWon, ib) {
						boardsWon = append(boardsWon, ib)
					}

					sum := 0
					for i := 0; i < 5; i++ {
						for j := 0; j < 5; j++ {
							if !slices.Contains(draws[:id+1], b.board[i][j]) {
								sum += b.board[i][j]
							}
						}
					}
					if firstToWin == 0 {
						firstToWin = sum * d
					}
					if len(boardsWon) == len(boards) {
						lastToWin = sum * d
						return firstToWin, lastToWin
					}
				}
			}
			boards[ib] = b
		}
	}

	return 0, 0
}

func main() {
	start := time.Now()
	fmt.Println("Day 4: Giant Squid")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 69579
	fmt.Println("\tPart Two:", p2) // 14877
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
