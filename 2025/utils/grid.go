package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

// directions   ↖ ↑ ↗ → ↘ ↓ ↙ ←
// directions	0 1 2 3 4 5 6 7
var Dir = []image.Point{{-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}}

// directions	↑ → ↓ ←
// directions	0 1 2 3
var XYDir = []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
var DiagDir = []image.Point{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}}

// directions	↑ → ↓ ←
// directions	^ > v <
var RuneDir = map[rune]image.Point{'^': {0, -1}, '>': {1, 0}, 'v': {0, 1}, '<': {-1, 0}}

var ASCIIBlocks = map[string]string{"full": "██", "half": "▒▒", "empty": "░░"}

type Grid map[image.Point]rune

func ReadGrid(in string) (g Grid, bounds image.Rectangle) {
	g = make(map[image.Point]rune)
	var w, h int
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		w = len(l)
		h = y + 1
		for x, r := range strings.TrimSpace(l) {
			g[image.Point{x, y}] = r
		}
	}
	bounds = image.Rect(0, 0, w, h)
	return
}

func (g Grid) Show(r image.Point, bounds image.Rectangle) {
	fmt.Println(bounds.Max.X, "x", bounds.Max.Y, ":")
	for y := 0; y <= bounds.Max.Y; y++ {
		for x := 0; x <= bounds.Max.X; x++ {
			fmt.Print(string(g[image.Pt(x, y)]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g Grid) Render(r image.Point, bounds image.Rectangle, palette map[rune]string) {
	fmt.Println(bounds.Max.X, "x", bounds.Max.Y, ":")
	for y := 0; y <= bounds.Max.Y; y++ {
		for x := 0; x <= bounds.Max.X; x++ {

			if _, ok := palette[g[image.Pt(x, y)]]; ok {
				fmt.Print(ASCIIBlocks[palette[g[image.Pt(x, y)]]])
			} else {
				fmt.Print(ASCIIBlocks["empty"])
			}

			// fmt.Print(string(g[image.Pt(x, y)]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func ABSPoint(d image.Point) image.Point {
	if d.X < 0 {
		d.X = -d.X
	}

	if d.Y < 0 {
		d.Y = -d.Y
	}
	return d
}

func (g Grid) renderPNG(file string, bounds image.Rectangle) {
	img := image.NewRGBA(bounds)
	for x := range bounds.Size().X {
		for y := range bounds.Size().Y {
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
		}
	}
	for p := range g {
		// img.Set(p.X, p.Y, color.White)
		img.Set(p.X, p.Y, color.RGBA{25, 255, 25, 255})
	}
	f, _ := os.Create(file)
	png.Encode(f, img)
}
