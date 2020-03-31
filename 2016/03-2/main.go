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

	nums := make([][]int, 3)
	for i := 0; i < 3; i++ {
		nums[i] = make([]int, 3)
	}

	var idx int
	for r.Scan() {
		fmt.Sscanf(r.Text(), "%d %d %d", &nums[0][idx], &nums[1][idx], &nums[2][idx])
		if idx == 2 {
			for i := 0; i < 3; i++ {
				sum, max := sumAndMax(nums[i])
				if sum-max > sum/2 {
					valid++
				}
			}
		}
		idx = (idx + 1) % 3
	}
	fmt.Printf("%d valid triangles\n", valid)
}
