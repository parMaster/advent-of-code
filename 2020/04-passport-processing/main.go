package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

func solve(f string) (p1, p2 int) {

	in, _ := os.ReadFile(f)
	ps := strings.Split(strings.TrimSpace(string(in)), "\n\n")
	for _, p := range ps {
		p = strings.Replace(p, "\n", " ", -1)
		pfields := strings.Split(p, " ")
		fields := []string{}
		p2score := 0
		for _, pf := range pfields {
			field := pf[0:strings.Index(pf, ":")]
			if !slices.Contains(fields, field) {
				fields = append(fields, field)
			}
			val := pf[strings.Index(pf, ":")+1:]

			if field == "byr" {
				iv, err := strconv.Atoi(val)
				if iv >= 1920 && iv <= 2002 && err == nil {
					p2score++
				}
			}
			if field == "iyr" {
				iv, err := strconv.Atoi(val)
				if iv >= 2010 && iv <= 2020 && err == nil {
					p2score++
				}
			}
			if field == "eyr" {
				iv, err := strconv.Atoi(val)
				if iv >= 2020 && iv <= 2030 && err == nil {
					p2score++
				}
			}

			if field == "hgt" {
				if strings.HasSuffix(val, "cm") {
					iv, err := strconv.Atoi(val[:len(val)-2])
					if iv >= 150 && iv <= 193 && err == nil {
						p2score++
					}
				}
				if strings.HasSuffix(val, "in") {
					iv, err := strconv.Atoi(val[:len(val)-2])
					if iv >= 59 && iv <= 76 && err == nil {
						p2score++
					}
				}
			}
			if field == "hcl" {
				// a # followed by exactly six characters 0-9 or a-f
				var re = regexp.MustCompile(`#[0-9a-f]{6}`)
				if re.Match([]byte(val)) && len(val) == 7 {
					p2score++
				}
			}
			if field == "pid" {
				// a nine-digit number, including leading zeroes
				var re = regexp.MustCompile(`[0-9]{9}`)
				if re.Match([]byte(val)) && len(val) == 9 {
					p2score++
				}
			}
			if field == "ecl" {
				if slices.Index([]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}, val) != -1 {
					p2score++
				}
			}
		}

		if p2score == 7 {
			p2++
		}
		if len(fields) == 8 || (len(fields) == 7 && !slices.Contains(fields, "cid")) {
			p1++
		}

	}

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 04: Passport Processing")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 228
	fmt.Println("\tPart Two:", p2) // 175
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
