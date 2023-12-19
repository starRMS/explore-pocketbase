package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"

	"github.com/starRMS/explore-pocketbase/tools/utils"
)

var (
	// TODO: read from environment
	key string = "my32digitkey12345678901234567890"
	iv  string = "my16digitIvKey12"
)

func AES_CBC_Encrypt(plaintext any) (string, error) {
	block, err := aes.NewCipher(utils.ConvertStringToSliceByte(key))
	if err != nil {
		return "", err
	}

	var encrypted []byte

	switch plaintext := plaintext.(type) { // Eliminate type assertion
	case string:
		padded := pad(utils.ConvertStringToSliceByte(plaintext))
		encrypted = make([]byte, len(padded))
		mode := cipher.NewCBCEncrypter(block, utils.ConvertStringToSliceByte(iv))
		mode.CryptBlocks(encrypted, padded)
		return base64.StdEncoding.EncodeToString(encrypted), nil
		// TODO: More type cases
	default:
		return "", nil
	}
}

func AES_CBC_Decrypt(decoded string) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(decoded)
	if err != nil {
		return nil, errors.New("decode ciphertext failed")
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("invalid ciphertext content, not a valid aes encryption")
	}

	block, err := aes.NewCipher(utils.ConvertStringToSliceByte(key))
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, utils.ConvertStringToSliceByte(iv))
	mode.CryptBlocks(ciphertext, ciphertext)
	plaintext, err := unpad(ciphertext)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
