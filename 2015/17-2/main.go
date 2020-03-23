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

	perNumber := make([]int, len(containers)+1)

	var f func(int, int, int) bool
	f = func(i int, used int, liters int) bool {
		if liters < 0 {
			return false
		}
		if liters == 0 {
			perNumber[used]++
		}

		for j, c := range containers[i:] {
			if !f(i+j+1, used+1, liters-c) {
				return true
			}
		}
		return true
	}

	f(0, 0, 150)
	for i, n := range perNumber {
		if n != 0 {
			log.Printf("Can fill 150 using %d containers %d different ways", i, n)
			break
		}
	}
}
