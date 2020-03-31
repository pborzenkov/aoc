package main

import (
	"bufio"
	"fmt"

	"github.com/pborzenkov/aoc/pkg/input"
)

type freqTracker struct {
	pos []map[rune]int
	top []rune
}

func (f *freqTracker) process(s string) {
	if f.pos == nil {
		f.pos = make([]map[rune]int, len(s))
		for i := 0; i < len(s); i++ {
			f.pos[i] = make(map[rune]int)
		}
		f.top = make([]rune, len(s))
	}

	for i, c := range s {
		f.pos[i][c]++
		if f.pos[i][c] > f.pos[i][f.top[i]] {
			f.top[i] = c
		}
	}
}

func (f *freqTracker) output() string {
	out := ""
	for _, c := range f.top {
		out += string(c)
	}
	return out
}

func main() {
	r := bufio.NewScanner(input.NewFileOrStdin())
	var f freqTracker

	for r.Scan() {
		f.process(r.Text())
	}

	fmt.Printf("Resulting string is %q\n", f.output())
}
