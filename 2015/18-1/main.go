package main

import (
	"bufio"
	"log"

	"github.com/pborzenkov/aoc/pkg/bitmap"
	"github.com/pborzenkov/aoc/pkg/input"
)

const (
	side  = 100
	steps = 100
)

func main() {
	states := []*bitmap.Bitmap{
		bitmap.New(side * side),
		bitmap.New(side * side),
	}

	r := input.NewFileOrStdin()
	s := bufio.NewScanner(r)

	var line uint
	for s.Scan() {
		for i, c := range s.Text() {
			if c == '#' {
				states[0].Set(line*side+uint(i), 1)
			}
		}
		line++
	}

	for step := 0; step < steps; step++ {
		curState := states[step%2]
		nextState := states[(step+1)%2]

		for i := 0; i < side; i++ {
			for j := 0; j < side; j++ {
				on := 0

				for k := i - 1; k < i+2; k++ {
					for l := j - 1; l < j+2; l++ {
						if k >= 0 && l >= 0 && k < side && l < side && (k != i || l != j) {
							if curState.IsSet(uint(k*side + l)) {
								on++
							}
						}
					}
				}

				if curState.IsSet(uint(i*side + j)) {
					if on != 2 && on != 3 {
						nextState.Clear(uint(i*side+j), 1)
					} else {
						nextState.Set(uint(i*side+j), 1)
					}
				}
				if !curState.IsSet(uint(i*side + j)) {
					if on == 3 {
						nextState.Set(uint(i*side+j), 1)
					} else {
						nextState.Clear(uint(i*side+j), 1)
					}
				}

			}
		}
	}

	log.Printf("Total number of turned on lights after %d steps: %d", steps, states[steps%2].Weight())
}
