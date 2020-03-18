package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/pborzenkov/aoc/pkg/input"
)

func strtonum(str string) uint64 {
	var out uint64 = 0
	base := uint64('z' - 'a' + 1)

	for i := 0; i < len(str); i++ {
		if i > 0 {
			out *= base
		}
		out += uint64(str[i] - 'a')
	}

	return out
}

func numtostr(num uint64) string {
	var out string
	base := uint64('z' - 'a' + 1)

	for num != 0 {
		out = string(rune('a'+(num%base))) + out
		num /= base
	}
	if len(out) < 8 {
		for i := len(out); i < 8; i++ {
			out = "a" + out
		}
	}

	return out
}

func next(str string) string {
	if str == "zzzzzzzz" {
		return str
	}

	return numtostr(strtonum(str) + 1)
}

func good(str string) bool {
	var last, lastPair rune
	var num, pairs int

	var three bool
	for i, c := range str {
		if c == last && c != lastPair {
			pairs++
			lastPair = c
		}
		if i > 0 && c == last+1 {
			num++
		} else {
			num = 1
		}
		last = c
		if num >= 3 {
			three = true
		}
	}
	forbiddenChars := strings.ContainsAny(str, "iol")

	return !forbiddenChars && three && pairs >= 2
}

func main() {
	var str string

	r := input.NewFileOrStdin()
	fmt.Fscanf(r, "%s", &str)

	for {
		if str == "zzzzzzzz" {
			log.Fatal("couldn't find valid password")
		}

		if good(str) {
			break
		}
		str = next(str)
	}

	log.Printf("Next valid password is %q", str)
}
