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
	file := "../aoc-inputs/2023/20/input02.txt"
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
	for n := 0; n < 4; n++ {

		res := modules.send(startPulse)
		sums[low] += res[low]
		sums[high] += res[high]

	}
	fmt.Println(sums)
	fmt.Println(sums[low] * sums[high])
}

var pulseLabel map[PulseType]string = map[PulseType]string{low: "->low->", high: "->high->"}

type Sums map[PulseType]int

func (m *Modules) send(p Pulse) Sums {
	fmt.Println("Sending: ", p.from, pulseLabel[p.pulse], p.to)
	var sums = Sums{low: 0, high: 0}
	sums[p.pulse]++

	receiver := (*m)[p.to]

	if receiver.action == broad {
		for _, n := range receiver.next {
			res := m.send(Pulse{p.to, n, p.pulse})
			sums[low] += res[low]
			sums[high] += res[high]
		}
	}

	if receiver.action == ff && p.pulse == low {
		receiver.state = flipState[receiver.state]
		(*m)[p.to] = receiver

		for _, n := range receiver.next {
			res := m.send(Pulse{p.to, n, high})
			sums[low] += res[low]
			sums[high] += res[high]
		}
	}

	if receiver.action == con {
		receiver.inputs[p.from] = p.pulse
		(*m)[p.to] = receiver

		for _, sendPulse := range []PulseType{low, high} {
			if !slices.Contains(maps.Values(receiver.inputs), sendPulse) { // all high
				for _, n := range receiver.next {
					res := m.send(Pulse{p.to, n, PulseType(flipState[State(sendPulse)])})
					sums[low] += res[low]
					sums[high] += res[high]
				}
			}
		}

	}

	return sums
}
