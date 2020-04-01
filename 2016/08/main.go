package main

import (
	"bufio"
	"fmt"

	"github.com/pborzenkov/aoc/pkg/input"
)

const (
	width  = 50
	height = 6
)

func dumpScreen(s [height][width]bool) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if s[i][j] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	r := bufio.NewScanner(input.NewFileOrStdin())

	var screen [height][width]bool
	for r.Scan() {
		var x, y, count int

		if n, err := fmt.Sscanf(r.Text(), "rect %dx%d", &x, &y); err == nil && n == 2 {
			for i := 0; i < y; i++ {
				for j := 0; j < x; j++ {
					screen[i][j] = true
				}
			}
		} else if n, err := fmt.Sscanf(r.Text(), "rotate column x=%d by %d", &x, &count); err == nil && n == 2 {
			var tmp [height]bool
			for j := 0; j < height; j++ {
				tmp[j] = screen[(j+height-count)%height][x]
			}
			for j := 0; j < height; j++ {
				screen[j][x] = tmp[j]
			}
		} else if n, err := fmt.Sscanf(r.Text(), "rotate row y=%d by %d", &y, &count); err == nil && n == 2 {
			var tmp [width]bool
			for i := 0; i < width; i++ {
				tmp[i] = screen[y][(i+width-count)%width]
			}
			for i := 0; i < width; i++ {
				screen[y][i] = tmp[i]
			}
		}
	}

	enabled := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if screen[i][j] {
				enabled++
			}
		}
	}

	fmt.Printf("Total number of enabled pixels is %d\n", enabled)

	dumpScreen(screen)
}
