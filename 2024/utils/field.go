package utils

import (
	"fmt"
	"image"
	"strings"
)

// Field is a map of Points (x,y coordinates) to runes
type Field map[image.Point]rune

// ReadField reads a string and returns a Field, width and height
func ReadField(in string) (m Field, w int, h int) {
	m = make(map[image.Point]rune)
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		w = len(l) - 1
		h = y
		for x, r := range strings.TrimSpace(l) {
			m[image.Point{x, y}] = r
		}
	}
	return
}

// Show prints the Field
func (f Field) Show(w, h int) {
	fmt.Println(w+1, "x", h+1, ":")
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			fmt.Print(string(f[image.Pt(x, y)]))
		}
		fmt.Println()
	}
	fmt.Println()
}

// Flatten returns the Field as a string
func (f Field) Flatten(w, h int) (s string) {
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			s += string(f[image.Pt(x, y)])
		}
	}
	return
}

type RField [][]rune

var dir = [][2]int{{-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}}

// recursive search around for term
func (f RField) search(term string, y, x int) (res int) {

	r := rune(term[0])
	if y < 0 || y >= len(f) || x < 0 || x >= len(f[y]) {
		return
	}

	if f[y][x] == r {
		if len(term) == 1 {
			return 1
		}
		for _, d := range dir {
			ny := y + d[0]
			nx := x + d[1]
			res += f.search(term[1:], ny, nx)
		}
	}

	return res
}
