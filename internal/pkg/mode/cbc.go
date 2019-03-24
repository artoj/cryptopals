// Cipher block chaining (CBC) mode.

package mode

import (
	"crypto/cipher"
)

func xor(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("xor: lenghts do not match")
	}
	result := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = a[i] ^ b[i]
	}
	return result
}

type cbcDecrypter struct {
	b  cipher.Block
	iv []byte
}

// NewCBCDecrypter returns a BlockMode which decrypts in
// cipher block chaining mode.
func NewCBCDecrypter(block cipher.Block, iv []byte) cipher.BlockMode {
	return &cbcDecrypter{
		b:  block,
		iv: iv,
	}
}

func (x *cbcDecrypter) BlockSize() int { return x.b.BlockSize() }

// CryptBlocks decrypts the a number of blocks.
// The inputs must not overlap.
func (x *cbcDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.BlockSize() != 0 {
		panic("cbc: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("cbc: output smaller than input")
	}

	// XXX: Support overlapping inputs

	var tmp, previousBlock []byte
	ivDone := false
	for len(src) > 0 {
		x.b.Decrypt(dst[:x.BlockSize()], src[:x.BlockSize()])
		if !ivDone {
			tmp = xor(x.iv, dst[:x.BlockSize()])
			ivDone = true
		} else {
			tmp = xor(previousBlock, dst[:x.BlockSize()])
		}
		copy(dst, tmp)
		previousBlock = src[:x.BlockSize()]

		src = src[x.BlockSize():]
		dst = dst[x.BlockSize():]

	}
}

type cbcEncrypter struct {
	b  cipher.Block
	iv []byte
}

// NewCBCEncrypter returns a BlockMode which encrypts in
// cipher block chaining mode.
func NewCBCEncrypter(block cipher.Block, iv []byte) cipher.BlockMode {
	return &cbcEncrypter{
		b:  block,
		iv: iv,
	}
}

func (x *cbcEncrypter) BlockSize() int { return x.b.BlockSize() }

// CryptBlocks decrypts the a number of blocks.
// The inputs must not overlap.
func (x *cbcEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.BlockSize() != 0 {
		panic("cbc: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("cbc: output smaller than input")
	}

	var tmp, previousBlock []byte
	ivDone := false
	for len(src) > 0 {
		if !ivDone {
			tmp = xor(x.iv, src[:x.BlockSize()])
			ivDone = true
		} else {
			tmp = xor(previousBlock, src[:x.BlockSize()])
		}

		x.b.Encrypt(dst[:x.BlockSize()], tmp)

		previousBlock = dst[:x.BlockSize()]
		src = src[x.BlockSize():]
		dst = dst[x.BlockSize():]
	}
}
