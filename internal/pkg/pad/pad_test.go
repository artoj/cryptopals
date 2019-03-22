package pad

import (
	"bytes"
	"testing"
)

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
		if !bytes.Equal(padded, table.expected) {
			t.Errorf("ZeroPad(%q, %d) = %q; want: %q", table.input, table.blocksize, padded, table.expected)
		}
	}
}

func TestPKCS7Pad(t *testing.T) {
	tables := []struct {
		input     []byte
		blocksize int
		expected  []byte
	}{
		{[]byte("This is a test ab"), 8, []byte("This is a test ab\x07\x07\x07\x07\x07\x07\x07")},
		{[]byte("This is a test 1"), 8, []byte("This is a test 1")},
		{[]byte("This is a test"), 8, []byte("This is a test\x02\x02")},
		{[]byte("abcdefgh"), 8, []byte("abcdefgh")},
		{[]byte("abc"), 8, []byte("abc\x05\x05\x05\x05\x05")},
	}
	for _, table := range tables {
		padded := PKCS7Pad(table.input, table.blocksize)
		if !bytes.Equal(padded, table.expected) {
			t.Errorf("PKCS7Pad(%q, %d) = %q; want: %q", table.input, table.blocksize, padded, table.expected)
		}
	}
}
