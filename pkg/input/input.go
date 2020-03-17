package input

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// NewFileOrStdin returns input source either from the first file mentioned on
// the command line arguments, or from stdin if none are provided.
func NewFileOrStdin() io.ReadCloser {
	if len(os.Args) < 2 {
		return ioutil.NopCloser(os.Stdin)
	}

	r, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Cannot open input file: %v", err)
	}

	return r
}
