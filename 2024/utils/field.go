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
