package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
)

type Action byte

const (
	ff    Action = 0
	con   Action = 1
	broad Action = 3
)

var actionMap map[rune]Action = map[rune]Action{
	'&': con,
	'%': ff,
	'b': broad,
}

type State bool

const (
	off State = false
	on  State = true
)

var flipState map[State]State = map[State]State{
	on:  off,
	off: on,
}

type PulseType bool

const (
	low  PulseType = false
	high PulseType = true
)

var statePulse map[State]PulseType = map[State]PulseType{
	on:  high,
	off: low,
}

type Pulse struct {
	from  string
	to    string
	pulse PulseType
}

type Queue []Pulse

type Module struct {
	inputs map[string]PulseType
	action Action
	state  State
	next   []string
}

type Modules map[string]Module

func read(file string) Modules {
	in, _ := os.ReadFile(file)

	modules := Modules{}
	cons := map[string][]string{}
	for _, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		sr := strings.Split(l, "->")
		// sender side
		sender := strings.TrimSpace(sr[0])
		// receiver side
		next := strings.Split(strings.TrimSpace(sr[1]), ", ")
		m := Module{action: actionMap[rune(sender[0])], state: off, next: next, inputs: map[string]PulseType{}}
		modules[sender[1:]] = m
		if actionMap[rune(sender[0])] == con {
			cons[sender[1:]] = []string{}
		}
	}

	for con := range cons {
		for i := range modules {
			if slices.Index(modules[i].next, con) != -1 {
				ins := modules[con].inputs
				ins[i] = low
			}
		}
	}
	return modules
}

func main() {
	fmt.Println("Day 20: Pulse Propagation")

	file := "../aoc-inputs/2023/20/input.txt"
	modules := read(file)

	combinator := ""
	// input data is structured this way: one module is signalling rx, which one?
	for mn := range modules {
		if slices.Index(modules[mn].next, "rx") != -1 {
			combinator = mn
			break
		}
	}
	// log.Println("combinator '", combinator, "' is signalling to ' rx '")

	// and there are some modules that signalling to combinator, which ones?
	cons := []string{}
	for mn := range modules {
		if slices.Index(modules[mn].next, combinator) != -1 {
			cons = append(cons, mn)
		}
	}
	// log.Println(cons, " are signalling to combinator '", combinator, "'")

	// detect cycles in modules that signalling to combinator
	var cycles = map[string]int{}
	var sums = Sums{low: 0, high: 0}
	var startPulse = Pulse{from: "button", to: "roadcaster", pulse: low}
	for n := 1; n < 10000; n++ {
		q := Queue{startPulse}
		sums[startPulse.pulse]++
		// fmt.Println()

		for len(q) > 0 {
			nq := Queue{}

			for i := range q {
				nq = append(nq, modules.call(q[i])...)
			}

			q = nq
			for i := range nq {

				// detecting cycles
				if slices.Index(cons, nq[i].from) != -1 && nq[i].to == "ql" && nq[i].pulse == high {
					cycles[nq[i].from] = n
					// log.Println("cycle detected for", nq[i].from, "length = ", n)
					cons = slices.DeleteFunc(cons, func(con string) bool { return con == nq[i].from })
				}
				if len(cons) == 0 {
					// Least Common Multiple of these cycles must be the button press when all of them are equally high
					fmt.Println("\tPart Two:", lcm(maps.Values(cycles))) // 212986464842911
					os.Exit(0)
				}

				sums[nq[i].pulse]++
			}
		}

		if n == 1000 {
			fmt.Println("\tPart One:", sums[low]*sums[high]) // 787056720
		}
	}
	fmt.Println(sums)
	fmt.Println(sums[low] * sums[high])
}

var pulseLabel map[PulseType]string = map[PulseType]string{low: "->low->", high: "->high->"}

type Sums map[PulseType]int

// call is processing one pulse, thet goes to the receiver module p.to,
// depending on it type, it does its thing and returns pulses to send next to all the receiver.next modules:
// - ff module: receiver state changed + pulses returned
// - con module: inputs changed + pulses returned
// - broad module: next pulses returned
func (m *Modules) call(p Pulse) []Pulse {

	receiver := (*m)[p.to]
	pulses := []Pulse{}

	if receiver.action == broad {
		for _, n := range receiver.next {
			pulses = append(pulses, Pulse{p.to, n, p.pulse})
		}
	}

	if receiver.action == ff && p.pulse == low {
		receiver.state = flipState[receiver.state]
		(*m)[p.to] = receiver

		for _, n := range receiver.next {
			pulses = append(pulses, Pulse{p.to, n, PulseType(receiver.state)})
		}
	}

	if receiver.action == con {
		receiver.inputs[p.from] = p.pulse
		(*m)[p.to] = receiver

		nextPulse := high
		if !slices.Contains(maps.Values(receiver.inputs), low) {
			nextPulse = low
		}

		for _, n := range receiver.next {
			pulses = append(pulses, Pulse{p.to, n, nextPulse})
		}
	}

	return pulses
}

// utils

// Returns prime numbers from 1 to n. Sieve of Eratosthenes
func primes(n int) (result []int) {

	a := []bool{}
	for i := 0; i <= n; i++ {
		a = append(a, true)
	}

	for i := 2; i <= n; i++ {
		if a[i] {
			j := i * i
			for j <= n {
				a[j] = false
				j += i
			}
		}
	}

	for i := 1; i <= n; i++ {
		if a[i] {
			result = append(result, i)
		}
	}
	return
}

// Returns Least Common Multiple of the integers slice
// LCM with division method:
// Divide numbers by prime numbers as long as at least one of the
// numbers is evenly divisible by a prime number.
func lcm(a []int) int {
	res := -1
	primes := primes(slices.Max(a))[1:]
	dividers := []int{}
	for {
		nodiv := true
		for _, p := range primes {
			nodiv := true
			for i := range a {
				if a[i]%p == 0 && a[i] > 1 {
					a[i] /= p
					nodiv = false
				}
			}
			if !nodiv {
				dividers = append(dividers, p)
			}
		}
		if nodiv {
			if slices.Max(a) == 1 {
				for i := range dividers {
					res *= dividers[i]
				}
				return -res
			}
			return res
		}
	}
}
