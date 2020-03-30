package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pborzenkov/aoc/pkg/bitmap"
	"github.com/pborzenkov/aoc/pkg/input"
)

const (
	X = 1000
	Y = 1000
)

type mover struct {
	dir     int
	x, y    int
	visited *bitmap.Bitmap
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
	for i := 0; i < d; i++ {
		switch m.dir {
		case 0:
			m.y++
		case 1:
			m.x++
		case 2:
			m.y--
		case 3:
			m.x--
		}
		if m.visited.IsSet(uint((m.y+Y/2)*X + m.x + X/2)) {
			log.Printf("First location visited twice is (%d, %d)", m.x, m.y)
			log.Printf("Distance to it is %d", abs(m.x)+abs(m.y))
			os.Exit(0)
		}
		m.visited.Set(uint((m.y+Y/2)*X+m.x+X/2), 1)
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

	m := mover{
		visited: bitmap.New(X * Y),
	}
	instruction := strings.Split(string(data), ",")
	for _, instr := range instruction {
		instr = strings.Trim(instr, " \n")
		m.turn(rune(instr[0]))

		d, err := strconv.Atoi(instr[1:])
		if err != nil {
			log.Fatalf("failed to parse instruction %q: %v", instr, err)
		}
		m.move(d)
	}
}
