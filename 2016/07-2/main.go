package main

import (
	"bufio"
	"fmt"

	"github.com/pborzenkov/aoc/pkg/input"
)

func main() {
	r := bufio.NewScanner(input.NewFileOrStdin())

	var withSSL int
	for r.Scan() {
		var inHypernet bool
		var roll [3]rune
		var idx int
		abas := make(map[[3]rune]struct{})
		babs := make(map[[3]rune]struct{})

		for _, c := range r.Text() {
			switch c {
			case '[', ']':
				inHypernet = !inHypernet
				roll = [3]rune{0, 0, 0}
			default:
				roll[idx] = c
				if roll[idx] == roll[(idx+1)%3] && roll[idx] != roll[(idx+2)%3] {
					if inHypernet {
						babs[[3]rune{roll[(idx+2)%3], roll[idx], roll[(idx+2)%3]}] = struct{}{}
					} else {
						abas[[3]rune{roll[(idx+1)%3], roll[(idx+2)%3], roll[idx]}] = struct{}{}
					}
				}
				idx = (idx + 1) % 3
			}
		}
		for r := range babs {
			if _, ok := abas[r]; ok {
				withSSL++
				break
			}
		}
	}

	fmt.Printf("Total number of addresses with SSL support is %d\n", withSSL)
}
