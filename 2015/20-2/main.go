package main

import (
	"fmt"
	"log"

	"github.com/pborzenkov/aoc/pkg/input"
)

func main() {
	var target int
	fmt.Fscanf(input.NewFileOrStdin(), "%d", &target)

	houses := make([]int, target/11)
	for i := 1; i < target/11; i++ {
		for j, k := i, 0; j < target/11 && k < 50; j, k = j+i, k+1 {
			houses[j] += i * 11
		}
	}
	for i := 0; i < target/11; i++ {
		if houses[i] > target {
			log.Printf("House %d, presents %d", i, houses[i])
			return
		}
	}
}
