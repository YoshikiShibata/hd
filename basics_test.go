// Copyright © 2019 Yoshiki Shibata. All rights reserved.

package hd

import (
	"testing"
)

func TestTurnOffTheRightmost1Bit(t *testing.T) {
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
		r := TurnOffTheRightmost1Bit(test.x)
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

func TestTurnOnTheRIghtmost0Bit(t *testing.T) {
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
		r := TurnOnTheRightmost0Bit(test.x)
		if r != test.expected {
			t.Errorf("%x: result is %x, but want %x", test.x, r, test.expected)
		}
	}
}
