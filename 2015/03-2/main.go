package main

import (
	"bufio"
	"log"

	"github.com/pborzenkov/aoc/pkg/input"
)

func main() {
	r := input.NewFileOrStdin()

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)

	type point struct {
		x, y int
	}

	var curPos [2]point
	path := make(map[point]struct{})
	path[point{}] = struct{}{}

	token := 0
	for s.Scan() {
		p := &curPos[token%2]

		switch s.Text() {
		case ">":
			p.x += 1
		case "<":
			p.x -= 1
		case "^":
			p.y += 1
		case "v":
			p.y -= 1
		default:
			log.Fatalf("Unsupported input token %q", s.Text())
		}
		path[*p] = struct{}{}
		token += 1
	}
	if s.Err() != nil {
		log.Fatalf("Scanning completed with error (%v), result might be unreliable", s.Err())
	}
	log.Printf("%d houses will receive at least one present", len(path))
}
