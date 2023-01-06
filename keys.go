package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"

	"github.com/gofrs/uuid"
)

type Key struct {
	key       []byte
	createdAt time.Time
}

type Keys struct {
	keys       map[string]*Key
	currentKid string
}

func NewKeys() (*Keys, error) {
	k := &Keys{
		keys: make(map[string]*Key),
	}
	err := k.GenerateNewKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generateNewKey: %w", err)
	}

	return k, nil
}

func (k *Keys) GenerateNewKey() error {
	newKey := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, newKey)
	if err != nil {
		return fmt.Errorf("failed to generate new key")
	}

	uid, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("failed to generate uuid")
	}

	k.keys[uid.String()] = &Key{
		key:       newKey,
		createdAt: time.Now(),
	}
	k.currentKid = uid.String()

	return nil
}

func (k *Keys) Key(kid string) *Key {
	return k.keys[kid]
}

func (k *Keys) CurrentKey() *Key {
	return k.keys[k.currentKid]
}

func (k *Keys) CurrentKid() string {
	return k.currentKid
}
