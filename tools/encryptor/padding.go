package encryptor

import (
	"bytes"
	"crypto/aes"
	"errors"
)

func pad(v []byte) []byte {
	length := aes.BlockSize - len(v)%aes.BlockSize
	pad := bytes.Repeat([]byte{uint8(length)}, length)
	return append(v, pad...)
}

func unpad(v []byte) ([]byte, error) {
	if len(v) == 0 {
		return nil, errors.New("unpad called with 0 length data")
	}

	unpadding := int(v[len(v)-1])
	if unpadding > len(v) {
		return nil, errors.New("error decrypting data")
	}

	return v[:(len(v) - unpadding)], nil
}
