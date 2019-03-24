package mode

import (
	"bytes"
	"crypto/aes"
	"testing"
)

func TestCBCMode(t *testing.T) {
	tables := []struct {
		plaintext []byte
		iv        []byte
		key       []byte
	}{
		{[]byte("This is a test. Another test #2."), bytes.Repeat([]byte{0}, 16), []byte("YELLOW SUBMARINE")},
	}
	for _, table := range tables {
		aesCipher, err := aes.NewCipher(table.key)
		if err != nil {
			t.Fatal(err)
		}
		encrypter := NewCBCEncrypter(aesCipher, table.iv)
		decrypter := NewCBCDecrypter(aesCipher, table.iv)

		ciphertext := make([]byte, len(table.plaintext))
		encrypter.CryptBlocks(ciphertext, table.plaintext)

		ptext := make([]byte, len(table.plaintext))
		decrypter.CryptBlocks(ptext, ciphertext)

		if !bytes.Equal(ptext, table.plaintext) {
			t.Errorf("TestCBCMode(%q, %q, %q) = %q; want: %q", table.plaintext, table.iv, table.key, ptext, table.plaintext)
		}
	}
}
