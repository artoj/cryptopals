package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/artoj/cryptopals/internal/pkg/mode"
)

const (
	inFile = "10.txt"
	key    = "YELLOW SUBMARINE"
	iv     = "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
)

func main() {
	in, err := ioutil.ReadFile(inFile)
	if err != nil {
		log.Fatal(err)
	}

	c, err := base64.StdEncoding.DecodeString(string(in))
	if err != nil {
		log.Fatal(err)
	}

	aesCipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	aesCBC := mode.NewCBCDecrypter(aesCipher, []byte(iv))
	plaintext := make([]byte, len(c))
	aesCBC.CryptBlocks(plaintext, c)
	fmt.Println(string(plaintext))
}
