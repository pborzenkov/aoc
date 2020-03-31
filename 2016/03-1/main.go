package main

import (
	"bufio"
	"fmt"

	"github.com/pborzenkov/aoc/pkg/input"
)

func sumAndMax(n []int) (sum int, max int) {
	for _, i := range n {
		sum += i
		if i > max {
			max = i
		}
	}
	return
}

func main() {
	r := bufio.NewScanner(input.NewFileOrStdin())

	var valid int
	for r.Scan() {
		var i1, i2, i3 int
		fmt.Sscanf(r.Text(), "%d %d %d", &i1, &i2, &i3)

		sum, max := sumAndMax([]int{i1, i2, i3})
		if sum-max > sum/2 {
			valid++
		}
	}
	fmt.Printf("%d valid triangles\n", valid)
}
