// Copyright © 2019 Yoshiki Shibata. All rights reserved.

package hd

func TurnOffTheRightmost1Bit(x int) int {
	return x & (x - 1)
}

func IsPowerOf2(x uint) bool {
	return (x & (x - 1)) == 0
}

func TurnOnTheRightmost0Bit(x int) int {
	return x | (x + 1)
}
