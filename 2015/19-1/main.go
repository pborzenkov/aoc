package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"

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

	var molecule string
	var transes []trans
	for s.Scan() {
		if readMolecule {
			molecule = s.Text()
			continue
		}
		if s.Text() == "" {
			readMolecule = true
			continue
		}

		var from, to string
		fmt.Sscanf(s.Text(), "%s => %s", &from, &to)
		transes = append(transes, trans{from, to})
	}

	distinct := 0
	molecules := make(map[string]struct{})
	for _, t := range transes {
		idx := 0
		for idx < len(molecule) {
			i := strings.Index(molecule[idx:], t.from)
			if i < 0 {
				break
			}
			i += idx
			str := molecule[:i] + t.to + molecule[i+len(t.from):]
			if _, ok := molecules[str]; !ok {
				distinct++
				molecules[str] = struct{}{}
			}
			idx = i + len(t.from)
		}
	}

	log.Printf("total distinct molecules: %d", distinct)
}
