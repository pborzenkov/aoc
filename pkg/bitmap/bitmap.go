package bitmap

import (
	"errors"
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

var (
	// ErrOutOfBounds is returned when a bitmap operation results in out-of-bound access.
	ErrOutOfBounds = errors.New("out of bounds")
	// ErrMalformedInput is returned when a string to bitmap conversion fails
	ErrMalformedInput = errors.New("malformed input")
)

const wordSize = bits.UintSize

// Bitmap represents a single bitmap.
type Bitmap struct {
	size uint
	data []uint
}

// New returns a new instance of a bitmap capable of holding up to size bits.
// The bitmap is intialized with zeros.
func New(size uint) *Bitmap {
	words := (size + wordSize - 1) / wordSize

	return &Bitmap{
		size: size,
		data: make([]uint, words),
	}
}

// FromString returns a bitmap from its string representation
func FromString(s string) (*Bitmap, error) {
	b := New(uint(len(s)))
	r := strings.NewReader(s)
	bitsToScan := b.size % wordSize
	if bitsToScan == 0 {
		bitsToScan = wordSize
	}

	for i := len(b.data) - 1; i >= 0; i-- {
		buf := make([]byte, bitsToScan)
		if n, err := r.Read(buf); n != len(buf) || err != nil {
			return nil, ErrMalformedInput
		}
		n, err := strconv.ParseUint(string(buf), 2, 64)
		if err != nil {
			return nil, ErrMalformedInput
		}
		b.data[i] = uint(n)
		bitsToScan = wordSize
	}

	return b, nil
}

// String returns a string representation of the bitmap (least significant bit
// is on the right).
func (b *Bitmap) String() string {
	w := new(strings.Builder)
	bitsToPrint := b.size % wordSize
	if bitsToPrint == 0 {
		bitsToPrint = wordSize
	}

	for i := len(b.data) - 1; i >= 0; i-- {
		fmt.Fprintf(w, "%0[2]*[1]b", b.data[i], bitsToPrint)
		bitsToPrint = wordSize
	}

	return w.String()
}

// Set sets nbits starting from bit start.
func (b *Bitmap) Set(start, nbits uint) error {
	if start >= b.size || (start+nbits) > b.size || (start+nbits < start) {
		return ErrOutOfBounds
	}

	i := start / wordSize
	bitsToSet := int(wordSize - start%wordSize)
	mask := ^uint(0) << (start & (wordSize - 1))

	l := int(nbits)
	for l-bitsToSet >= 0 {
		b.data[i] |= mask
		l -= bitsToSet
		bitsToSet = wordSize
		mask = ^uint(0)
		i++
	}
	if l > 0 {
		mask &= ^uint(0) >> (-(start + nbits) & (wordSize - 1))
		b.data[i] |= mask
	}

	return nil
}

// Clear clears nbits starting from bit start.
func (b *Bitmap) Clear(start, nbits uint) error {
	if start >= b.size || (start+nbits) > b.size || (start+nbits < start) {
		return ErrOutOfBounds
	}

	i := start / wordSize
	bitsToClear := int(wordSize - start%wordSize)
	mask := ^uint(0) << (start & (wordSize - 1))

	l := int(nbits)
	for l-bitsToClear >= 0 {
		b.data[i] &^= mask
		l -= bitsToClear
		bitsToClear = wordSize
		mask = ^uint(0)
		i++
	}
	if l > 0 {
		mask &= ^uint(0) >> (-(start + nbits) & (wordSize - 1))
		b.data[i] &^= mask
	}

	return nil
}

// Not applies bitwise NOT operation to nbits starting from start.
func (b *Bitmap) Not(start, nbits uint) error {
	if start >= b.size || (start+nbits) > b.size || (start+nbits < start) {
		return ErrOutOfBounds
	}

	i := start / wordSize
	bitsToNot := int(wordSize - start%wordSize)
	mask := ^uint(0) << (start & (wordSize - 1))

	l := int(nbits)
	for l-bitsToNot >= 0 {
		b.data[i] ^= mask
		l -= bitsToNot
		bitsToNot = wordSize
		mask = ^uint(0)
		i++
	}
	if l > 0 {
		mask &= ^uint(0) >> (-(start + nbits) & (wordSize - 1))
		b.data[i] ^= mask
	}

	return nil
}

// IsSet returns true if the given bit is set.
func (b *Bitmap) IsSet(bit uint) bool {
	if bit >= b.size {
		return false
	}
	i := bit / wordSize
	return (b.data[i] & (uint(1) << (bit % wordSize))) != 0
}

// Weight returns a Hamming weight of the bitmap (number of set bits).
func (b *Bitmap) Weight() int {
	var weight int

	for _, d := range b.data {
		weight += bits.OnesCount(d)
	}

	return weight
}
