package main

import (
	"fmt"
	"log"

	"github.com/pborzenkov/aoc/pkg/input"
)

func main() {
	r := input.NewFileOrStdin()

	var row, column int
	fmt.Fscanf(r, "To continue, please consult the code grid in the manual.  Enter the code at row %d, column %d.", &row, &column)
	diagonal := row + column - 1

	num := (diagonal*(diagonal-1))/2 + column

	log.Printf("Number #%d", num)

	cur := int64(20151125)
	for i := 0; i < num-1; i++ {
		cur = (cur * 252533) % 33554393
	}

	log.Printf("Code is %d", cur)
}
