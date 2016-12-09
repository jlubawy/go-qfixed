package qfixed

import (
	"testing"
)

func TestEncode(t *testing.T) {
}

func TestDecode(t *testing.T) {
	assert := func(f *Format, number uint16, expected float64) {
		actual := f.Decode(number)
		if expected != actual {
			t.Errorf("TestDecode: expected %f but got %f", expected, actual)
		}
	}

	// Q8.4
	assert(Q8_4, 0x000, +000.0000)
	assert(Q8_4, 0x800, -128.0000)
	assert(Q8_4, 0x7FF, +127.9375)
}
