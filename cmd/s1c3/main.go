package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/artoj/cryptopals/internal/pkg/utils"
)

// relative letter frequencies in the English alphabet with some made-up weights
var weights = map[byte]int{
	'e': 12,
	't': 9,
	'a': 8,
	'o': 7,
	'i': 7,
	'n': 7,
	's': 6,
	'h': 6,
	'r': 6,
	'd': 4,
	'l': 4,
	'c': 3,
	'u': 3,
	'm': 2,
	'w': 2,
	'f': 2,
	'g': 2,
	'y': 2,
	'p': 2,
	'b': 2,
}

func weight(b byte) int {
	w, ok := weights[b]
	if !ok {
		if !utils.IsPrintable(b) {
			return -100
		} else {
			return 0
		}
	}
	return w
}

func total_weight(b []byte) int {
	result := 0
	for _, w := range b {
		result += weight(w)
	}
	return result
}

func xor(k byte, c []byte) []byte {
	result := make([]byte, len(c))
	for i := range c {
		result[i] = c[i] ^ k
	}
	return result
}

func main() {
	c, err := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		fmt.Println("c: invalid hex string")
		os.Exit(1)
	}
	in_weights := make(map[byte]int)
	for k := byte(0x20); k < 0x7f; k++ {
		in_weights[k] = total_weight(xor(k, c))
	}

	for k, w := range in_weights {
		if w > 0 {
			fmt.Printf("k: '%c'\tw: %d\tvalue: %s\n", k, w, xor(k, c))
		}
	}
}
