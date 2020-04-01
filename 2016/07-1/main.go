package main

import (
	"bufio"
	"fmt"

	"github.com/pborzenkov/aoc/pkg/input"
)

func main() {
	r := bufio.NewScanner(input.NewFileOrStdin())

	var withTLS int
	for r.Scan() {
		var hasTLS bool
		var inABBA bool
		var roll [4]rune
		var idx int

	loop:
		for _, c := range r.Text() {
			switch c {
			case '[', ']':
				inABBA = !inABBA
				roll = [4]rune{0, 0, 0, 0}
			default:
				roll[idx] = c
				if roll[idx] == roll[(idx+1)%4] && roll[(idx+3)%4] == roll[(idx+2)%4] && roll[idx] != roll[(idx+3)%4] {
					if inABBA {
						hasTLS = false
						break loop
					}
					hasTLS = true
				}
				idx = (idx + 1) % 4
			}
		}
		if hasTLS {
			withTLS++
		}
	}

	fmt.Printf("Total number of addresses with TLS support is %d\n", withTLS)
}
