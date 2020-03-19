package main

import (
	"bufio"
	"log"
	"regexp"
	"strconv"

	"github.com/pborzenkov/aoc/pkg/input"
	"github.com/pborzenkov/aoc/pkg/iterutils"
)

var (
	reSitting = regexp.MustCompile(`(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+)\.`)

	happiness = make(map[string]map[string]int)
)

func setHappiness(who string, nextTo string, d int) {
	if _, ok := happiness[who]; !ok {
		happiness[who] = make(map[string]int)
	}
	happiness[who][nextTo] = d
}

func main() {
	r := input.NewFileOrStdin()
	s := bufio.NewScanner(r)

	for s.Scan() {
		groups := reSitting.FindStringSubmatch(s.Text())
		if groups == nil {
			log.Fatalf("failed to parse sitting spec %q", s.Text())
		}

		h, err := strconv.Atoi(groups[3])
		if err != nil {
			log.Fatalf("failed to parse distance %q: %v", groups[3], err)
		}
		if groups[2] == "lose" {
			h = -h
		}

		setHappiness(groups[1], groups[4], h)
	}

	// Adding myself, Go's default zero values are really helpful here
	happiness["me"] = make(map[string]int)

	persons := make([]string, 0, len(happiness))
	for i := range happiness {
		persons = append(persons, i)
	}

	left := func(i int) int {
		if i == 0 {
			return len(persons) - 1
		}
		return i - 1
	}
	right := func(i int) int {
		if i == len(persons)-1 {
			return 0
		}
		return i + 1
	}

	max := 0
	iterutils.ForAllPerms(persons, func(persons []string) {
		var d int

		for i := 0; i < len(persons); i++ {
			d += happiness[persons[i]][persons[left(i)]]
			d += happiness[persons[i]][persons[right(i)]]
		}

		if d > max {
			max = d
		}
	})
	log.Printf("Max happiness is %d", max)
}
