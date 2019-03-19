package utils

import "testing"

func TestZeroPadd(t *testing.T) {
	tables := []struct {
		input     []byte
		blocksize int
		paddedLen int
	}{
		{[]byte("This is a test ab"), 8, 24},
		{[]byte("This is a test 1"), 8, 16},
		{[]byte("This is a test"), 8, 16},
		{[]byte("abcdefgh"), 8, 8},
		{[]byte("abc"), 8, 8},
	}
	for _, table := range tables {
		padded := ZeroPad(table.input, table.blocksize)
		if len(padded) != table.paddedLen {
			t.Errorf("ZeroPad of (\"%s\", %d) was incorrect, got: %d, want: %d", table.input, table.blocksize, len(padded), table.paddedLen)
		}
	}
}

func TestHammingDistance(t *testing.T) {
	dist, err := HammingDistance([]byte("this is a test"), []byte("wokka wokka!!!"))
	if err != nil {
		t.Error("Hamming distance calculation failed")
	}
	// the value "37" was given as correct distance in set #1 exercise #6
	if dist != 37 {
		t.Errorf("Incorrect hamming distance, got: %d, want: %d", dist, 37)
	}
}
