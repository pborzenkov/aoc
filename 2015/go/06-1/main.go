package main

import (
	"bufio"
	"log"
	"regexp"
	"strconv"

	"github.com/pborzenkov/aoc/pkg/go/bitmap"
	"github.com/pborzenkov/aoc/pkg/go/input"
)

const (
	side = 1000
)

var (
	baseCommand = `(\d+),(\d+) through (\d+),(\d+)`
	reTurnOn    = regexp.MustCompile(`turn on ` + baseCommand)
	reTurnOff   = regexp.MustCompile(`turn off ` + baseCommand)
	reToggle    = regexp.MustCompile(`toggle ` + baseCommand)
)

type point struct {
	x, y uint
}

func makePoint(xs, ys string) *point {
	x, err := strconv.ParseUint(xs, 10, 32)
	if err != nil {
		log.Fatalf("failed to parse %v: %v", xs, err)
	}
	y, err := strconv.ParseUint(ys, 10, 32)
	if err != nil {
		log.Fatalf("failed to parse %v: %v", ys, err)
	}
	return &point{
		x: uint(x),
		y: uint(y),
	}
}

func applyOp(op func(start, num uint) error, p1, p2 *point) {
	for i := p1.x; i <= p2.x; i++ {
		op(i*side+p1.y, p2.y-p1.y+1)
	}
}

func main() {
	b := bitmap.New(side * side)

	r := input.NewFileOrStdin()
	s := bufio.NewScanner(r)
	for s.Scan() {
		if groups := reTurnOn.FindStringSubmatch(s.Text()); groups != nil {
			applyOp(b.Set, makePoint(groups[1], groups[2]), makePoint(groups[3], groups[4]))
			continue
		}
		if groups := reTurnOff.FindStringSubmatch(s.Text()); groups != nil {
			applyOp(b.Clear, makePoint(groups[1], groups[2]), makePoint(groups[3], groups[4]))
			continue
		}
		if groups := reToggle.FindStringSubmatch(s.Text()); groups != nil {
			applyOp(b.Not, makePoint(groups[1], groups[2]), makePoint(groups[3], groups[4]))
			continue
		}
		log.Fatalf("unknown command: %q", s.Text())
	}
	if s.Err() != nil {
		log.Fatalf("Scanning completed with error (%v), result might be unreliable", s.Err())
	}
	log.Printf("Total number of turned on lights: %d", b.Weight())
}
