package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/pborzenkov/aoc/pkg/input"
)

func appendstr(bldr *strings.Builder, char rune, num int) {
	if num > 0 {
		fmt.Fprintf(bldr, "%d", num)
		fmt.Fprintf(bldr, "%s", string(char))
	}
}

func main() {
	var seq string
	var iters int

	r := input.NewFileOrStdin()
	fmt.Fscanf(r, "%s %d", &seq, &iters)

	for i := 0; i < iters; i++ {
		var out strings.Builder
		var num int
		var char rune

		for _, c := range seq {
			if c != char {
				appendstr(&out, char, num)
				char = c
				num = 1
			} else {
				num++
			}
		}
		appendstr(&out, char, num)
		seq = out.String()
	}

	log.Printf("Result length is %d", len(seq))
}
