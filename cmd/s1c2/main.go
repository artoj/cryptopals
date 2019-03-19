package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

func fixedXor(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("lengths do not match, %d != %d", len(a), len(b))
	}
	result := make([]byte, len(a))
	for i, _ := range a {
		result[i] = a[i] ^ b[i]
	}
	return result, nil
}

func main() {
	a, err := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	if err != nil {
		fmt.Println("a: invalid hex string")
		os.Exit(1)
	}
	b, err := hex.DecodeString("686974207468652062756c6c277320657965")
	if err != nil {
		fmt.Println("b: invalid hex string")
		os.Exit(1)
	}
	xorred, err := fixedXor(a, b)
	if err != nil {
		fmt.Println("xor failed:", err)
		os.Exit(1)
	}
	fmt.Printf("%x\n", xorred)
}
