package main

import (
	"bufio"
	"log"

	"github.com/pborzenkov/aoc/pkg/go/input"
)

var vowels = map[rune]struct{}{
	'a': struct{}{},
	'e': struct{}{},
	'i': struct{}{},
	'o': struct{}{},
	'u': struct{}{},
}
var pairs = map[string]struct{}{
	"ab": struct{}{},
	"cd": struct{}{},
	"pq": struct{}{},
	"xy": struct{}{},
}

func isNice(str string) bool {
	numVowels := 0
	numPairs := 0
	for i, c := range str {
		if _, ok := vowels[c]; ok {
			numVowels++
		}
		if i > 0 {
			if str[i-1] == str[i] {
				numPairs++
			}
			if _, ok := pairs[str[i-1:i+1]]; ok {
				return false
			}
		}
	}
	return numVowels >= 3 && numPairs >= 1
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
