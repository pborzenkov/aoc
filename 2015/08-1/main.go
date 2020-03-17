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

	var codeLen, memLen int
	for s.Scan() {
		str := s.Text()

		codeLen += len(str)
		unquoted, err := strconv.Unquote(str)
		if err != nil {
			log.Fatalf("Failed to unquote %q: %v", str, err)
		}
		memLen += len(unquoted)
	}
	if s.Err() != nil {
		log.Fatalf("Scanning completed with error (%v), result might be unreliable", s.Err())
	}

	log.Printf("Difference: %d", codeLen-memLen)
}
