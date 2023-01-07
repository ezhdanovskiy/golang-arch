package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
)

func NewEncryptWriter(w io.Writer, key []byte) (io.Writer, error) {
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create new cipher.Block: %w", err)
	}

	s := cipher.NewCTR(cipherBlock, make([]byte, aes.BlockSize))

	return cipher.StreamWriter{
		S: s,
		W: w,
	}, nil
}
