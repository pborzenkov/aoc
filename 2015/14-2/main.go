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

type deer struct {
	speed int
	flies int
	rests int

	distance int
	score    int
}

func main() {
	r := input.NewFileOrStdin()
	s := bufio.NewScanner(r)

	total := 2503
	var deers []*deer
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

		deers = append(deers, &deer{
			speed: speed,
			flies: flies,
			rests: rests,
		})
	}

	var maxScore int
	for i := 1; i <= total; i++ {
		var maxDistance int

		for _, d := range deers {
			t := i % (d.flies + d.rests)
			if t > 0 && t <= d.flies {
				d.distance += d.speed
			}
			if d.distance > maxDistance {
				maxDistance = d.distance
			}
		}
		for _, d := range deers {
			if d.distance == maxDistance {
				d.score++
			}
			if d.score > maxScore {
				maxScore = d.score
			}
		}
	}

	log.Printf("Max score is %d", maxScore)
}
