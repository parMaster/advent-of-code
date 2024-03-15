package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"slices"
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

func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, cell := range row {
			if cell == 0 {
				fmt.Printf("\033[1;37m%d\033[0m", cell)
				continue
			}
			fmt.Printf("\033[0;37m%d\033[0m", cell)
			// fmt.Printf("%d", cell)

		}
		fmt.Println()
	}
	fmt.Println()
}

var neighbors = [][]int{
	{-1, 0},  // up
	{1, 0},   // down
	{0, -1},  // left
	{0, 1},   // right
	{-1, -1}, // up-left
	{-1, 1},  // up-right
	{1, -1},  // down-left
	{1, 1},   // down-right
}

func checkAndFlashNeighbors(grid [][]int) int {
	flashed := 0
	for i, row := range grid {
		for j, cell := range row {
			if cell > 9 {
				flashed++
				for _, n := range neighbors {
					x := i + n[0]
					y := j + n[1]
					if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0 {
						continue
					}
					grid[x][y]++
				}
				grid[i][j] = 0
			}
		}
	}
	return flashed
}

func fadeFlashed(grid [][]int) {
	for i, row := range grid {
		for j, cell := range row {
			if cell > 9 {
				grid[i][j] = 0
			}
		}
	}
}

func incStep(grid [][]int) {
	for i, row := range grid {
		for j := range row {
			grid[i][j]++
		}
	}
}

func solve(filename string) (int, int) {

	grid := readFile(filename)

	if slices.Contains(os.Args, "--visual") {
		fmt.Println("Initial")
		printGrid(grid)
	}

	p1 := 0
	p2 := 0
	totalFlashed := 0
	for step := range 9999 {

		incStep(grid)

		stepFlashed := 0
		for {
			flashed := checkAndFlashNeighbors(grid)
			stepFlashed += flashed
			if flashed == 0 {
				break
			}
		}
		fadeFlashed(grid)

		totalFlashed += stepFlashed

		if step == 99 {
			p1 = totalFlashed
		}

		if slices.Contains(os.Args, "--visual") {
			fmt.Println("Step", step)
			printGrid(grid)
			fmt.Println("Flashed:", totalFlashed)
		}

		if stepFlashed == len(grid)*len(grid[0]) {
			p2 = step + 1
			break
		}
	}
	return p1, p2
}

func main() {
	start := time.Now()
	fmt.Println("Day 11: Dumbo Octopus")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 1625
	fmt.Println("\tPart Two:", p2) // 244
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
