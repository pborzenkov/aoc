package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/pborzenkov/aoc/pkg/input"
)

type scope struct {
	sum     int64
	char    byte
	has_red bool
}

type stack []*scope

func (s *stack) push(char byte) {
	*s = append(*s, &scope{
		char: char,
	})
}

func (s *stack) pop() *scope {
	sc := (*s)[len(*s)-1]

	*s = (*s)[:len(*s)-1]

	return sc
}

func (s stack) last() *scope {
	return s[len(s)-1]
}

func main() {
	data, err := ioutil.ReadAll(input.NewFileOrStdin())
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	var s stack
	// dummy scope to hold total sum
	s.push('{')
	for i := 0; i < len(data); i++ {
		d := data[i]

		switch {
		case d == '[' || d == '{':
			s.push(d)
		case d == ']' || d == '}':
			sc := s.pop()
			if sc.char == '[' || !sc.has_red {
				s.last().sum += sc.sum
			}
		case d == '-' || d >= '0' && d <= '9':
			start := i

			for ; i+1 < len(data) && data[i+1] >= '0' && data[i+1] <= '9'; i++ {
			}
			num, err := strconv.ParseInt(string(data[start:i+1]), 10, 64)
			if err != nil {
				log.Fatalf("failed to parse %s: %v", string(data[start:i+1]), err)
			}
			s.last().sum += num
		case d == '"':
			start := i + 1
			i++

			for ; i < len(data) && data[i] != '"'; i++ {
			}
			if bytes.Compare(data[start:i], []byte("red")) == 0 {
				s.last().has_red = true
			}
		}
	}

	log.Printf("Total sum is %d", s.last().sum)
}
