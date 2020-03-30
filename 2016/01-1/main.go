package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/pborzenkov/aoc/pkg/input"
)

type mover struct {
	dir  int
	x, y int
}

func (m *mover) turn(d rune) {
	switch d {
	case 'R':
		m.dir = (m.dir + 1) % 4
	case 'L':
		m.dir = (m.dir + 3) % 4
	default:
		panic("Unknown direction")
	}
}

func (m *mover) move(d int) {
	switch m.dir {
	case 0:
		m.y += d
	case 1:
		m.x += d
	case 2:
		m.y -= d
	case 3:
		m.x -= d
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	data, err := ioutil.ReadAll(input.NewFileOrStdin())
	if err != nil {
		log.Fatalf("failed to read data: %v", err)
	}

	var loc mover
	instruction := strings.Split(string(data), ",")
	for _, instr := range instruction {
		instr = strings.Trim(instr, " \n")
		loc.turn(rune(instr[0]))

		d, err := strconv.Atoi(instr[1:])
		if err != nil {
			log.Fatalf("failed to parse instruction %q: %v", instr, err)
		}
		loc.move(d)
	}

	log.Printf("Final locations is (%d, %d)", loc.x, loc.y)
	log.Printf("Distance from the start is %d", abs(loc.x)+abs(loc.y))
}
