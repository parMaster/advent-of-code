package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Bag struct {
	num  int
	name string
}

func search(m map[string][]Bag, start string) bool {
	res := false
	for _, bag := range m[start] {
		if bag.name == "shiny gold" {
			return true
		}
		if search(m, bag.name) {
			res = true
		}
	}
	return res
}

func count(m map[string][]Bag, parentBag string, numParents int) int {

	res := numParents

	for _, insideBag := range m[parentBag] {
		res = res + numParents*count(m, insideBag.name, insideBag.num)

	}

	return res
}

func solve(f string) (p1, p2 int) {
	in, _ := os.ReadFile(f)
	lines := strings.Split(strings.TrimSpace(string(in)), "\n")

	bags := map[string][]Bag{}
	for _, line := range lines {
		line = strings.ReplaceAll(line, "bags", "")
		line = strings.ReplaceAll(line, "bag", "")
		parts := strings.Split(line, "contain")
		bag := strings.TrimSpace(parts[0])
		content := strings.Split(parts[1][:len(parts[1])-1], ",")

		bags[bag] = []Bag{}
		for _, cont := range content {
			cont = strings.TrimSpace(cont)
			num, _ := strconv.Atoi(cont[:strings.Index(cont, " ")])
			name := cont[strings.Index(cont, " ")+1:]
			bags[bag] = append(bags[bag], Bag{num, name})
		}

	}
	for from := range bags {
		if search(bags, from) {
			p1++
		}
	}

	p2 = count(bags, "shiny gold", 1)
	p2-- // don't count shiny gold bag itself

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 07: Handy Haversacks")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 316
	fmt.Println("\tPart Two:", p2) // 11310
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
