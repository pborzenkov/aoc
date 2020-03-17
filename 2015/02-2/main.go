package main

import (
	"bufio"
	"fmt"
	"log"

	"github.com/pborzenkov/aoc/pkg/input"
)

func max(val int, vals ...int) int {
	mv := val
	for _, v := range vals {
		if v > mv {
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

		total += uint64((w+h+l-max(w, h, l))*2 + w*h*l)
	}
	if s.Err() != nil {
		log.Fatalf("Scanning completed with error (%v), result might be unreliable", s.Err())
	}
	log.Printf("Total feet of ribbon required: %v", total)
}
