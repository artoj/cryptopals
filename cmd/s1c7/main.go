package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/artoj/cryptopals/internal/pkg/ecb"
)

const (
	inFile = "7.txt"
	key    = "YELLOW SUBMARINE"
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

	aesECB := ecb.NewECBDecrypter(aesCipher)
	aesECB.CryptBlocks(c, c)
	fmt.Println(string(c))
}
