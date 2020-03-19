package bitmap_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/pborzenkov/aoc/pkg/bitmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromString(t *testing.T) {
	var tests = []struct {
		input string
		err   error
	}{
		{"001", nil},
		{"00000010000000000000000000010000", nil},
		{"000000000000001000000000000000100", nil},
		{"0010000011010000111000000000000000000000000000000000000000000000", nil},
		{"00000111000000000000000000000000000000000000000000000000000001111", nil},
		{"", nil},
		{"00010110110101a0011011", bitmap.ErrMalformedInput},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%d", i+1), func(t *testing.T) {
			bmap, err := bitmap.FromString(test.input)
			if test.err != nil {
				assert.Equal(t, test.err, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.input, bmap.String())
		})
	}
}

func TestString(t *testing.T) {
	var tests = []struct {
		get func() *bitmap.Bitmap
		str string
	}{
		{func() *bitmap.Bitmap { return bitmap.New(3) }, "000"},
		{func() *bitmap.Bitmap { return bitmap.New(32) }, "00000000000000000000000000000000"},
		{func() *bitmap.Bitmap { return bitmap.New(33) }, "000000000000000000000000000000000"},
		{func() *bitmap.Bitmap { return bitmap.New(64) },
			"0000000000000000000000000000000000000000000000000000000000000000"},
		{func() *bitmap.Bitmap { return bitmap.New(65) },
			"00000000000000000000000000000000000000000000000000000000000000000"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%d", i+1), func(t *testing.T) {
			bmap := test.get()
			assert.Equal(t, test.str, bmap.String())
		})
	}
}

func TestSet(t *testing.T) {
	var size uint = 65

	var tests = []struct {
		start uint
		size  uint
		err   error
		str   string
	}{
		{0, 65, nil, "11111111111111111111111111111111111111111111111111111111111111111"},
		{0, 16, nil, "00000000000000000000000000000000000000000000000001111111111111111"},
		{60, 5, nil, "11111000000000000000000000000000000000000000000000000000000000000"},
		{64, 1, nil, "10000000000000000000000000000000000000000000000000000000000000000"},
		{65, 1, bitmap.ErrOutOfBounds, ""},
		{65, 0, bitmap.ErrOutOfBounds, ""},
		{30, 0, nil, "00000000000000000000000000000000000000000000000000000000000000000"},
		{62, 3, nil, "11100000000000000000000000000000000000000000000000000000000000000"},
		{15, math.MaxUint64, bitmap.ErrOutOfBounds, ""},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%d", i+1), func(t *testing.T) {
			bmap := bitmap.New(size)
			err := bmap.Set(test.start, test.size)
			if test.err != nil {
				assert.Equal(t, test.err, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.str, bmap.String())
		})
	}
}

func TestClear(t *testing.T) {
	var size uint = 65

	var tests = []struct {
		start uint
		size  uint
		err   error
		str   string
	}{
		{0, 65, nil, "00000000000000000000000000000000000000000000000000000000000000000"},
		{0, 16, nil, "11111111111111111111111111111111111111111111111110000000000000000"},
		{60, 5, nil, "00000111111111111111111111111111111111111111111111111111111111111"},
		{64, 1, nil, "01111111111111111111111111111111111111111111111111111111111111111"},
		{65, 1, bitmap.ErrOutOfBounds, ""},
		{65, 0, bitmap.ErrOutOfBounds, ""},
		{30, 0, nil, "11111111111111111111111111111111111111111111111111111111111111111"},
		{62, 3, nil, "00011111111111111111111111111111111111111111111111111111111111111"},
		{15, math.MaxUint64, bitmap.ErrOutOfBounds, ""},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%d", i+1), func(t *testing.T) {
			bmap := bitmap.New(size)
			require.NoError(t, bmap.Set(0, size))
			err := bmap.Clear(test.start, test.size)
			if test.err != nil {
				assert.Equal(t, test.err, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.str, bmap.String())
		})
	}
}

func TestNot(t *testing.T) {
	var tests = []struct {
		input string
		start uint
		size  uint
		err   error
		str   string
	}{
		{
			"10101010101010101010101010101010101010101010101010101010101010101",
			0, 65, nil,
			"01010101010101010101010101010101010101010101010101010101010101010",
		},
		{
			"10101010101010101010101010101010101010101010101010101010101010101",
			65, 1, bitmap.ErrOutOfBounds, "",
		},
		{
			"10101010101010101010101010101010101010101010101010101010101010101",
			65, 0, bitmap.ErrOutOfBounds, "",
		},
		{
			"10101010101010101010101010101010101010101010101010101010101010101",
			15, math.MaxUint64, bitmap.ErrOutOfBounds, "",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%d", i+1), func(t *testing.T) {
			bmap, err := bitmap.FromString(test.input)
			require.NoError(t, err)
			err = bmap.Not(test.start, test.size)
			if test.err != nil {
				assert.Equal(t, test.err, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.str, bmap.String())
		})
	}
}

func TestWeight(t *testing.T) {
	var tests = []struct {
		input  string
		weight int
	}{
		{"10101010101010101010101010101010101010101010101010101010101010101", 33},
		{"00000000000000000000000000000000000000000000000000000000000000000", 0},
		{"11111111111111111111111111111111111111111111111111111111111111111", 65},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%d", i+1), func(t *testing.T) {
			bmap, err := bitmap.FromString(test.input)
			require.NoError(t, err)
			weight := bmap.Weight()
			assert.Equal(t, test.weight, weight)
		})
	}
}
