package main

import (
	"bufio"
	"fmt"
	"log"
	"sort"

	"github.com/pborzenkov/aoc/pkg/input"
)

func main() {
	r := input.NewFileOrStdin()
	s := bufio.NewScanner(r)

	var containers []int
	for s.Scan() {
		var c int
		fmt.Sscanf(s.Text(), "%d", &c)
		containers = append(containers, c)
	}

	sort.Ints(containers)

	var combinations int

	var f func(int, int) bool
	f = func(i int, liters int) bool {
		if liters < 0 {
			return false
		}
		if liters == 0 {
			combinations++
		}

		for j, c := range containers[i:] {
			if !f(i+j+1, liters-c) {
				return true
			}
		}
		return true
	}

	f(0, 150)
	log.Printf("Total combinations: %d", combinations)
}
