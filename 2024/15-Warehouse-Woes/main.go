package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"slices"
	"strings"
	"time"
)

type Grid map[image.Point]rune

func ReadGrid(in string) (g Grid, bounds image.Rectangle, robot image.Point) {
	g = make(map[image.Point]rune)
	var w, h int
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		w = len(l)
		h = y + 1
		for x, r := range strings.TrimSpace(l) {
			if r == '@' {
				robot = image.Pt(x, y)
				g[image.Point{x, y}] = '.'
				continue
			}
			g[image.Point{x, y}] = r
		}
	}
	bounds = image.Rect(0, 0, w, h)
	return
}

var dir = map[rune]image.Point{'^': {0, -1}, '>': {1, 0}, 'v': {0, 1}, '<': {-1, 0}}

func (g Grid) boxesToMove(points map[image.Point]rune, d rune) map[image.Point]rune {
	foundNew := true
	for foundNew {
		foundNew = false
		newPoints := map[image.Point]rune{}
		for p, r := range points {
			newPoints[p] = r
			np := p.Add(dir[d])

			if _, ok := points[np]; !ok && slices.Contains([]rune{'[', ']', 'O'}, g[np]) {
				newPoints[np] = g[np]
				foundNew = true
				if d == '^' || d == 'v' {
					if g[np] == '[' {
						newPoints[np.Add(dir['>'])] = g[np.Add(dir['>'])]
					}
					if g[np] == ']' {
						newPoints[np.Add(dir['<'])] = g[np.Add(dir['<'])]
					}
				}
			}
		}
		points = newPoints
	}
	return points
}

func (g Grid) moveAllOrNothing(ps map[image.Point]rune, d rune) bool {
	for p := range ps {
		if g[p.Add(dir[d])] == '#' {
			return false
		}
	}
	for p := range ps {
		g[p] = '.'
	}
	for p, r := range ps {
		g[p.Add(dir[d])] = r
	}

	return true
}

func (grid Grid) moveRobot(robot image.Point, moves string) (res int) {
	for _, m := range moves {
		np := robot.Add(dir[m])
		if grid[np] == '.' {
			robot = np
		}
		if slices.Contains([]rune{'[', ']', 'O'}, grid[np]) {
			b2m := grid.boxesToMove(map[image.Point]rune{robot: '.'}, m)
			if grid.moveAllOrNothing(b2m, m) {
				robot = np
			}
		}
	}

	for p := range grid {
		if grid[p] == '[' || grid[p] == 'O' {
			res += 100*p.Y + p.X
		}
	}

	return
}

func solve(file string) (p1, p2 int) {
	in, _ := os.ReadFile(file)
	parts := strings.Split(string(in), "\n\n")
	moves := strings.ReplaceAll(parts[1], "\n", "")
	field := parts[0]
	grid, bounds, robot := ReadGrid(field)

	grid.Show(robot, bounds)
	p1 = grid.moveRobot(robot, moves)

	field = strings.ReplaceAll(field, "#", "##")
	field = strings.ReplaceAll(field, ".", "..")
	field = strings.ReplaceAll(field, "O", "[]")
	field = strings.ReplaceAll(field, "@", "@.")
	grid, bounds, robot = ReadGrid(field)
	grid.Show(robot, bounds)
	p2 = grid.moveRobot(robot, moves)

	grid.renderPNG("final.png", bounds, robot)

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 15: Warehouse Woes")
	// p1, p2 := solve("input.txt")
	p1, p2 := solve("../aoc-inputs/2024/15/input.txt")
	fmt.Println("\tPart One:", p1) // 1448589
	fmt.Println("\tPart Two:", p2) // 1472235
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}

func (g Grid) Show(r image.Point, bounds image.Rectangle) {
	fmt.Println(bounds.Max.X, "x", bounds.Max.Y, ":")
	for y := 0; y <= bounds.Max.Y; y++ {
		for x := 0; x <= bounds.Max.X; x++ {
			if r == image.Pt(x, y) {
				fmt.Print("@")
				continue
			}
			fmt.Print(string(g[image.Pt(x, y)]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g Grid) renderPNG(file string, bounds image.Rectangle, robot image.Point) {
	img := image.NewRGBA(bounds)
	for x := range bounds.Size().X {
		for y := range bounds.Size().Y {
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
		}
	}
	for p, r := range g {
		if p == robot {
			img.Set(p.X, p.Y, color.RGBA{25, 25, 255, 255})
		}
		if r == '[' || r == ']' {
			img.Set(p.X, p.Y, color.RGBA{255, 255, 255, 255})
		}
		if r == '#' {
			img.Set(p.X, p.Y, color.RGBA{100, 100, 100, 255})
		}
	}
	f, _ := os.Create(file)
	png.Encode(f, img)
}
