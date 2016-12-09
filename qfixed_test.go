package qfixed

import (
	"testing"
)

const epsilon = 0.0001

func TestEncode(t *testing.T) {
	assert := func(f *Format, number float64, expected Number) {
		actual := f.Encode(number)
		t.Logf("Encode: format=%s, number=%.4f, expected=0x%04X, actual=0x%04X", f, number, expected, actual)
		if expected != actual {
			t.Errorf("TestEncode: expected 0x%04X but got 0x%04X", expected, actual)
		}
	}

	// Q9.4
	assert(Q9_4, +000.0000, 0x0000)
	assert(Q9_4, -256.0000, 0x1000)
	assert(Q9_4, +255.9375, 0x0FFF)
	assert(Q9_4, -000.0625, 0x1FFF)
	assert(Q9_4, +002.0000, 0x0020)
	assert(Q9_4, +123.0625, 0x07B1)
	assert(Q9_4, -122.9375, 0x1851)

	// Q14.0
	assert(Q14_0, +0000.0, 0x0000)
	assert(Q14_0, -8192.0, 0x2000)
	assert(Q14_0, +8191.0, 0x1FFF)

	// Q8.4
	assert(Q8_4, +000.0000, 0x000)
	assert(Q8_4, -128.0000, 0x800)
	assert(Q8_4, +127.9375, 0x7FF)

	// Q12.3
	assert(Q12_3, +0000.000, 0x0000)
	assert(Q12_3, -2048.000, 0x4000)
	assert(Q12_3, +2047.875, 0x3FFF)

	// Q15.1
	assert(Q15_1, +00000.0, 0x0000)
	assert(Q15_1, -16384.0, 0x8000)
	assert(Q15_1, +16383.5, 0x7FFF)
}

func TestDecode(t *testing.T) {
	assert := func(f *Format, number Number, expected float64) {
		actual := f.Decode(number)
		t.Logf("Decode: format=%s, number=0x%04X, expected=%.4f, actual=%.4f", f, number, expected, actual)
		if absDiff(expected, actual) > epsilon {
			t.Errorf("TestDecode: expected %f but got %f", expected, actual)
		}
	}

	// Q9.4
	assert(Q9_4, 0x0000, +000.0000)
	assert(Q9_4, 0x1000, -256.0000)
	assert(Q9_4, 0x0FFF, +255.9375)
	assert(Q9_4, 0x1FFF, -000.0625)
	assert(Q9_4, 0x0020, +002.0000)
	assert(Q9_4, 0x07B1, +123.0625)
	assert(Q9_4, 0x1851, -122.9375)

	// Q14.0
	assert(Q14_0, 0x0000, +0000.0)
	assert(Q14_0, 0x2000, -8192.0)
	assert(Q14_0, 0x1FFF, +8191.0)

	// Q8.4
	assert(Q8_4, 0x000, +000.0000)
	assert(Q8_4, 0x800, -128.0000)
	assert(Q8_4, 0x7FF, +127.9375)

	// Q12.3
	assert(Q12_3, 0x0000, +0000.000)
	assert(Q12_3, 0x4000, -2048.000)
	assert(Q12_3, 0x3FFF, +2047.875)

	// Q15.1
	assert(Q15_1, 0x0000, +00000.0)
	assert(Q15_1, 0x8000, -16384.0)
	assert(Q15_1, 0x7FFF, +16383.5)
}

func absDiff(a, b float64) float64 {
	if a < 0.0 {
		a = -a
	}
	if b < 0.0 {
		b = -b
	}
	return a - b
}
