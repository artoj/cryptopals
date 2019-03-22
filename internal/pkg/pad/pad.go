package pad

// paddedLen calculates the total length of the ciphertext when padded to blockSize.
func paddedLen(ptlen, blockSize int) int {
	if ptlen <= blockSize {
		return blockSize
	}
	size := ptlen / blockSize * blockSize
	if ptlen%blockSize != 0 {
		size += blockSize
	}
	return size
}

// ZeroPad pads the input b with zeroes to a given blockSize.
func ZeroPad(b []byte, blockSize int) []byte {
	padded := make([]byte, paddedLen(len(b), blockSize))
	copy(padded, b)
	return padded
}

// PKCS7Pad implements PKCS#7 padding of b to blockSize.
func PKCS7Pad(b []byte, blockSize int) []byte {
	plen := paddedLen(len(b), blockSize)
	padded := make([]byte, plen)
	copy(padded, b)
	for i := len(b); i < plen; i++ {
		padded[i] = byte(plen - len(b))
	}
	return padded
}
