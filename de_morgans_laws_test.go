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
			b1f: func(x, y int) int { return ^(x & y) },
			b2f: func(x, y int) int { return ^x | ^y },
		},
		{
			b1f: func(x, y int) int { return ^(x | y) },
			b2f: func(x, y int) int { return ^x & ^y },
		},
		{
			b1f: func(x, _ int) int { return ^(x + 1) },
			b2f: func(x, _ int) int { return ^x - 1 },
		},
		{
			b1f: func(x, _ int) int { return ^(x - 1) },
			b2f: func(x, _ int) int { return ^x + 1 },
		},
		{
			b1f: func(x, _ int) int { return ^-x },
			b2f: func(x, _ int) int { return x - 1 },
		},
		{
			b1f: func(x, y int) int { return ^(x ^ y) },
			b2f: func(x, y int) int { return ^x ^ y },
		},
	} {
		for i := 0; i < 10; i++ {
			x := rand.Int()
			y := rand.Int()
			if test.b1f(x, y) != test.b2f(x, y) {
				t.Errorf("test.b1f(%d, %d) is %d, but want %d",
					x, y, test.b1f(x, y), test.b2f(x, y))
			}
		}
	}
}
