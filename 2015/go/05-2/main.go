package main

import (
	"bufio"
	"log"

	"github.com/pborzenkov/aoc/pkg/go/input"
)

func isNice(str string) bool {
	pairs := make(map[string]int)
	hasRepeatedPair := false
	hasRepeatedLetter := false

	lastWasAdded := false
	lastPair := ""
	for i := 1; i < len(str); i++ {
		pair := str[i-1 : i+1]
		if lastPair != pair || !lastWasAdded {
			pairs[pair]++
			lastWasAdded = true
			if pairs[pair] >= 2 {
				hasRepeatedPair = true
			}
		} else {
			lastWasAdded = false
		}
		lastPair = pair
		if i > 1 {
			if str[i-2] == str[i] {
				hasRepeatedLetter = true
			}
		}
	}
	return hasRepeatedPair && hasRepeatedLetter
}

func main() {
	r := input.NewFileOrStdin()

	s := bufio.NewScanner(r)

	niceStrings := 0
	for s.Scan() {
		if isNice(s.Text()) {
			niceStrings++
		}
	}
	if s.Err() != nil {
		log.Fatalf("Scanning completed with error (%v), result might be unreliable", s.Err())
	}
	log.Printf("Found %d nice strings", niceStrings)
}
