package main

import (
	"bufio"
	"log"
	"regexp"
	"strconv"

	"github.com/pborzenkov/aoc/pkg/input"
)

var (
	reValue = regexp.MustCompile(`value (\d+) goes to bot (\d+)`)
	reBot   = regexp.MustCompile(`bot (\d+) gives low to (\w+) (\d+) and high to (\w+) (\d+)`)
)

const (
	bot = iota
	output
	maxValueHolder
)

type target struct {
	whoType   int
	whoNumber int
}

type state struct {
	actions map[int][]target
	holders [2]map[int][]int
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to convert %q to int: %v", s, err)
	}
	return i
}

func min(i []int) int {
	if i[0] < i[1] {
		return i[0]
	}
	return i[1]
}

func max(i []int) int {
	if i[0] > i[1] {
		return i[0]
	}
	return i[1]
}

func toType(s string) int {
	switch s {
	case "bot":
		return bot
	case "output":
		return output
	default:
		log.Fatalf("unexpected type %q", s)
		return 0
	}
}

func main() {
	r := bufio.NewScanner(input.NewFileOrStdin())

	s := &state{
		actions: make(map[int][]target),
		holders: [2]map[int][]int{make(map[int][]int), make(map[int][]int)},
	}

	for r.Scan() {
		if groups := reValue.FindStringSubmatch(r.Text()); groups != nil {
			b := toInt(groups[2])
			if len(s.holders[bot][b]) == 2 {
				log.Fatalf("bot %d already has two values", b)
			}
			s.holders[bot][b] = append(s.holders[bot][b], toInt(groups[1]))

		} else if groups := reBot.FindStringSubmatch(r.Text()); groups != nil {
			b := toInt(groups[1])
			if _, ok := s.actions[b]; ok {
				log.Fatalf("bot %d already has an action", b)
			}
			s.actions[b] = []target{
				{toType(groups[2]), toInt(groups[3])},
				{toType(groups[4]), toInt(groups[5])},
			}
		}
	}

	hasSomething := true
	for hasSomething {
		hasSomething = false
		for b, values := range s.holders[bot] {
			if len(values) < 2 {
				continue
			}
			if min(values) == 17 && max(values) == 61 {
				log.Printf("Bot %d handles 17 and 61", b)
			}

			low := s.actions[b][0]
			high := s.actions[b][1]

			lowLen := len(s.holders[low.whoType][low.whoNumber])
			if low.whoType == output && lowLen >= 1 || low.whoType == bot && lowLen >= 2 {
				log.Fatalf("output overflow: type %d, len %d", low.whoType, lowLen)
			}
			s.holders[low.whoType][low.whoNumber] = append(s.holders[low.whoType][low.whoNumber], min(s.holders[bot][b]))

			highLen := len(s.holders[high.whoType][high.whoNumber])
			if high.whoType == output && highLen >= 1 || high.whoType == bot && highLen >= 2 {
				log.Fatalf("output overflow: type %d, len %d", high.whoType, highLen)
			}
			s.holders[high.whoType][high.whoNumber] = append(s.holders[high.whoType][high.whoNumber], max(s.holders[bot][b]))

			s.holders[bot][b] = s.holders[bot][b][:0]

			hasSomething = true
		}
	}

	log.Printf("output value is %d", s.holders[output][0][0]*s.holders[output][1][0]*s.holders[output][2][0])
}
