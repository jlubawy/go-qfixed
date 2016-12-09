package qfixed

import (
	"errors"
	"fmt"
)

// Format is a Qm.n signed fixed-point number of format Qm.n where m is the
// integer width and n is the fractional width.
type Format struct {
	m     uint
	n     uint
	width uint
	mask  uint16
}

// NewFormat returns a fixed-point Qm.n format.
func NewFormat(m, n uint) *Format {
	width := uint(m + n)
	return &Format{
		m:     m,
		n:     n,
		width: width,
		mask:  mask(width),
	}
}

var Q9_4 = NewFormat(9, 4)
var Q8_4 = NewFormat(8, 4)
var Q14_0 = NewFormat(14, 0)
var Q12_3 = NewFormat(12, 3)

func (f *Format) String() string {
	return fmt.Sprintf("Q%d.%d", f.m, f.n)
}

// Encode converts a float64 number to it's Qm.n fixed-point representation.
func (f *Format) Encode(number float64) {
	panic(errors.New("not yet implemented"))
}

// Decode converts a Qm.m fixed-point number to its float64 representation.
func (f *Format) Decode(number uint16) (r float64) {
	if number > f.mask {
		panic(fmt.Errorf("qfixed: number %d is too large for format %s", number, f))
	}

	isNegative := (number&(1<<(f.width-1)) != 0)

	var magnitude uint16
	if isNegative {
		magnitude = (^number + 1) & f.mask
	} else {
		magnitude = number
	}

	res := 1 << f.n
	resolution := 1.0 / float64(res)

	result := float64(magnitude) * resolution

	if isNegative {
		r = -result
	} else {
		r = result
	}

	return
}

// Mask is a helper function which generates a mask of a specified width.
func mask(width uint) uint16 {
	return (1 << width) - 1
}
