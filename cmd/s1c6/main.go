package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"github.com/artoj/cryptopals/internal/pkg/utilities"
)

const (
	inFile     = "6.txt"
	maxKeySize = 40
)

// KeySizeDistance is a key size - hamming distance pair
type KeySizeDistance struct {
	KeySize, Distance int
}

type byHammingDistance []KeySizeDistance

func (k byHammingDistance) Len() int           { return len(k) }
func (k byHammingDistance) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
func (k byHammingDistance) Less(i, j int) bool { return k[i].Distance < k[j].Distance }

// KeyWeight is a byte key - word frequency weight pair
type KeyWeight struct {
	Key    byte
	Weight int
}

type byWeight []KeyWeight

func (k byWeight) Len() int           { return len(k) }
func (k byWeight) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
func (k byWeight) Less(i, j int) bool { return k[i].Weight < k[j].Weight }

func calcDistances(c []byte) []KeySizeDistance {
	var distances []KeySizeDistance

	for i := 2; i <= maxKeySize; i++ {

		blocks := [][]byte{c[0:i], c[i : i*2], c[i*2 : i*3], c[i*3 : i*4]}

		var totalDist int
		for i := 0; i < len(blocks)-1; i++ {
			dist, err := utilities.HammingDistance(blocks[i], blocks[i+1])
			if err != nil {
				log.Fatal(err)
			}
			totalDist += dist
		}
		distances = append(distances, KeySizeDistance{i, totalDist / i})
	}

	return distances
}

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
		}
		return -1
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

// Solves a single key XOR encryption in c
func solveBlock(c []byte) byte {
	var keyWeights []KeyWeight
	for k := byte(0x20); k < 0x7f; k++ {
		xorred := utilities.RepeatedKeyXor([]byte{k}, c)
		weight := totalWeight(xorred)

		keyWeights = append(keyWeights, KeyWeight{k, weight})
	}

	// Largest weights first
	sort.Sort(sort.Reverse(byWeight(keyWeights)))
	return keyWeights[0].Key
}

func main() {
	in, err := ioutil.ReadFile(inFile)
	if err != nil {
		log.Fatal(err)
	}

	c, err := base64.StdEncoding.DecodeString(string(in))
	if err != nil {
		log.Fatal(err)
	}

	distances := calcDistances(c)
	sort.Sort(byHammingDistance(distances))
	for _, keySizeDistance := range distances[:5] {
		var key []byte
		split := utilities.Split(c, keySizeDistance.KeySize)
		transposed := utilities.Transpose(split)

		for _, block := range transposed {
			key = append(key, solveBlock(block))
		}

		fmt.Printf("key: '%s'\tplain: %q\n", key, utilities.RepeatedKeyXor(key, c)[0:100])
	}
}
