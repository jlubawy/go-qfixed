package qfixed

import (
	"testing"
)

const epsilon = 0.0001

func TestEncode(t *testing.T) {
}

func TestDecode(t *testing.T) {
	assert := func(f *Format, number uint16, expected float64) {
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
