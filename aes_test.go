package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestAES(t *testing.T) {
	const (
		msg      = "This is totally fun get hands-on and learning it from the ground up."
		password = "ilovedogs"
	)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	require.NoError(t, err)
	fmt.Printf("hashedPassword: %v\n", string(hashedPassword))

	require.NoError(t, bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)))

	hashedPassword = hashedPassword[:16]
	//hashedPassword = []byte("$2a$04$d85Ym4ikX")
	fmt.Printf("hashedPassword: %v\n", string(hashedPassword))

	var encryptedMsg string
	wtr := &bytes.Buffer{}
	{
		encryptWriter, err := NewEncryptWriter(wtr, hashedPassword)
		require.NoError(t, err)

		_, err = io.WriteString(encryptWriter, msg)
		require.NoError(t, err)

		fmt.Printf("encryptedMsg: %v\n", base64.StdEncoding.EncodeToString(wtr.Bytes()))
		encryptedMsg = wtr.String()
	}

	{
		wtr.Reset()
		encryptWriter, err := NewEncryptWriter(wtr, hashedPassword)
		require.NoError(t, err)

		_, err = io.WriteString(encryptWriter, encryptedMsg)
		require.NoError(t, err)

		fmt.Printf("decryptedMsg: %v\n", wtr.String())
		assert.Equal(t, msg, wtr.String())
	}
}
