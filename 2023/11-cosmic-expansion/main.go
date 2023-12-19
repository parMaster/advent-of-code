package main

import (
	"fmt"
	"image"
	"os"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
)

type galaxies map[image.Point]int

func readexpand(in string, expansion int) galaxies {
	m := make(galaxies)

	ll := strings.Split(strings.TrimSpace(string(in)), "\n")

	dx := map[int]int{}
	for x := 1; x < len(ll[0]); x++ {
		emptyCol := true
		for y := 0; y < len(ll); y++ {
			if ll[y][x] == '#' {
				emptyCol = false
			}
		}
		if emptyCol {
			dx[x] = dx[x-1] + expansion
			continue
		}
		dx[x] = dx[x-1]
	}

	i := 0
	dy := 0
	for y, l := range ll {
		if !strings.Contains(l, "#") {
			dy += expansion
		}
		for x, r := range strings.TrimSpace(l) {
			if r == '#' {
				m[image.Point{x + dx[x], y + dy}] = i
				i++
			}
		}
	}
	return m
}

func (g galaxies) maxPoint() (mx, my int) {
	for p := range g {
		mx = max(mx, p.X)
		my = max(my, p.Y)
	}
	return mx, my
}

func absPoint(d image.Point) image.Point {
	if d.X < 0 {
		d.X = -d.X
	}

	if d.Y < 0 {
		d.Y = -d.Y
	}
	return d
}

func distance(p1, p2 image.Point) int {
	dist := absPoint(p1.Sub(p2))
	return dist.X + dist.Y
}

func (g galaxies) show() {
	mx, my := g.maxPoint()
	fmt.Println("Loop: mx=", mx, " my=", my, ":")
	if mx+my > 600 {
		fmt.Println("Too big, won't show")
		return
	}
	for y := 0; y <= my; y++ { // lines
		for x := 0; x <= mx; x++ { // cols
			_, ok := g[image.Pt(x, y)]
			if !ok {
				fmt.Print(".")
				continue
			}
			fmt.Print("#")
		}
		fmt.Println()
	}
	fmt.Println()
}

func solve(f string, expansion int) int {
	in, _ := os.ReadFile(f)
	g := readexpand(string(in), expansion)

	if slices.Index(os.Args[1:], "--visual") != -1 {
		g.show()
	}

	points := maps.Keys(g)

	sum := 0
	for i := 0; i < len(points); i++ {
		for j := i; j < len(points); j++ {
			sum += distance(points[i], points[j])
			// fmt.Println(i, j, g[points[i]], g[points[j]], points[i], points[j], distance(points[i], points[j]), minD)
		}
	}

	return sum
}

func main() {
	fmt.Println("Day 11: Cosmic Expansion")
	fmt.Println("\tPart One:", solve("../aoc-inputs/2023/11/input1.txt", 1))         // 10154062
	fmt.Println("\tPart Two:", solve("../aoc-inputs/2023/11/input1.txt", 1000000-1)) // 553083047914
}
