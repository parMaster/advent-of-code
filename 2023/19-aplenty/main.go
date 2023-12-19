package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Part struct {
	props map[rune]int
	curr  string
}

type Step struct {
	prop rune
	sign rune
	val  int
	next string
}

type Flow struct {
	steps []Step
	def   string
}

func p1(file string) int {
	in, _ := os.ReadFile(file)
	workparts := strings.Split(strings.TrimSpace(string(in)), "\n\n")

	var re = regexp.MustCompile(`(?m)([a-z]+)\{(.+)}`)

	flows := map[string]Flow{}
	// workflows
	for _, wf := range strings.Split(strings.TrimSpace(workparts[0]), "\n") {
		namerules := re.FindStringSubmatch(wf)

		steps := []Step{}
		for _, r := range strings.Split(namerules[2], ",") {
			propsign, next, found := strings.Cut(r, ":")
			if !found {
				flows[namerules[1]] = Flow{steps, r}
				break
			}

			prop, val, found := strings.Cut(propsign, "<")
			sign := '<'
			if !found {
				prop, val, found = strings.Cut(propsign, ">")
				sign = '>'
			}
			ival, _ := strconv.Atoi(val)

			steps = append(steps, Step{rune(prop[0]), sign, ival, next})
		}
	}

	// parts
	parts := []Part{}
	for _, p := range strings.Split(strings.TrimSpace(workparts[1]), "\n") {
		tp := Part{props: map[rune]int{}, curr: "in"}
		for _, cond := range strings.Split(strings.Trim(p, "{}"), ",") {
			condParts := strings.Split(cond, "=")
			v, _ := strconv.Atoi(condParts[1])
			tp.props[rune(condParts[0][0])] = v
		}
		parts = append(parts, tp)
	}

	acc := 0
	for len(parts) > 0 {

		for ip, part := range parts {

			if part.curr == "R" || part.curr == "A" {
				if part.curr == "A" {
					acc += part.props['x'] + part.props['m'] + part.props['a'] + part.props['s']
				}
				parts = slices.Delete(parts, ip, ip+1)
				break
			}

			for _, step := range flows[part.curr].steps {

				parts[ip].curr = flows[part.curr].def
				if prop, ok := part.props[step.prop]; ok {
					if step.sign == '<' && prop < step.val {
						parts[ip].curr = step.next
						break
					}
					if step.sign == '>' && prop > step.val {
						parts[ip].curr = step.next
						break
					}
				}
			}
		}

	}

	return acc
}

func main() {
	input := "../aoc-inputs/2023/19/input.txt"

	fmt.Println("Day 19: Aplenty")
	fmt.Println("\tPart One:", p1(input)) //
	// fmt.Println("\tPart Two:", p2(input)) //

}
