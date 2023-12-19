package main

import (
	"fmt"
	"image"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

var moves map[rune]image.Point = map[rune]image.Point{
	'R': {1, 0},
	'L': {-1, 0},
	'U': {0, -1},
	'D': {0, 1},
}

var xyDir = []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

type Field map[image.Point]int

var f Field
var w, h int

func read(in string) (m Field, w int, h int) {

	points := []image.Point{}

	m = make(Field)
	minx, miny := math.MaxInt, math.MaxInt
	pos := image.Point{0, 0}
	for _, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {

		parts := strings.Split(strings.TrimSpace(l), " ")
		if len(parts) == 0 {
			continue
		}

		move := l[0]
		amount, _ := strconv.Atoi(parts[1])
		for i := 0; i < amount; i++ {
			pos = pos.Add(moves[rune(move)])
			points = append(points, pos)

			w = max(w, pos.X)
			minx = min(minx, pos.X)
			h = max(h, pos.Y)
			miny = min(miny, pos.Y)
		}
	}
	w = w + (-minx) + 2
	h = h + (-miny) + 2
	for _, p := range points {
		m[image.Pt(p.X+(-minx)+1, p.Y+(-miny)+1)] = 1
	}

	return
}

func fill(v *Field, start image.Point) {
	_, ok := (*v)[start]
	if !ok {
		(*v)[start] = 2
		for i := range xyDir {
			candidate := start.Add(xyDir[i])
			// (?) in bounds
			if candidate.X > w || candidate.X < 0 || candidate.Y > h || candidate.Y < 0 {
				continue
			}
			fill(v, candidate)
		}
	}
}

func show(m map[image.Point]int, w, h int) {
	asciiBlocks := []string{"░░", "██", "▒▒"}
	fmt.Println(w+1, "x", h+1, ":")
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			fmt.Print(asciiBlocks[m[image.Pt(x, y)]])
		}
		fmt.Println()
	}
	fmt.Println()
}

func p1(file string) int {
	in, _ := os.ReadFile(file)
	f, w, h = read(string(in))
	fill(&f, image.Point{0, 0})

	if slices.Index(os.Args[1:], "--visual") != -1 {
		show(f, w, h)
	}

	lava := 0
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			if v, ok := f[image.Pt(x, y)]; !ok || v == 1 {
				lava++
			}
		}
	}

	return lava
}

// Part two
type Line struct {
	start image.Point
	end   image.Point
}

func readLines(in string) ([]Line, int, int) {
	var lines = []Line{}
	var numberMoves = map[rune]image.Point{'0': moves['R'], '1': moves['D'], '2': moves['L'], '3': moves['U']}
	var p image.Point
	var mx, my, nx, ny int = math.MaxInt, math.MaxInt, 0, 0
	for _, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		parts := strings.Split(strings.TrimSpace(l), " ")
		if len(parts) == 0 {
			continue
		}

		mx = min(mx, p.X)
		my = min(my, p.Y)

		n, _ := strconv.ParseInt(parts[2][2:len(parts[2])-2], 16, 32)

		start := p
		end := p.Add(numberMoves[rune(parts[2][len(parts[2])-2])].Mul(int(n)))
		lines = append(lines, Line{start, end})
		p = end
	}

	for i := range lines {
		lines[i].start = lines[i].start.Add(image.Point{-mx, -my})
		lines[i].end = lines[i].end.Add(image.Point{-mx, -my})
		nx = max(nx, lines[i].start.X)
		ny = max(ny, lines[i].start.Y)
	}
	return lines, nx, ny
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

func p2(file string) int {
	in, _ := os.ReadFile(file)
	lines, _, _ := readLines(string(in))

	// https://en.wikipedia.org/wiki/Shoelace_formula
	// https://en.wikipedia.org/wiki/Pick%27s_theorem

	A := 0
	for _, l := range lines {
		A += (l.start.Y + l.end.Y) * (l.start.X - l.end.X)
	}
	A /= 2

	b := 0
	for _, l := range lines {
		b += distance(l.start, l.end)
	}

	return A + b/2 + 1
}

func main() {
	input := "../aoc-inputs/2023/18/input.txt"

	fmt.Println("Day 18: Lavaduct Lagoon")
	fmt.Println("\tPart One:", p1(input))              // 62573
	fmt.Println("\tPart Two:", p2(input), draw(input)) // 54662804037719
}

func draw(file string) string {
	in, _ := os.ReadFile(file)
	lines, _, _ := readLines(string(in))

	p := plot.New()
	for _, line := range lines {
		pts := make(plotter.XYs, 2)
		pts[0].X = float64(line.start.X)
		pts[0].Y = float64(line.start.Y)
		pts[1].X = float64(line.end.X)
		pts[1].Y = float64(line.end.Y)

		linePlot, _ := plotter.NewLine(pts)

		p.Add(linePlot)
	}
	p.Title.Text = "Lavaduct Lagoon"
	err := p.Save(500, 500, "shape_plot.png")
	if err != nil {
		log.Fatal(err)
	}

	return "\tPlot saved as shape_plot.png"
}
