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
		{0x0FFF00FF, 0x0FFF0000},
		{0x0F00FFFF, 0x0F000000},
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

func TestReplaceTrailing1sWith0s(t *testing.T) {
	for _, test := range []struct {
		x        int
		expected int
	}{
		{0, -1}, // all 1's if none
		{0xA7, -8},
		{0x0F000007, -8},
		{0x7FFFFFFD, -2},
	} {
		r := ReplaceTrailing1sWith0s(test.x)
		if r != test.expected {
			t.Errorf("%x: result is %x, but want %x", test.x, r, test.expected)
		}
	}
}

func TestIsolateRightMost1Bit(t *testing.T) {
	for _, test := range []struct {
		x        int
		expected int
	}{
		{0, 0}, // 0 if none
		{0x58, 0x08},
		{0x70000000, 0x10000000},
	} {
		r := IsolateRightMost1Bit(test.x)
		if r != test.expected {
			t.Errorf("%x: result is %x, but want %x", test.x, r, test.expected)
		}
	}
}

func TestReplaceRightMost1BitAndTrailing0sWith1s(t *testing.T) {
	for _, test := range []struct {
		x        int
		expected int
	}{
		{0, -1}, // all 1's if no 1-bit
		{-1, 1}, // the integer 1 if no trailing 0's
		{0x58, 0xF},
	} {
		r := ReplaceRightMost1BitAndTrailing0sWith1s(test.x)
		if r != test.expected {
			t.Errorf("%x: result is %x, but want %x", test.x, r, test.expected)
		}
	}
}

func TestReplaceRightMost0BitAndTrailing1sWith0s(t *testing.T) {
	for _, test := range []struct {
		x        int
		expected int
	}{
		{-1, -1}, // all 1's if no 0-bit
		{0, 1},   // the integer 1 if no trailing 1's
		{0x57, 0x0F},
	} {
		r := ReplaceRightMost0BitAndTrailing1sWith0s(test.x)
		if r != test.expected {
			t.Errorf("%x: result is %x, but want %x", test.x, r, test.expected)
		}
	}
}

func TestTurnOffRightMostContinousStringOf1s(t *testing.T) {
	for _, test := range []struct {
		x        int
		expected int
	}{
		{0x5C, 0x40},
		{1<<4 - 1<<2, 0},
		{1<<15 - 1<<10, 0},
	} {
		r := TurnOffRightMostContinousStringOf1s(test.x)
		if r != test.expected {
			t.Errorf("%x: result is %x, but want %x", test.x, r, test.expected)
		}
	}
}

func TestTurnOffRightMostContinousStringOf1sAnother(t *testing.T) {
	for _, test := range []struct {
		x        int
		expected int
	}{
		{0x5C, 0x40},
		{1<<4 - 1<<2, 0},
		{1<<15 - 1<<10, 0},
	} {
		r := TurnOffRightMostContinousStringOf1sAnother(test.x)
		if r != test.expected {
			t.Errorf("%x: result is %x, but want %x", test.x, r, test.expected)
		}
	}
}

func BenchmarkTurnOffRightMostContinousStringOf1s(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result += TurnOffRightMostContinousStringOf1s(i)
	}
	finalResult = result
}

func BenchmarkTurnOffRightMostContinousStringOf1sAnother(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result += TurnOffRightMostContinousStringOf1sAnother(i)
	}
	finalResult = result
}
