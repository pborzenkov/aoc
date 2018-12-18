package main

import (
	"bufio"
	"log"

	"github.com/pborzenkov/aoc/pkg/go/input"
)

func main() {
	r := input.NewFileOrStdin()

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)

	type point struct {
		x, y int
	}
	path := make(map[point]struct{})

	x, y := 0, 0
	path[point{x, y}] = struct{}{}
	for s.Scan() {
		switch s.Text() {
		case ">":
			x += 1
		case "<":
			x -= 1
		case "^":
			y += 1
		case "v":
			y -= 1
		default:
			log.Fatalf("Unsupported input token %q", s.Text())
		}
		path[point{x, y}] = struct{}{}
	}
	if s.Err() != nil {
		log.Fatalf("Scanning completed with error (%v), result might be unreliable", s.Err())
	}
	log.Printf("%d houses will receive at least one present", len(path))
}
