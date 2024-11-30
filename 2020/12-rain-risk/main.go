package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var directions = map[int]string{
	0:   "N",
	180: "S",
	90:  "E",
	270: "W",
}

type Ship struct {
	dir         int // 90:east 180:south 270:west 0:north
	east, north int
}

func (ship *Ship) move(dir string, val int) {
	switch dir {
	case "N":
		ship.north = ship.north + val
	case "E":
		ship.east = ship.east + val
	case "W":
		ship.east = ship.east - val
	case "S":
		ship.north = ship.north - val
	case "F":
		ship.move(directions[ship.dir], val)
	}
}

func (ship *Ship) rotate(degrees int) {
	ship.dir += degrees
	if ship.dir < 0 {
		ship.dir = 360 - (-ship.dir)
	}
	if ship.dir >= 360 {
		ship.dir = ship.dir % 360
	}
}

func solve(f string) (p1, p2 int) {
	in, _ := os.ReadFile(f)
	lines := strings.Split(strings.TrimSpace(string(in)), "\n")

	ship := &Ship{90, 0, 0}

	for _, line := range lines {

		cmd := line[:1]
		val, _ := strconv.Atoi(line[1:])

		fmt.Print(ship.dir, directions[ship.dir], ship)
		fmt.Print(" -> CMD ", cmd, val, " -> ")

		switch cmd {
		case "N", "W", "E", "S", "F":
			ship.move(cmd, val)
		case "R":
			ship.rotate(val)
		case "L":
			ship.rotate(-val)
		}

		fmt.Println(ship.dir, directions[ship.dir], ship)
	}

	if ship.east < 0 {
		ship.east = -ship.east
	}
	if ship.north < 0 {
		ship.north = -ship.north
	}

	p1 = ship.east + ship.north

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 12: Rain Risk")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 757
	fmt.Println("\tPart Two:", p2) //
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
