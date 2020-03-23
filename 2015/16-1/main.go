package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"

	"github.com/pborzenkov/aoc/pkg/input"
)

func main() {
	r := input.NewFileOrStdin()
	s := bufio.NewScanner(r)

	need := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	for s.Scan() {
		parts := strings.SplitN(s.Text(), ":", 2)

		suspect := true
		for _, s := range strings.Split(parts[1], ",") {
			p := strings.Split(s, ":")

			what := strings.Trim(p[0], " ")
			much, err := strconv.Atoi(strings.Trim(p[1], " "))
			if err != nil {
				log.Fatalf("failed to parse %s quantity (%s) for %s", what, p[1], parts[0])
			}

			if need[what] != much {
				suspect = false
				break
			}
		}
		if suspect {
			log.Printf("maybe %s", parts[0])
		}
	}
}
