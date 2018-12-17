package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	var r io.Reader = os.Stdin
	var err error
	if len(os.Args) >= 2 {
		r, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("Cannot open input file: %v", err)
		}
	}
	s := bufio.NewScanner(r)
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
