package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func read(filename string, padSize int) (alg []int, grid [][]int, w, h int) {
	in, _ := os.ReadFile(filename)
	parts := strings.Split(strings.TrimSpace(string(in)), "\n\n")
	alg = make([]int, 512)
	for i, v := range string(parts[0]) {
		if v == '#' {
			alg[i] = 1
		} else {
			alg[i] = 0
		}
	}
	gridLines := strings.Split(parts[1], "\n")
	w, h = len(gridLines[0])+padSize*2, len(gridLines)+padSize*2
	pad := make([]int, w)

	grid = [][]int{}
	for range padSize {
		grid = append(grid, pad)
	}
	for _, line := range gridLines {
		row := make([]int, w)
		for i, v := range line {
			if v == '#' {
				row[i+padSize] = 1
			} else {
				row[i+padSize] = 0
			}
		}
		grid = append(grid, row)
	}
	for range padSize {
		grid = append(grid, pad)
	}
	return
}

var grid3x3 = [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {0, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

func enhance(grid [][]int, alg []int, def int) [][]int {
	enhanced := [][]int{}
	for i := 0; i < len(grid); i++ {
		row := make([]int, len(grid[i]))
		enhanced = append(enhanced, row)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			num := 0
			for binplace, v := range grid3x3 {
				y, x := i+v[1], j+v[0]
				val := def
				if y > 0 && y < len(grid) && x > 0 && x < len(grid[y]) {
					val = grid[y][x]
				}
				num += val * int(math.Pow(2, 8-float64(binplace)))
			}
			enhanced[i][j] = alg[num]
		}
	}

	return enhanced
}

func countLit(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			count += grid[i][j]
		}
	}
	return count
}

func solve(filename string) (p1 int, p2 int) {
	alg, grid, _, _ := read(filename, 50)

	for i := range 50 {
		grid = enhance(grid, alg, grid[0][0])
		if i == 1 {
			p1 = countLit(grid)
		}
	}
	p2 = countLit(grid)

	return p1, p2
}

func main() {
	start := time.Now()
	fmt.Println("Day 20: Trench Map")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 5479
	fmt.Println("\tPart Two:", p2) // 19012
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
