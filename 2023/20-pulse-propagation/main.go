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

func main() {
	file := "../aoc-inputs/2023/20/input.txt"
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

	var sums = Sums{low: 0, high: 0}
	var startPulse = Pulse{from: "button", to: "roadcaster", pulse: low}
	fmt.Println(startPulse.from, startPulse.pulse, startPulse.to)
	for n := 0; n < 1000; n++ {
		q := Queue{startPulse}
		sums[startPulse.pulse]++
		fmt.Println()

		for len(q) > 0 {
			nq := Queue{}

			for i := range q {
				fmt.Println(q[i].from, q[i].pulse, q[i].to)
				nq = append(nq, modules.call(q[i])...)
			}

			q = nq
			for i, _ := range nq {
				sums[nq[i].pulse]++
				// fmt.Println(v.from, v.pulse, v.to)
			}
		}
	}
	fmt.Println(sums)
	fmt.Println(sums[low] * sums[high])
}

var pulseLabel map[PulseType]string = map[PulseType]string{low: "->low->", high: "->high->"}

type Sums map[PulseType]int

// pulse is received by p.to, p.to believes that state and inputs was changed before
// so only pulses sent for every this.next module:
// - ff module: next state changed + pulse returned
// - con module: inputs changed + pulse returned
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
