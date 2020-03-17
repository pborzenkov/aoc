package main

import (
	"bufio"
	"log"
	"regexp"
	"strconv"

	"github.com/pborzenkov/aoc/pkg/input"
)

var (
	reDistance = regexp.MustCompile(`(\w+) to (\w+) = (\d+)`)

	distances = make(map[string]map[string]int)
)

func setDistance(from string, to string, d int) {
	if _, ok := distances[from]; !ok {
		distances[from] = make(map[string]int)
	}
	if _, ok := distances[to]; !ok {
		distances[to] = make(map[string]int)
	}
	distances[from][to] = d
	distances[to][from] = d
}

func forAllPerms(cities []string, cb func([]string)) {
	s := make([]string, len(cities))
	copy(s, cities)

	var p func(int)
	p = func(i int) {
		if i > len(s) {
			cb(s)
			return
		}

		p(i + 1)
		for j := i + 1; j < len(s); j++ {
			s[i], s[j] = s[j], s[i]
			p(i + 1)
			s[i], s[j] = s[j], s[i]
		}
	}
	p(0)
}

func main() {
	r := input.NewFileOrStdin()
	s := bufio.NewScanner(r)

	for s.Scan() {
		groups := reDistance.FindStringSubmatch(s.Text())
		if groups == nil {
			log.Fatalf("failed to parse distance spec %q", s.Text())
		}

		distance, err := strconv.Atoi(groups[3])
		if err != nil {
			log.Fatalf("failed to parse distance %q: %v", groups[3], err)
		}
		setDistance(groups[1], groups[2], distance)
	}

	cities := make([]string, 0, len(distances))
	for i := range distances {
		cities = append(cities, i)
	}

	max := 0
	forAllPerms(cities, func(cities []string) {
		var d int
		for i := 1; i < len(cities); i++ {
			d += distances[cities[i-1]][cities[i]]
		}
		if d > max {
			max = d
		}
	})
	log.Printf("Longest route is %d", max)
}
