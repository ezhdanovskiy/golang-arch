package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
)

func signMsg(msg []byte) (signature []byte, err error) {
	const key = "key"
	h := hmac.New(sha512.New, []byte(key))

	_, err = h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to write msg: %w", err)
	}

	return h.Sum(nil), nil
}

func checkSign(msg, sign []byte) (bool, error) {
	newSign, err := signMsg(msg)
	if err != nil {
		return false, fmt.Errorf("failed to sign msg: %w", err)
	}

	return hmac.Equal(sign, newSign), nil
}
