package main

import (
	"bufio"
	"fmt"

	"github.com/pborzenkov/aoc/pkg/input"
)

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	x := 1
	y := 1

	r := bufio.NewScanner(input.NewFileOrStdin())
	for r.Scan() {
		for _, c := range r.Text() {
			switch c {
			case 'R':
				x = min(x+1, 2)
			case 'L':
				x = max(x-1, 0)
			case 'D':
				y = min(y+1, 2)
			case 'U':
				y = max(y-1, 0)
			}
		}
		fmt.Printf("%d", y*3+x+1)
	}
	fmt.Printf("\n")
}
