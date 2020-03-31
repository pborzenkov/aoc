package main

import (
	"bufio"
	"fmt"

	"github.com/pborzenkov/aoc/pkg/input"
)

type freqTracker struct {
	pos []map[rune]int
}

func (f *freqTracker) process(s string) {
	if f.pos == nil {
		f.pos = make([]map[rune]int, len(s))
		for i := 0; i < len(s); i++ {
			f.pos[i] = make(map[rune]int)
		}
	}

	for i, c := range s {
		f.pos[i][c]++
	}
}

func (f *freqTracker) output() string {
	out := ""
	for _, p := range f.pos {
		var c rune
		var count int

		for k, v := range p {
			if count == 0 || v < count {
				c = k
				count = v
			}
		}
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
