package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
)

func signMsg(msg []byte) (signature []byte, err error) {
	const key = "52e0b4616b25ab5aa1958d957e5b2f09757405416f15e1507368894127d4daba01cda189feb46774ca067eb118791fb98db85e5addfd84c531dde5d597e4fd68"
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
