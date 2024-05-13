package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func solve(f string) (p1, p2 int) {
	input := read(f)
	for _, in := range input {
		cnt := strings.Count(in.password, in.char)
		if cnt >= in.in[0] && cnt <= in.in[1] {
			p1++
		}
	}

	for _, in := range input {
		if (in.password[in.in[0]-1] == in.char[0] && in.password[in.in[1]-1] != in.char[0]) ||
			(in.password[in.in[0]-1] != in.char[0] && in.password[in.in[1]-1] == in.char[0]) {
			p2++
		}
	}

	return p1, p2
}

type policy struct {
	password string
	char     string
	in       [2]int
}

func read(filename string) []policy {
	fc, _ := os.ReadFile(filename)

	lines := strings.Split(strings.TrimSpace(string(fc)), "\n")
	pa := []policy{}
	for _, line := range lines {
		pt := strings.Split(line, " ")
		p := policy{}
		json.Unmarshal([]byte("["+strings.Replace(pt[0], "-", ",", 1)+"]"), &p.in)
		p.char = string(pt[1][0])
		p.password = pt[2]

		pa = append(pa, p)
	}
	return pa
}

func main() {
	start := time.Now()
	fmt.Println("Day 02: Password Philosophy")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 398
	fmt.Println("\tPart Two:", p2) // 562
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
