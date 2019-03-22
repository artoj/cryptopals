package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/artoj/cryptopals/internal/pkg/utils"
)

const (
	inFile    = "8.txt"
	blockSize = 16
)

// inputNumDuplicate is a input - number of duplicates pair
type inputNumDuplicates struct {
	Input         string
	NumDuplicates int
}

type byDuplicates []inputNumDuplicates

func (d byDuplicates) Len() int           { return len(d) }
func (d byDuplicates) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d byDuplicates) Less(i, j int) bool { return d[i].NumDuplicates < d[j].NumDuplicates }

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
	f, err := os.Open(inFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Idea for detection: since the operation of ECB is deterministric
	// same values exists for the same key - plaintext pairs.
	//
	// 1. Split the input to KEYSIZE blocks
	// 2. Transpose the input to get the same position elements to the
	//    same slice
	// 3. Count the total number of duplicate elements in the slice.
	// 4. Sort the list of input - number of duplicate pairs. The input
	//    with the largest number of duplicates is the input encrypted
	//    with ECB.
	scanner := bufio.NewScanner(f)
	var results []inputNumDuplicates
	for scanner.Scan() {
		c, err := base64.StdEncoding.DecodeString(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		blocks := utils.Split(c, blockSize)
		transposed := utils.Transpose(blocks)

		totalNumDuplicates := 0
		for _, t := range transposed {
			totalNumDuplicates += countDuplicates(t)
		}
		results = append(results, inputNumDuplicates{scanner.Text(), totalNumDuplicates})
	}
	sort.Sort(sort.Reverse(byDuplicates(results)))
	for _, r := range results[:5] {
		fmt.Printf("%s...: %d\n", r.Input[:16], r.NumDuplicates)
	}
}
