package main

import (
	"fmt"

	"github.com/artoj/cryptopals/internal/pkg/utilities"
)

const (
	key = "ICE"
	p   = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
)

func main() {
	fmt.Printf("%x\n", utilities.RepeatedKeyXor([]byte(key), []byte(p)))
}
