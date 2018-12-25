package main

import (
	"context"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"golang.org/x/sync/errgroup"
)

type resultErr struct {
	num int
}

func (e *resultErr) Error() string {
	return fmt.Sprintf("result: %v", e.num)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %v <input>", os.Args[0])
	}

	numc := make(chan int)
	group, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < runtime.NumCPU(); i++ {
		group.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case num := <-numc:
					input := make([]byte, len(os.Args[1]))
					copy(input, os.Args[1])
					input = strconv.AppendInt(input, int64(num), 10)
					hash := md5.Sum(input)
					if binary.BigEndian.Uint32(hash[:])&0xfffff000 == 0 {
						return &resultErr{
							num: num,
						}
					}
				}
			}
		})
	}

loop:
	for num := 0; ; num++ {
		select {
		// sending number ranges will probably be more efficient, as
		// this one spends lots of time doing context switches...
		case numc <- num:
		case <-ctx.Done():
			break loop
		}
	}

	err := group.Wait()
	if rerr, ok := err.(*resultErr); ok {
		log.Printf("Found the answer: %v", rerr.num)
		return
	}
	log.Fatalf("Got unexpected error from the group: %v", err)
}
