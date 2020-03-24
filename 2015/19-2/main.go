package main

import (
	"bufio"
	"log"
	"strings"
	"unicode"

	"github.com/pborzenkov/aoc/pkg/input"
)

type trans struct {
	from string
	to   string
}

func main() {
	r := input.NewFileOrStdin()
	s := bufio.NewScanner(r)

	readMolecule := false
	var target string
	for s.Scan() {
		if readMolecule {
			target = s.Text()
			continue
		}
		if s.Text() == "" {
			readMolecule = true
			continue
		}
	}

	Rn := strings.Count(target, "Rn")
	Y := strings.Count(target, "Y")
	elements := 0
	strings.Map(func(r rune) rune {
		if unicode.IsUpper(r) {
			elements++
		}
		return r
	}, target)

	log.Printf("Steps: %d", elements-Rn*2-2*Y-1)
}
