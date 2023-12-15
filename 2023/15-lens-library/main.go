package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var cache map[string]int = map[string]int{}

// p for projector - array of boxes with lenses 0..255
var p [256][]string = [256][]string{}

func hash(s string) (h int) {
	if v, ok := cache[s]; ok {
		// fmt.Println("hit", len(cache))
		return v
	}
	// fmt.Println("miss")
	for _, c := range s {
		h += int(c)
		h *= 17
		h %= 256
	}
	cache[s] = h
	return
}

func hashSeq(seq string) int {
	sum := 0
	for _, s := range strings.Split(strings.TrimSpace(seq), ",") {
		sum += hash(s)
	}
	return sum
}

func p1(f string) int {
	in, _ := os.ReadFile(f)
	return hashSeq(strings.TrimSpace(string(in)))
}

func p2(f string) int {
	in, _ := os.ReadFile(f)
	for _, step := range strings.Split(strings.TrimSpace(string(in)), ",") {
		// rn=1
		// cm-
		if strings.Index(step, "-") != -1 {
			label := step[:strings.Index(step, "-")]
			box := hash(label)
			p[box] = slices.DeleteFunc(p[box], func(e string) bool {
				return strings.HasPrefix(e, label+" ")
			})
		} else {
			parts := strings.Split(step, "=") // label=focal_length
			label := parts[0]
			focal := parts[1]
			box := hash(label)

			found := false
			for i := range p[box] {
				if strings.HasPrefix(p[box][i], label+" ") {
					p[box][i] = label + " " + focal
					found = true
					break
				}
			}

			if !found {
				p[box] = append(p[box], label+" "+focal)
			}
		}
	}

	sum := 0
	for i, box := range p {
		for slot, lense := range box {
			fl, _ := strconv.Atoi(strings.Split(lense, " ")[1])
			fp := (i + 1) * (slot + 1) * fl
			// fmt.Println(lense, "|", "box", i+1, "slot", slot+1, "focal", fl, "fp=", fp)
			sum += fp
		}
	}

	return sum
}

func main() {
	fmt.Println("Day 15: Lens Library")
	fmt.Println("\tPart One:", p1("../aoc-inputs/2023/15/input.txt")) // 513643
	fmt.Println("\tPart Two:", p2("../aoc-inputs/2023/15/input.txt")) // 265345
}
