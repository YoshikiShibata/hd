// Copyright © 2019 Yoshiki Shibata. All rights reserved.

package hd

// TurnOffRightMost1Bit turns off the rightmost 1-bit in a word, producing 0 if
// none (e.g., 01011000 ⇒ 01010000).
func TurnOffRightMost1Bit(x int) int {
	return x & (x - 1)
}

// IsPowerOf2 determines if an unsigned integer is a power of 2.
func IsPowerOf2(x uint) bool {
	if x == 0 {
		return false
	}
	return (x & (x - 1)) == 0
}

// TurnOfRightMost0Bit turns on the rightmost 0-bit in a word, producing all 1’s
// if none (e.g., 10100111 ⇒ 10101111).
func TurnOnRightMost0Bit(x int) int {
	return x | (x + 1)
}

// TurnOffTrailing1s turns off the trailing 1’s in a word, producing x if none
// (e.g., 10100111 ⇒ 10100000).
func TurnOffTrailing1s(x int) int {
	return x & (x + 1)
}

// TurnOnTrailing0s turns on the trailing 0’s in a word, producing x if none
// (e.g., 10101000 ⇒ 10101111).
func TurnOnTrailing0s(x int) int {
	return x | (x - 1)
}

// Signle1BitAtRightMost0Bit creates a word with a single 1-bit at the position
// of the rightmost 0-bit in x, producing 0 if none (e.g., 10100111 ⇒ 00001000).
func Single1BitAtRightMost0Bit(x int) int {
	return (x + 1) &^ x
}

// Single0BitAtRightMost1Bit creates a word with a single 0-bit at the
// position of the rightmost 1-bit in x, producing all 1’s if none
// (e.g., 10101000 ⇒ 11110111):
func Single0BitAtRightMost1Bit(x int) int {
	return ^x | (x - 1)
}

// ReplaceTrailing0sWith1sA creates a word with 1’s at the positions of the
// trailing 0’s in x, and 0’s elsewhere, producing 0 if none
// (e.g., 01011000 ⇒ 00000111)
func ReplaceTrailing0sWith1sA(x int) int {
	return ^x & (x - 1)
}

// ReplaceTrailing0sWith1sB creates a word with 1’s at the positions of the
// trailing 0’s in x, and 0’s elsewhere, producing 0 if none
// (e.g., 01011000 ⇒ 00000111)
func ReplaceTrailing0sWith1sB(x int) int {
	return ^(x | -x)
}

// ReplaceTrailing0sWith1sC creates a word with 1’s at the positions of the
// trailing 0’s in x, and 0’s elsewhere, producing 0 if none
// (e.g., 01011000 ⇒ 00000111)
func ReplaceTrailing0sWith1sC(x int) int {
	return (x & -x) - 1
}

// ReplaceTrailing1sWith0s creates a word with 0’s at the positions of the
// trailing 1’s in x, and 1’s elsewhere, producing all 1’s if none
// (e.g., 10100111 ⇒ 11111000):
func ReplaceTrailing1sWith0s(x int) int {
	return ^x | (x + 1)
}

// IsolateRightMost1Bit isolates the rightmost 1-bit, producing 0 if none
// (e.g., 01011000 ⇒ 00001000):
func IsolateRightMost1Bit(x int) int {
	return x & (-x)
}

// ReplaceRightMost1BitAndTrailing0sWith1s creates a word with 1’s at the
// positions of the rightmost 1-bit and the trailing 0’s in x, producing all
// 1’s if no 1-bit, and the integer 1 if no trailing 0’s
// (e.g., 01011000 ⇒ 00001111):
func ReplaceRightMost1BitAndTrailing0sWith1s(x int) int {
	return x ^ (x - 1)
}
