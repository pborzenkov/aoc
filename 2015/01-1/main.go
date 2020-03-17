package main

import (
	"bufio"
	"log"

	"github.com/pborzenkov/aoc/pkg/input"
)

func main() {
	r := input.NewFileOrStdin()
	defer r.Close()

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)

	floor := 0
	for s.Scan() {
		switch s.Text() {
		case "(":
			floor += 1
		case ")":
			floor -= 1
		default:
			log.Fatalf("Unsupported token in the input: %v", s.Text())
		}
	}
	if s.Err() != nil {
		log.Fatalf("Scanning completed with error (%v), result might be unreliable", s.Err())
	}
	log.Printf("Floor: %v", floor)
}
