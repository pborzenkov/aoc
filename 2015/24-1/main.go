package main

import (
	"bufio"
	"fmt"
	"log"
	"math"

	"github.com/gonum/stat/combin"
	"github.com/pborzenkov/aoc/pkg/input"
)

func main() {
	r := input.NewFileOrStdin()
	s := bufio.NewScanner(r)

	var numbers []int
	var target int
	for s.Scan() {
		var num int
		fmt.Sscanf(s.Text(), "%d", &num)
		numbers = append(numbers, num)
		target += num
	}
	target /= 3

	log.Printf("Target sum is %d", target)

	minLen := len(numbers) + 1
	var minQE int64 = math.MaxInt64
	for i := 1; i < len(numbers)/3 && i < minLen; i++ {
		g := combin.NewCombinationGenerator(len(numbers), i)
		data := make([]int, i)
		for g.Next() {
			data := g.Combination(data)
			sum := 0
			qe := int64(1)
			for _, v := range data {
				sum += numbers[v]
				qe *= int64(numbers[v])
			}
			if sum == target && qe < minQE {
				minLen = i
				minQE = qe
			}
		}
	}

	log.Printf("Min QE is %d", minQE)
}
