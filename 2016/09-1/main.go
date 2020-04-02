package main

import (
	"fmt"
	"strings"

	"github.com/pborzenkov/aoc/pkg/input"
)

func main() {
	var in string
	fmt.Fscanf(input.NewFileOrStdin(), "%s", &in)

	var out, repeated strings.Builder
	var inMarker, scanTimes bool
	var chars, times int
	for i := 0; i < len(in); i++ {
		switch {
		case in[i] == '(' && chars == 0:
			inMarker = true
		case in[i] == 'x' && inMarker:
			scanTimes = true
		case in[i] >= '0' && in[i] <= '9' && inMarker && !scanTimes:
			chars = chars*10 + int(in[i]-'0')
		case in[i] >= '0' && in[i] <= '9' && inMarker && scanTimes:
			times = times*10 + int(in[i]-'0')
		case in[i] == ')' && inMarker:
			inMarker = false
			scanTimes = false
		case !inMarker && chars > 0:
			repeated.WriteByte(in[i])
			chars--
			for ; chars == 0 && times > 0; times-- {
				out.WriteString(repeated.String())
			}
			if chars == 0 {
				repeated.Reset()
			}
		default:
			out.WriteByte(in[i])
		}
	}

	fmt.Printf("Length is %d\n", out.Len())
}
