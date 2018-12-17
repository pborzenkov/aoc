package main

import (
	"bufio"
	"log"

	"github.com/pborzenkov/aoc/pkg/go/input"
)

func main() {
	r := input.NewFileOrStdin()
	defer r.Close()

	s := bufio.NewScanner()
	s.Split(bufio.ScanRunes)

	floor := 0
	position := 0
	for s.Scan() {
		position += 1
		switch s.Text() {
		case "(":
			floor += 1
		case ")":
			floor -= 1
		default:
			log.Fatalf("Unsupported token in the input: %v", s.Text())
		}
		if floor == -1 {
			break
		}
	}
	if s.Err() != nil {
		log.Fatalf("Scanning completed with error (%v), result might be unreliable", s.Err())
	}
	log.Printf("Token position: %v", position)
}
