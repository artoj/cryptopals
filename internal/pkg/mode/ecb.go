// Electronic code book (ECB) mode.

package mode

import (
	"crypto/cipher"
)

type ecbDecrypter struct {
	b cipher.Block
}

// NewECBDecrypter returns a BlockMode which decrypts in
// electronic code book mode.
func NewECBDecrypter(block cipher.Block) cipher.BlockMode {
	return &ecbDecrypter{
		b: block,
	}
}

func (x *ecbDecrypter) BlockSize() int { return x.b.BlockSize() }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.BlockSize() != 0 {
		panic("ecb: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("ecb: output smaller than input")
	}

	// XXX: overlap not checked

	for len(src) > 0 {
		x.b.Decrypt(dst[:x.BlockSize()], src[:x.BlockSize()])
		src = src[x.BlockSize():]
		dst = dst[x.BlockSize():]
	}
}

type ecbEncrypter struct {
	b cipher.Block
}

// NewECBEncrypter returns a BlockMode which encrypts in
// electronic code book mode.
func NewECBEncrypter(block cipher.Block) cipher.BlockMode {
	return &ecbEncrypter{
		b: block,
	}
}

func (x *ecbEncrypter) BlockSize() int { return x.b.BlockSize() }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.BlockSize() != 0 {
		panic("ecb: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("ecb: output smaller than input")
	}

	// XXX: overlap not checked

	for len(src) > 0 {
		x.b.Encrypt(dst[:x.BlockSize()], src[:x.BlockSize()])
		src = src[x.BlockSize():]
		dst = dst[x.BlockSize():]
	}
}
