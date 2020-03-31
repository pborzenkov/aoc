package main

import (
	"bufio"
	"fmt"
	"sort"
	"unicode"

	"github.com/pborzenkov/aoc/pkg/input"
)

type code struct {
	c rune
	f int
}

func decrypt(s string, shift int) string {
	res := ""

	for _, c := range s {
		if c == '-' {
			res += " "
		} else {
			res += string((int(c-'a')+shift)%('z'-'a'+1) + 'a')
		}
	}
	return res
}

func main() {
	r := bufio.NewScanner(input.NewFileOrStdin())

	for r.Scan() {
		freq := make(map[rune]int)
		seenChecksumStart := false
		var sum, room string
		id := 0

		for i, s := range r.Text() {
			switch {
			case s == '-' || s == ']':
			case s == '[':
				seenChecksumStart = true
			case unicode.IsNumber(s):
				if room == "" {
					room = r.Text()[:i]
				}
				id = id*10 + int(s-'0')
			case seenChecksumStart:
				sum += string(s)
			default:
				freq[s]++
			}
		}

		codes := make([]code, 0, len(freq))
		for k, v := range freq {
			codes = append(codes, code{k, v})
		}
		sort.Slice(codes, func(i, j int) bool {
			if codes[j].f < codes[i].f {
				return true
			}
			if codes[j].f > codes[i].f {
				return false
			}
			return codes[i].c < codes[j].c
		})

		valid := true
		for i, c := range sum {
			if c != codes[i].c {
				valid = false
				break
			}
		}
		if valid {
			fmt.Printf("Room %q, ID %d\n", decrypt(room, id), id)
		}
	}
}
