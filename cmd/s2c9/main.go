package main

import (
	"bytes"
	"fmt"

	"github.com/artoj/cryptopals/internal/pkg/pad"
)

func main() {
	if bytes.Equal([]byte("YELLOW SUBMARINE\x04\x04\x04\x04"), pad.PKCS7Pad([]byte("YELLOW SUBMARINE"), 20)) {
		fmt.Println("ok")
	} else {
		fmt.Println("PKCS#7 padding failure")
	}
}
