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

// Part Two

type Interval struct {
	from int
	to   int
}

type Option struct {
	xmas map[rune]Interval
	next string
}

type Nodes map[string][]Option

var startInterval []Interval

func p2(file string) int {
	in, _ := os.ReadFile(file)

	nodes := Nodes{}
	for _, l := range strings.Split(strings.TrimSpace(strings.Split(string(in), "\n\n")[0]), "\n") {

		var re = regexp.MustCompile(`(?mi)(?m)([a-z]+)\{(.+)\}`)
		for _, match := range re.FindAllStringSubmatch(l, -1) {

			// last one is default - inverted to every condition before it
			// it also default for every condition because it got the inverted sum of the previous ones
			opts := strings.Split(match[2], ",")
			defaultOpt := Option{xmas: map[rune]Interval{'x': {1, 4000}, 'm': {1, 4000}, 'a': {1, 4000}, 's': {1, 4000}}}
			defaultOpt.next = opts[len(opts)-1]

			opts = opts[:len(opts)-1]
			for _, opt := range opts {

				newOpt := Option{xmas: map[rune]Interval{}}
				for dk, dv := range defaultOpt.xmas {
					newOpt.xmas[dk] = Interval{from: dv.from, to: dv.to}
				}
				newOpt.next = strings.Split(opt, ":")[1]

				rule := strings.Split(opt, ":")[0]
				if strings.Index(rule, ">") != -1 {
					xmas := rune(strings.Split(rule, ">")[0][0])
					from, _ := strconv.Atoi(strings.Split(rule, ">")[1])
					interval := newOpt.xmas[xmas]
					interval.from = from + 1
					newOpt.xmas[xmas] = interval

					interval = defaultOpt.xmas[xmas]
					interval.to = min(interval.to, from)
					defaultOpt.xmas[xmas] = interval
				}

				if strings.Index(rule, "<") != -1 {
					xmas := rune(strings.Split(rule, "<")[0][0])
					to, _ := strconv.Atoi(strings.Split(rule, "<")[1])
					interval := newOpt.xmas[xmas]
					interval.to = to - 1
					newOpt.xmas[xmas] = interval

					interval = defaultOpt.xmas[xmas]
					interval.from = max(interval.from, to)
					defaultOpt.xmas[xmas] = interval
				}
				nodes[match[1]] = append(nodes[match[1]], newOpt)
			}
			nodes[match[1]] = append(nodes[match[1]], defaultOpt)
		}
	}

	startInterval := map[rune]Interval{'x': {1, 4000}, 'm': {1, 4000}, 'a': {1, 4000}, 's': {1, 4000}}

	return walk(&nodes, startInterval, "in", "in")
}

func walk(nodes *Nodes, xmas map[rune]Interval, node string, path string) int {

	if node == "A" {
		// got it
		// fmt.Print("Got To A ", path, " ")
		mul := 1
		for _, ak := range []rune{'a', 'm', 's', 'x'} {
			av := xmas[ak]
			// fmt.Print(string(ak), "[", av.from, ",", av.to, "] ")
			if av.from < av.to {
				mul *= ((av.to - av.from) + 1)
			} else {
				fmt.Println("DANGER! from < to ", xmas)
			}
		}
		// fmt.Println(mul)

		return mul
	}

	if node == "R" {
		return 0
	}

	sum := 0
	for _, option := range (*nodes)[node] {
		isec := intersect(option.xmas, xmas)
		sum += walk(nodes, isec, option.next, path+"->"+option.next)
	}

	return sum
}

func intersect(a, b map[rune]Interval) map[rune]Interval {
	isec := map[rune]Interval{'x': {1, 4000}, 'm': {1, 4000}, 'a': {1, 4000}, 's': {1, 4000}}
	for ak, av := range a {
		bv := b[ak]
		isec[ak] = Interval{max(av.from, bv.from), min(av.to, bv.to)}
	}
	return isec
}

func main() {
	input := "../aoc-inputs/2023/19/input.txt"
	fmt.Println("Day 19: Aplenty")
	fmt.Println("\tPart One:", p1(input)) // 575412
	fmt.Println("\tPart Two:", p2(input)) // 126107942006821
}
