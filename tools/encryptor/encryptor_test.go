package encryptor

import (
	"fmt"
	"testing"
)

func TestEncryptionAndDecryption(t *testing.T) {
	tests := []struct {
		plaintext string
	}{
		{
			plaintext: "plaintext_1_test_case",
		},
		{
			plaintext: "plaintext_2_test_case",
		},
		{
			plaintext: "plaintext_3_test_case",
		},
		{
			plaintext: "pla",
		},
		{
			plaintext: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris finibus arcu at ipsum eleifend orci.",
		},
	}

	for num, test := range tests {
		t.Run(fmt.Sprintf("Run test case %d", num+1), func(t *testing.T) {
			ciphertext, err := AES_CBC_Encrypt(test.plaintext)
			if err != nil {
				t.Fatalf("error was found while encrypting %s\n", err)
			}

			plaintext, err := AES_CBC_Decrypt(ciphertext)
			if err != nil {
				t.Fatalf("error was found while decrypting %s\n", err)
			}

			if string(plaintext) != test.plaintext {
				t.Fatalf("mismatched plaintext result")
			}
		})
	}
}
