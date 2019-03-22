package utils

import (
	"fmt"
	"math/bits"

	"github.com/artoj/cryptopals/internal/pkg/pad"
)

// Transpose performs a matrix transpose operation on blocks.
func Transpose(blocks [][]byte) [][]byte {
	result := make([][]byte, len(blocks[0]))
	for i := 0; i < len(blocks); i++ {
		for j := 0; j < len(blocks[i]); j++ {
			result[j] = append(result[j], blocks[i][j])
		}
	}
	return result
}

// Split splits the input c in to blockSize length slices.
// Input is automatically zero padded.
func Split(c []byte, blockSize int) [][]byte {
	var blocks [][]byte

	padded := pad.ZeroPad(c, blockSize)

	for i := 0; i < len(padded); i += blockSize {
		blocks = append(blocks, padded[i:i+blockSize])
	}
	return blocks
}

// IsPrintable returns a boolean value whether b is an ASCII printable character or not.
func IsPrintable(b byte) bool {
	if b >= 0x20 && b <= 0x7e {
		return true
	}
	return false
}

// RepeatedKeyXor performs a XOR operation using a repeated key k
func RepeatedKeyXor(k, p []byte) []byte {
	c := make([]byte, len(p))

	for i := range p {
		c[i] = p[i] ^ k[i%len(k)]
	}
	return c
}

// HammingDistance calculates the hamming distance between two byte strings
func HammingDistance(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("lengths do not match, %d != %d", len(a), len(b))
	}

	dist := 0

	for i := range a {
		dist += bits.OnesCount8(uint8(a[i]) ^ uint8(b[i]))
	}
	return dist, nil
}
