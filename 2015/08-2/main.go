package main

import (
	"bufio"
	"log"
	"strconv"

	"github.com/pborzenkov/aoc/pkg/input"
)

func main() {
	r := input.NewFileOrStdin()
	s := bufio.NewScanner(r)

	var codeLen, quoteLen int
	for s.Scan() {
		str := s.Text()

		codeLen += len(str)
		quoteLen += len(strconv.Quote(str))
	}
	if s.Err() != nil {
		log.Fatalf("Scanning completed with error (%v), result might be unreliable", s.Err())
	}

	log.Printf("Difference: %d", quoteLen-codeLen)
}
