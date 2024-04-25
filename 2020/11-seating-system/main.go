package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// . : 0
// L : 1
// # : 2
func read(f string) (grid [][]int) {
	in, _ := os.ReadFile(f)
	lines := strings.Split(strings.TrimSpace(string(in)), "\n")

	for _, line := range lines {
		gridLine := []int{}
		for _, c := range line {
			if c == '.' {
				gridLine = append(gridLine, 0)
			}
			if c == 'L' {
				gridLine = append(gridLine, 1)
			}
		}
		grid = append(grid, gridLine)
	}
	return
}

var around = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func evolve(grid [][]int, rays bool, tolerance int) int {

	for {
		nGrid := [][]int{}
		for i := range grid {
			nGrid = append(nGrid, make([]int, len(grid[i])))
		}

		changed := false
		for i := range grid {
			for j := range grid[i] {

				occupied := 0
				for _, a := range around {

					x, y := i, j

					for {
						x, y = x+a[0], y+a[1]

						// out of bounds
						if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[i]) {
							break
						}

						if grid[x][y] == 1 {
							break
						}

						if grid[x][y] == 2 {
							occupied++
							break
						}
						if !rays {
							break
						}
					}
				}

				// If a seat is empty (L) and there are no occupied seats adjacent to it
				if grid[i][j] == 1 && occupied == 0 {
					// the seat becomes occupied
					nGrid[i][j] = 2
					changed = true
					continue
				}

				// If a seat is occupied (#) and five or more seats adjacent to it are also occupied
				if grid[i][j] == 2 && occupied >= tolerance {
					// the seat becomes empty
					nGrid[i][j] = 1
					changed = true
					continue
				}

				// Otherwise, the seat's state does not change.
				nGrid[i][j] = grid[i][j]

			}
		}

		if !changed {
			// reached final state
			occupied := 0
			for i := range grid {
				for j := range grid[i] {
					if grid[i][j] == 2 {
						occupied++
					}
				}
			}
			return occupied
		} else {
			// evolve further
			for i := range nGrid {
				copy(grid[i], nGrid[i])
			}
		}
	}
}

func solve(f string) (p1, p2 int) {

	p1 = evolve(read(f), false, 4)

	p2 = evolve(read(f), true, 5)

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 11: Seating System")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 2453
	fmt.Println("\tPart Two:", p2) // 2159
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
