package main

import (
	"bufio"
	"log"
	"regexp"
	"strconv"

	"github.com/pborzenkov/aoc/pkg/input"
)

var (
	reDescription = regexp.MustCompile(`\w+ can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds\.`)
)

func main() {
	r := input.NewFileOrStdin()
	s := bufio.NewScanner(r)

	total := 2503
	var max int
	for s.Scan() {
		groups := reDescription.FindStringSubmatch(s.Text())
		if groups == nil {
			log.Fatalf("failed to parse deer description %q", s.Text())
		}

		speed, err := strconv.Atoi(groups[1])
		if err != nil {
			log.Fatalf("failed to parse speed %q: %v", groups[1], err)
		}
		flies, err := strconv.Atoi(groups[2])
		if err != nil {
			log.Fatalf("failed to parse fly time %q: %v", groups[2], err)
		}
		rests, err := strconv.Atoi(groups[3])
		if err != nil {
			log.Fatalf("failed to parse rest time %q: %v", groups[3], err)
		}

		this := speed * (total / (flies + rests)) * flies
		rem := total % (flies + rests)
		if rem > flies {
			rem = flies
		}
		this += rem * speed
		if this > max {
			max = this
		}
	}

	log.Printf("Max distance is %d", max)
}
