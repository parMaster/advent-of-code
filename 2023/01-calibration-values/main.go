package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/exp/maps"
)

func stringsFile(filename string) ([]string, error) {

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return strings.Fields(string(content)), nil
}

/** Part 1, no regexps */

func isDigit(b byte) bool {
	if (int(b) >= int('0')) && (int(b) <= int('9')) {
		return true
	}
	return false
}

func calLine(s string) (n int) {

	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			n += (int(s[i]) - int('0')) * 10
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if isDigit(s[i]) {
			n += (int(s[i]) - int('0'))
			break
		}
	}

	return
}

func PartOne(file string) int {

	s, _ := stringsFile(file)

	var sum int
	for _, l := range s {
		sum += calLine(l)
	}

	return sum
}

/** Part Two **/

var strNumbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func PartTwo(file string) int {

	s, _ := stringsFile(file)

	first := regexp.MustCompile(`(` + strings.Join(maps.Keys(strNumbers), "|") + `)`)
	last := regexp.MustCompile(`.*(` + strings.Join(maps.Keys(strNumbers), "|") + `)`)

	var sum int
	for _, l := range s {
		firstFound := first.FindStringSubmatch(l)
		lastFound := last.FindStringSubmatch(l)

		v := strNumbers[firstFound[1]]*10 + strNumbers[lastFound[1]]
		sum += v
		// log.Printf("%s ::: %s | %s | +%d = %d ", l, firstFound[1], lastFound[1], v, sum)
	}

	return sum
}

func main() {
	fmt.Println("Day 1: Trebuchet?! \n\tPart One:", PartOne("../aoc-inputs/2023/01/input.txt")) // 53974
	fmt.Println("\tPart Two:", PartTwo("../aoc-inputs/2023/01/input2.txt"))                     // 52840
}
