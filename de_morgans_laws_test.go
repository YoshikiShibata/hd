// Copyright © 2019 Yoshiki Shibata. All rights reserved.

package hd

import (
	"math/rand"
	"testing"
)

type binary func(x, y int) int

func TestDeMorgansLaws(t *testing.T) {
	for _, test := range []struct {
		b1f binary
		b2f binary
	}{
		{
			b1f: func(x, y int) int { return ^(x & y) }, // ¬(x & y)
			b2f: func(x, y int) int { return ^x | ^y },  // ¬x | ¬y
		},
		{
			b1f: func(x, y int) int { return ^(x | y) }, // ¬(x | y)
			b2f: func(x, y int) int { return ^x & ^y },  // ¬x & ¬y
		},
		{
			b1f: func(x, _ int) int { return ^(x + 1) }, // ¬(x + 1)
			b2f: func(x, _ int) int { return ^x - 1 },   // ¬x - 1
		},
		{
			b1f: func(x, _ int) int { return ^(x - 1) }, // ¬(x - 1)
			b2f: func(x, _ int) int { return ^x + 1 },   // ¬x + 1
		},
		{
			b1f: func(x, _ int) int { return ^-x },   // ¬-x
			b2f: func(x, _ int) int { return x - 1 }, // x -1
		},
		{
			b1f: func(x, y int) int { return ^(x ^ y) }, // ¬(x ⊕ y)
			b2f: func(x, y int) int { return ^x ^ y },   // ¬x ⊕ y
		},
		{
			b1f: func(x, y int) int { return ^(x + y) }, // ¬(x + y)
			b2f: func(x, y int) int { return ^x - y },   // ¬x - y
		},
		{
			b1f: func(x, y int) int { return ^(x - y) }, // ¬(x - y)
			b2f: func(x, y int) int { return ^x + y },   // ¬x + y
		},

		// As an example of the application of these formulas,
		// ¬(x | –(x + 1)) = ¬x &¬–(x + 1) = ¬x & ((x + 1) – 1) = ¬x & x = 0.
		{
			b1f: func(x, y int) int { return ^(x | -(x + 1)) }, // ¬(x | –(x + 1))
			b2f: func(x, y int) int { return 0 },
		},
		{
			b1f: func(x, y int) int { return ^x & ^-(x + 1) }, // ¬x &¬–(x + 1)
			b2f: func(x, y int) int { return 0 },
		},
		{
			b1f: func(x, y int) int { return ^x & ((x + 1) - 1) }, // ¬x & ((x + 1) – 1)
			b2f: func(x, y int) int { return 0 },
		},
		{
			b1f: func(x, y int) int { return ^x & x }, // ¬x & x
			b2f: func(x, y int) int { return 0 },
		},
	} {
		for i := 0; i < 30; i++ {
			x := rand.Int()
			y := rand.Int()
			if test.b1f(x, y) != test.b2f(x, y) {
				t.Errorf("test.b1f(%d, %d) is %d, but want %d",
					x, y, test.b1f(x, y), test.b2f(x, y))
			}
		}
	}
}
