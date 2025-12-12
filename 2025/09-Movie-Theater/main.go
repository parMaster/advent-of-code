package main

import (
	"advent-of-code/2025/utils"
	"fmt"
	"image"
	"os"
	"strings"
	"time"
)

func solve(file string) (p1, p2 int) {
	g, reds := utils.Grid{}, []image.Point{}
	in, _ := os.ReadFile(file)
	mx, my := 0, 0
	for line := range strings.FieldsSeq(string(in)) {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		g[image.Pt(x, y)] = '#'
		reds = append(reds, image.Pt(x, y))
		mx, my = max(x, mx), max(my, y)
	}
	// g.Render(image.Pt(0, 0), image.Rect(0, 0, mx, my), map[rune]string{'.': "empty", '#': "full"})

	for i := range g {
		for j := range g {
			r := image.Rect(i.X, i.Y, j.X, j.Y)
			p1 = max(p1, (r.Dx()+1)*(r.Dy()+1))
		}
	}

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 09: Movie Theater")
	p1, p2 := solve("../aoc-inputs/2025/09/input.txt") // 1e10 tiles
	// p1, p2 := solve("input-pub.txt")
	fmt.Println("\tPart One:", p1) // 4776100539
	fmt.Println("\tPart Two:", p2) //
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
