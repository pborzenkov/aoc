package main

import (
	"bufio"
	"fmt"

	"github.com/pborzenkov/aoc/pkg/input"
)

func main() {
	keypad := []string{
		"", "", "1", "", "",
		"", "2", "3", "4", "",
		"5", "6", "7", "8", "9",
		"", "A", "B", "C", "",
		"", "", "D", "", "",
	}

	x := 0
	y := 2

	r := bufio.NewScanner(input.NewFileOrStdin())
	for r.Scan() {
		for _, c := range r.Text() {
			switch c {
			case 'R':
				if x < 4 && keypad[y*5+x+1] != "" {
					x++
				}
			case 'L':
				if x > 0 && keypad[y*5+x-1] != "" {
					x--
				}
			case 'D':
				if y < 4 && keypad[(y+1)*5+x] != "" {
					y++
				}
			case 'U':
				if y > 0 && keypad[(y-1)*5+x] != "" {
					y--
				}
			}
		}
		fmt.Printf("%s", keypad[y*5+x])
	}
	fmt.Printf("\n")
}
