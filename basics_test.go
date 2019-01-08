// Copyright © 2019 Yoshiki Shibata. All rights reserved.

package hd

import (
	"testing"
)

func TestTurnOffRightMost1Bit(t *testing.T) {
	minInt := func() int {
		bits := 32 << (^uint(0) >> 63)
		return -1 << uint((bits - 1))
	}

	for _, test := range []struct {
		x        int
		expected int
	}{
		{0, 0},
		{-1, -1 &^ 1},
		{minInt(), 0},
	} {
		r := TurnOffRightMost1Bit(test.x)
		if r != test.expected {
			t.Errorf("%x: result is %x, but want %x", test.x, r, test.expected)
		}
	}
}

func TestIsPowerOf2(t *testing.T) {
	for _, test := range []struct {
		x        uint
		expected bool
	}{
		{0, false},
		{1, true},
		{1 << 10, true},
		{3, false},
		{1<<10 + 1, false},
	} {
		r := IsPowerOf2(test.x)
		if r != test.expected {
			t.Errorf("%x: result is %t, but want %t", test.x, r, test.expected)
		}
	}
}

func TestTurnOnRightMost0Bit(t *testing.T) {
	minInt := func() int {
		bits := 32 << (^uint(0) >> 63)
		return -1 << uint((bits - 1))
	}

	for _, test := range []struct {
		x        int
		expected int
	}{
		{0, 1},
		{2, 3},
		{minInt(), minInt() + 1},
	} {
		r := TurnOnRightMost0Bit(test.x)
		if r != test.expected {
			t.Errorf("%x: result is %x, but want %x", test.x, r, test.expected)
		}
	}
}

func TestTurnOffTrailing1s(t *testing.T) {
	for _, test := range []struct {
		x        int
		expected int
	}{
		{0, 0},
		{-1, 0},
		{1, 0},
		{0xFF, 0},
		{0xFFFF00FF, 0xFFFF0000},
		{0xFF00FFFF, 0xFF000000},
	} {
		r := TurnOffTrailing1s(test.x)
		if r != test.expected {
			t.Errorf("%x: result is %x, but want %x", test.x, r, test.expected)
		}
	}
}

func TestTurnOnTrailing0s(t *testing.T) {
	for _, test := range []struct {
		x        int
		expected int
	}{
		{-1, -1},
		{0, -1},
		{0xF0, 0xFF},
		{0xFF00, 0xFFFF},
		{0xF0FF00, 0xF0FFFF},
	} {
		r := TurnOnTrailing0s(test.x)
		if r != test.expected {
			t.Errorf("%x: result is %x, but want %x", test.x, r, test.expected)
		}
	}
}

func TestSingle1BitAtRightMost0Bit(t *testing.T) {
	for _, test := range []struct {
		x        int
		expected int
	}{
		{-1, 0},
		{0xFF00, 1},
		{0xFF01, 2},
	} {
		r := Single1BitAtRightMost0Bit(test.x)
		if r != test.expected {
			t.Errorf("%x: result is %x, but want %x", test.x, r, test.expected)
		}
	}
}

func TestSingle0BitAtRightMost1Bit(t *testing.T) {
	for _, test := range []struct {
		x        int
		expected int
	}{
		{0, -1}, // all 1's if none
		{1, -2},
		{0xA8, -1 &^ (1 << 3)},
	} {
		r := Single0BitAtRightMost1Bit(test.x)
		if r != test.expected {
			t.Errorf("%x: result is %x, but want %x", test.x, r, test.expected)
		}
	}
}

func TestReplaceTrailing0sWith1s(t *testing.T) {
	for _, test := range []struct {
		x        int
		expected int
	}{
		{-1, 0},
		{0x58, 0x7},
		{0, -1},
	} {
		r := ReplaceTrailing0sWith1sA(test.x)
		if r != test.expected {
			t.Errorf("%x: result is %x, but want %x", test.x, r, test.expected)
		}

		r = ReplaceTrailing0sWith1sB(test.x)
		if r != test.expected {
			t.Errorf("%x: result is %x, but want %x", test.x, r, test.expected)
		}

		r = ReplaceTrailing0sWith1sC(test.x)
		if r != test.expected {
			t.Errorf("%x: result is %x, but want %x", test.x, r, test.expected)
		}
	}
}

var finalResult int

func BenchmarkReplaceTrailing0sWith1sA(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result += ReplaceTrailing0sWith1sA(i)
	}
	finalResult = result
}

func BenchmarkReplaceTrailing0sWith1sB(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result += ReplaceTrailing0sWith1sB(i)
	}
	finalResult = result
}

func BenchmarkReplaceTrailing0sWith1sC(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result += ReplaceTrailing0sWith1sC(i)
	}
	finalResult = result
}
