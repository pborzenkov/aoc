package main

import (
	"fmt"
	"log"

	"github.com/pborzenkov/aoc/pkg/input"
)

func main() {
	var target int
	fmt.Fscanf(input.NewFileOrStdin(), "%d", &target)

	houses := make([]int, target/10)
	for i := 1; i < target/10; i++ {
		for j := i; j < target/10; j += i {
			houses[j] += i * 10
		}
	}
	for i := 0; i < target/10; i++ {
		if houses[i] > target {
			log.Printf("House %d, presents %d", i, houses[i])
			return
		}
	}
}
