package main

import (
	"log"
	"io/ioutil"
	"strconv"

	"github.com/pborzenkov/aoc/pkg/input"
)

func main() {
	data, err := ioutil.ReadAll(input.NewFileOrStdin())
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	
	start := -1
	var sum int64
	for i := 0; i < len(data); i++ {
		if (data[i] == '-' || (data[i] >= '0' && data[i] <= '9')) && start == -1 {
			start = i
		} else if (data[i] < '0' || data[i] > '9') && start != -1 {
			str := string(data[start:i])
			log.Printf("parsing %q", str)

			num, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				log.Fatalf("Failed to parse number %q: %v", str, err)
			}
			sum += num
			start = -1
		}
	}

	log.Printf("Total sum is %d", sum)
}
