package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/artoj/cryptopals/internal/pkg/utilities"
)

const InFile = "4.txt"

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
		if !utilities.IsPrintable(b) {
			return -100
		} else {
			return 0
		}
	}
	return w
}

func totalWeight(b []byte) int {
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
	in, err := ioutil.ReadFile(InFile)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range bytes.Split(in, []byte("\n")) {
		decoded := make([]byte, hex.DecodedLen(len(s)))
		_, err = hex.Decode(decoded, s)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%x\n", decoded)
		for k := byte(0x20); k < 0x7f; k++ {
			xorred := xor(k, decoded)
			w := totalWeight(xorred)
			if w > 0 {
				fmt.Printf("\tk: '%c'\tw: %d\tvalue: %s\n", k, w, xorred)
			}
		}
	}
}
