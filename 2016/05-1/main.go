package main

import (
	"bytes"
	"crypto/md5"
	"fmt"

	"github.com/pborzenkov/aoc/pkg/input"
)

func main() {
	var doorID string

	fmt.Fscanf(input.NewFileOrStdin(), "%s", &doorID)

	var buf bytes.Buffer
	var i, found int
	var pass string
	for found < 8 {
		fmt.Fprintf(&buf, "%s%d", doorID, i)
		hash := md5.Sum(buf.Bytes())
		if hash[0] == 0 && hash[1] == 0 && hash[2] <= 0xf {
			found++
			pass += fmt.Sprintf("%x", hash[2])
		}
		buf.Reset()
		i++
	}
	fmt.Printf("Password is %q\n", pass)
}
