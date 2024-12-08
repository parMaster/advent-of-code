package main

import (
	"fmt"
	"image"
	"os"
	"strings"
	"time"
)

type Field map[image.Point]rune

func ReadField(in string) (f Field, ant map[rune][]image.Point, w int, h int) {
	f = make(map[image.Point]rune)
	ant = make(map[rune][]image.Point)
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		w = len(l) - 1
		h = y
		for x, r := range strings.TrimSpace(l) {
			if r != '.' {
				f[image.Point{x, y}] = r
				ant[r] = append(ant[r], image.Point{x, y})
			}
		}
	}
	return
}

func solve(file string) (p1, p2 int) {
	in, _ := os.ReadFile(file)
	f, freqAntenas, w, h := ReadField(string(in))

	bounds := image.Rect(0, 0, w+1, h+1)
	mp1, mp2 := map[image.Point]bool{}, map[image.Point]bool{}
	for _, ants := range freqAntenas {
		for _, ant := range ants {
			for _, pair := range ants {
				dist := ant.Sub(pair)

				for mul := range max(w, h) {
					delta := dist.Mul(mul)
					for _, pt := range []image.Point{
						ant.Add(delta), ant.Sub(delta),
					} {
						if pt.In(bounds) {
							mp2[pt] = true
							if mul == 1 && !pt.Eq(pair) && !pt.Eq(ant) {
								mp1[pt] = true
							}
						}
					}
				}
			}
		}
	}
	f.Show(w, h, mp2)

	return len(mp1), len(mp2)
}

func main() {
	start := time.Now()
	fmt.Println("Day 08: Resonant Collinearity")
	// p1, p2 := solve("input_pub.txt")
	p1, p2 := solve("../aoc-inputs/2024/08/input.txt")
	fmt.Println("\tPart One:", p1) // 261
	fmt.Println("\tPart Two:", p2) // 898
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}

// Show prints the Field
func (f Field) Show(w, h int, mp map[image.Point]bool) {
	fmt.Println(w+1, "x", h+1, ":")
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			if r, ok := f[image.Pt(x, y)]; ok {
				fmt.Print(string(r))
				continue
			}
			if _, ok := mp[image.Pt(x, y)]; ok {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
	fmt.Println()
}
