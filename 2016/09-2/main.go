package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/pborzenkov/aoc/pkg/input"
)

var (
	reMarker = regexp.MustCompile(`(\(\d+x\d+\))`)
)

func decompressedLength(in []byte) int {
	var length, chars, times int
	marker := reMarker.FindIndex(in)

	switch marker {
	case nil:
		length = len(in)
	default:
		fmt.Sscanf(string(in[marker[0]:marker[1]]), "(%dx%d)", &chars, &times)
		length = marker[0] + times*decompressedLength(in[marker[1]:marker[1]+chars]) +
			decompressedLength(in[marker[1]+chars:])
	}

	return length
}

func main() {
	in, err := ioutil.ReadAll(input.NewFileOrStdin())
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	l := decompressedLength(bytes.TrimSpace(in))

	fmt.Printf("Length is %d\n", l)
}
