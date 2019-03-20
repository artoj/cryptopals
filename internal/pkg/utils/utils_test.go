package utils

import "testing"

func isEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func isEqualMulti(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !isEqual(b[i], v) {
			return false
		}
	}
	return true
}

func TestTranspose(t *testing.T) {
	tables := []struct {
		input    [][]byte
		expected [][]byte
	}{
		{[][]byte{[]byte{0xff, 0x00, 0xff}, []byte{0x00, 0xff, 0x00}}, [][]byte{[]byte{0xff, 0x00}, []byte{0x00, 0xff}, []byte{0xff, 0x00}}},
	}
	for _, table := range tables {
		transposed := Transpose(table.input)
		if !isEqualMulti(transposed, table.expected) {
			t.Errorf("Transpose of %v was incorrect, got: %v, want: %v", table.input, transposed, table.expected)
		}
	}
}

func TestZeroPad(t *testing.T) {
	tables := []struct {
		input     []byte
		blocksize int
		expected  []byte
	}{
		{[]byte("This is a test ab"), 8, []byte("This is a test ab\x00\x00\x00\x00\x00\x00\x00")},
		{[]byte("This is a test 1"), 8, []byte("This is a test 1")},
		{[]byte("This is a test"), 8, []byte("This is a test\x00\x00")},
		{[]byte("abcdefgh"), 8, []byte("abcdefgh")},
		{[]byte("abc"), 8, []byte("abc\x00\x00\x00\x00\x00")},
	}
	for _, table := range tables {
		padded := ZeroPad(table.input, table.blocksize)
		if !isEqual(padded, table.expected) {
			t.Errorf("ZeroPad of (\"%s\", %d) was incorrect, got: %q, want: %q", table.input, table.blocksize, padded, table.expected)
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
