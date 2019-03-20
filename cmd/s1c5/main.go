package main

import (
	"fmt"

	"github.com/artoj/cryptopals/internal/pkg/utils"
)

const (
	key = "ICE"
	p   = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
)

func main() {
	fmt.Printf("%x\n", utils.RepeatedKeyXor([]byte(key), []byte(p)))
}
