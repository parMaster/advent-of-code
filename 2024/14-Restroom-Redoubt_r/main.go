package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
	"time"
)

type Vector struct {
	p image.Point
	v image.Point
}

func solve(file string) (p1, p2 int) {
	in, _ := os.ReadFile(file)
	lines := strings.Split(string(in), "\n")
	vecs := []Vector{}
	res := map[image.Point]int{}
	w, h, sec := 101, 103, 100
	var p, v image.Point
	for _, line := range lines {
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &p.X, &p.Y, &v.X, &v.Y)
		vecs = append(vecs, Vector{p, v})
	}

	for i := 0; p2 == 0; i++ {
		for vi, v := range vecs {
			v.p = v.p.Add(v.v)

			// wrap around
			if v.p.X < 0 {
				v.p.X += w
			}
			if v.p.Y < 0 {
				v.p.Y += h
			}

			v.p = v.p.Mod(image.Rect(0, 0, w, h))
			vecs[vi] = v
		}

		grid := map[image.Point]int{}
		for _, v := range vecs {
			grid[v.p] = 1
		}

		if i == sec-1 {
			for _, v := range vecs {
				res[v.p]++
			}
		}

		// p2
		middle := 0 // 1/9 of the grid
		for y := h / 3; y <= h/3*2; y++ {
			for x := w / 3; x <= w/3*2; x++ {
				if _, ok := grid[image.Pt(x, y)]; ok {
					middle++
				}
			}
		}
		// more than half of the points are in the middle (>1/2 of pixels in 1/9 of the grid)
		if middle > len(grid)/2 {
			p2 = i + 1
			show(w, h, grid)
			renderPNG("day14.png", w, h, grid)
			break
		}
	}

	p1 = 1
	for _, r := range []image.Rectangle{
		image.Rect(0, 0, w/2, h/2),
		image.Rect(w/2+1, 0, w+1, h/2),
		image.Rect(0, h/2+1, w/2, h),
		image.Rect(w/2+1, h/2+1, w, h),
	} {
		sum := 0
		for p, n := range res {
			if p.In(r) {
				sum += n
			}
		}
		p1 *= sum
	}

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 14: Restroom Redoubt")
	// p1, p2 := solve("input.txt")
	p1, p2 := solve("../aoc-inputs/2024/14/input.txt")
	fmt.Println("\tPart One:", p1) // 236628054
	fmt.Println("\tPart Two:", p2) // 7584
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}

func show(w, h int, grid map[image.Point]int) {
	fmt.Println(w, "x", h, ":")
	var ASCIIBlocks = map[string]string{"full": "██", "half": "▒▒", "empty": "░░"}
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			if _, ok := grid[image.Pt(x, y)]; ok {
				fmt.Print(ASCIIBlocks["full"])
			} else {
				fmt.Print(ASCIIBlocks["empty"])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func renderPNG(file string, w, h int, grid map[image.Point]int) {
	rect := image.Rect(0, 0, w, h)
	img := image.NewRGBA(rect)
	for p := range grid {
		img.Set(p.X, p.Y, color.RGBA{25, 255, 25, 255})
	}
	f, _ := os.Create(file)
	png.Encode(f, img)
}
