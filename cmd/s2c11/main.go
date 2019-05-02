package main

import (
	"bytes"
	"crypto/aes"
	crand "crypto/rand"
	"log"
	"math/rand"
	"time"

	"github.com/artoj/cryptopals/internal/pkg/mode"
	"github.com/artoj/cryptopals/internal/pkg/pad"
	"github.com/artoj/cryptopals/internal/pkg/utils"
)

type CipherMode int

const (
	ECB CipherMode = iota
	CBC
)

func randRead(c []byte) {
	_, err := crand.Read(c)
	if err != nil {
		log.Fatal(err)
	}
}

func encryptionOracle(input []byte) ([]byte, CipherMode) {
	// create a random AES key
	key := make([]byte, 16)
	randRead(key)

	rand.Seed(time.Now().UnixNano())

	// append bytes before and after plaintext
	// append 'A' bytes
	prefix := bytes.Repeat([]byte("A"), rand.Intn(6)+5)
	postfix := bytes.Repeat([]byte("A"), rand.Intn(6)+5)
	log.Printf("prepending %d bytes, appending %d bytes\n", len(prefix), len(postfix))

	totalLen := len(input) + len(prefix) + len(postfix)
	plainText := make([]byte, totalLen)

	copy(plainText[0:len(prefix)], prefix)
	copy(plainText[len(prefix):len(input)+len(prefix)], input)
	copy(plainText[len(prefix)+len(input):len(prefix)+len(input)+len(postfix)], postfix)

	paddedPlainText := pad.PKCS7Pad(plainText, len(key))
	log.Printf("padded plain text = %q\n", paddedPlainText)
	cipherText := make([]byte, len(paddedPlainText))

	aesCipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	var selectedMode CipherMode
	if rand.Int()%2 == 0 {
		iv := make([]byte, 16)
		randRead(iv)

		aesCBC := mode.NewCBCEncrypter(aesCipher, iv)
		aesCBC.CryptBlocks(cipherText, paddedPlainText)
		selectedMode = CBC
	} else {
		aesECB := mode.NewECBEncrypter(aesCipher)
		aesECB.CryptBlocks(cipherText, paddedPlainText)
		selectedMode = ECB
	}
	return cipherText, selectedMode
}

// countDuplicates returns the number of duplicate entries in s.
func countDuplicates(s []byte) int {
	bytes := make(map[byte]int)
	for _, b := range s {
		if _, ok := bytes[b]; ok != true {
			bytes[b] = 1
		} else {
			bytes[b]++
		}
	}
	num := 0
	for _, v := range bytes {
		if v > 1 {
			num += v - 1
		}
	}
	return num
}

func main() {
	cipherText, mode := encryptionOracle(bytes.Repeat([]byte("B"), 128))
	log.Printf("cipher text = %v, len = %d\n", cipherText, len(cipherText))

	// Use similar detection mechanism as in set s1c8:
	// 1. split the cipher tex to KEYSIZE blocks
	// 2. transpose the blocks
	// 3. calculate the number of duplicates

	blocks := utils.Split(cipherText, 16)
	transposed := utils.Transpose(blocks)

	totalNumDuplicates := 0
	for _, t := range transposed {
		totalNumDuplicates += countDuplicates(t)
	}

	var guess CipherMode
	if totalNumDuplicates > 20 {
		log.Print("guess: ECB")
		guess = ECB
	} else {
		log.Print("guess: CBC")
		guess = CBC
	}

	if guess == mode {
		log.Print("Correct!")
	} else {
		log.Print("Incorrect!")
	}
}
