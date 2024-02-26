package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func readFile(filename string) [][]int {
	grid := [][]int{}
	data, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	for _, line := range lines {
		row := []int{}
		for i := range len(line) {
			row = append(row, int(line[i]-0x30))
		}
		grid = append(grid, row)
	}
	return grid
}

func printGrid(grid [][]int, basins map[[2]int]bool) {
	if !slices.Contains(os.Args, "--visual") {
		return
	}

	for i, row := range grid {
		for j, cell := range row {
			if _, ok := basins[[2]int{i, j}]; ok {
				fmt.Printf("\033[1;37m%d\033[0m", cell)
				continue
			}
			fmt.Printf("\033[1;30m%d\033[0m", cell)
		}
		fmt.Println()
	}
}

var cross = [][]int{
	{-1, 0}, // up
	{1, 0},  // down
	{0, -1}, // left
	{0, 1},  // right
}

func solve(filename string) (int, int) {
	grid := readFile(filename)
	p1 := 0
	p2 := 1
	basins := []int{}
	basinsCells := map[[2]int]bool{}
	for i, row := range grid {
		for j, cell := range row {
			higher := 0
			for _, c := range cross {
				x := i + c[0]
				y := j + c[1]

				if x < 0 || x >= len(grid) || y < 0 || y >= len(row) {
					higher++
					continue
				}

				if grid[x][y] > cell {
					higher++
				}
			}
			if higher == 4 {
				p1 = p1 + cell + 1
				basin, basinCells := basin(grid, i, j)
				for k, v := range basinCells {
					basinsCells[k] = v
				}
				basins = append(basins, basin)
			}
		}
	}

	slices.Sort(basins)
	slices.Reverse(basins)
	for _, b := range basins[:3] {
		p2 *= b
	}
	printGrid(grid, basinsCells)

	return p1, p2
}

func basin(grid [][]int, x, y int) (int, map[[2]int]bool) {
	q := [][2]int{{x, y}}
	visited := map[[2]int]bool{}

	for len(q) > 0 {
		x, y = q[0][0], q[0][1]
		visited[q[0]] = true
		q = q[1:]

		for _, c := range cross {
			dx := x + c[0]
			dy := y + c[1]

			if _, ok := visited[[2]int{dx, dy}]; ok {
				continue
			}

			if dx < 0 || dx >= len(grid) || dy < 0 || dy >= len(grid[0]) {
				continue
			}
			if grid[dx][dy] == 9 {
				continue
			}

			if grid[dx][dy] <= grid[x][y] {
				continue
			}

			q = append(q, [2]int{dx, dy})
		}
	}
	return len(visited), visited
}

func main() {
	start := time.Now()
	fmt.Println("Day 9: Smoke Basin")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 423
	fmt.Println("\tPart Two:", p2) // 1198704
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
