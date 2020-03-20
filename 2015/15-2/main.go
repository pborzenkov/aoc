package main

import (
	"bufio"
	"log"
	"regexp"
	"strconv"

	"github.com/pborzenkov/aoc/pkg/input"
)

var (
	reRecipe = regexp.MustCompile(`(\w+): capacity (\S+), durability (\S+), flavor (\S+), texture (\S+), calories (\S+)`)
)

const (
	capacity int = iota
	durability
	flavor
	texture
	calories
	maxIngredient
)

type ingredient struct {
	name  string
	props map[int]int
}

func main() {
	r := input.NewFileOrStdin()
	s := bufio.NewScanner(r)

	var ingredients []*ingredient
	for s.Scan() {
		groups := reRecipe.FindStringSubmatch(s.Text())
		if groups == nil {
			log.Fatalf("failed to parse recipe %q", s.Text())
		}

		ing := &ingredient{
			name:  groups[1],
			props: make(map[int]int),
		}
		for i := 0; i < maxIngredient; i++ {
			val, err := strconv.Atoi(groups[i+2])
			if err != nil {
				log.Fatalf("failed to parse ingredient %s %d (%s): %v", groups[1], i, groups[i+2], err)
			}
			ing.props[i] = val
		}
		ingredients = append(ingredients, ing)
	}

	var max int
	var maxIngs []int

	var find func([]int, int, int)
	find = func(set []int, cur int, left int) {
		if cur == len(ingredients) {
			val := 1
			for i := 0; i < calories; i++ {
				t := 0

				for j, ing := range ingredients {
					t += ing.props[i] * set[j]
				}
				if t < 0 {
					t = 0
				}

				val *= t
			}

			var cals int
			for j, ing := range ingredients {
				cals += ing.props[calories] * set[j]
			}

			if cals == 500 && val > max {
				max = val
				maxIngs = make([]int, len(set))
				copy(maxIngs, set)
			}
			return
		}
		if cur == len(ingredients)-1 {
			set[cur] = left
			find(set, cur+1, 0)
		} else {
			for i := 0; i < left; i++ {
				set[cur] = i
				find(set, cur+1, left-i)
			}
		}
	}

	find(make([]int, len(ingredients)), 0, 100)

	log.Printf("Max score is %d, ingredients: %+v", max, maxIngs)
}
