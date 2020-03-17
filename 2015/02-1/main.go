package main

import (
	"bufio"
	"fmt"
	"log"

	"github.com/pborzenkov/aoc/pkg/input"
)

func min(val int, vals ...int) int {
	mv := val
	for _, v := range vals {
		if v < mv {
			mv = v
		}
	}
	return mv
}

func main() {
	r := input.NewFileOrStdin()
	defer r.Close()

	s := bufio.NewScanner(r)

	var total uint64
	for s.Scan() {
		var w, h, l int

		_, err := fmt.Sscanf(s.Text(), "%dx%dx%d", &w, &h, &l)
		if err != nil {
			log.Fatalf("Could not fmt.Sscanf line %q: %v", s.Text(), err)
		}

		s1, s2, s3 := w*h, w*l, h*l
		total += uint64(2*(s1+s2+s3) + min(s1, s2, s3))
	}
	if s.Err() != nil {
		log.Fatalf("Scanning completed with error (%v), result might be unreliable", s.Err())
	}
	log.Printf("Total square feet of wrapping paper required: %v", total)
}
