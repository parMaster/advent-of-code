package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

var orders = map[rune]int{'2': 0, '3': 1, '4': 2, '5': 3, '6': 4, '7': 5, '8': 6, '9': 7, 'T': 8, 'J': 9, 'Q': 10, 'K': 11, 'A': 12}

type hand struct {
	cards []rune
	bid   int
	j     bool
}

// strength of a hand: 0 - highest card ... 6 - five of a kind
func (h hand) strength() int {
	g := map[rune]int{}
	for _, c := range h.cards {
		g[c]++
	}

	j := 0
	if h.j {
		if jokers, ok := g['J']; ok {
			j = jokers
			delete(g, 'J')
		}

	}

	s := maps.Values(g)
	slices.Sort(s)
	slices.Reverse(s)

	if h.j && j > 0 {
		if len(s) > 0 {
			s[0] += j
		} else {
			s = append(s, j)
		}
	}

	switch {
	case s[0] == 5:
		return 6
	case s[0] == 4:
		return 5
	case s[0] == 3 && s[1] == 2:
		return 4
	case s[0] == 3 && s[1] == 1:
		return 3
	case s[0] == 2 && s[1] == 2:
		return 2
	case s[0] == 2 && s[1] == 1:
		return 1
	default:
		return 0
	}
}

type hands []hand

func (hh hands) read(f string, j bool) hands {
	cat, _ := os.ReadFile(f)
	lines := strings.Split((strings.TrimSpace(string(cat))), "\n")
	for _, l := range lines {
		cards := []rune(strings.TrimSpace(strings.Fields(l)[0]))
		bid, _ := strconv.Atoi(strings.TrimSpace(strings.Fields(l)[1]))
		hh = append(hh, hand{cards, bid, j})
	}
	return hh
}

func (hh hands) score(joker bool) (sum int) {
	if joker {
		orders['J'] = -1
	}
	sort.Sort(hh)
	for i, h := range hh {
		sum += (i + 1) * h.bid
	}
	return
}

// implementing sort.Interface
func (hh hands) Less(i, j int) bool {
	if hh[i].strength() < hh[j].strength() {
		return true
	} else if hh[i].strength() > hh[j].strength() {
		return false
	}

	for c := 0; c < 5; c++ {
		if orders[hh[i].cards[c]] < orders[hh[j].cards[c]] {
			return true
		} else if orders[hh[i].cards[c]] > orders[hh[j].cards[c]] {
			return false
		}
	}

	return false
}

func (hh hands) Swap(i, j int) {
	hh[i], hh[j] = hh[j], hh[i]
}

func (hh hands) Len() int {
	return len(hh)
}

func main() {
	fmt.Println("Day 7: Camel Cards \n\tPart One:", hands{}.read("../aoc-inputs/2023/07/input1.txt", false).score(false)) // p1: 253638586
	fmt.Println("\tPart Two:", hands{}.read("../aoc-inputs/2023/07/input1.txt", true).score(true))                        // p2: 253253225
}

// dbg
func (hh hands) dump() {
	for i, h := range hh {
		fmt.Println(i+1, string(h.cards), h.bid, h.strength())
	}
}
