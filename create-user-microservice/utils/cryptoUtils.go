package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
)

// Encryption key (must be 16, 24, or 32 bytes)
var encryptionKey = []byte("your_encryption_key_32bytes")

// Encrypt data using AES
func Encrypt(data string) string {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		log.Fatal(err)
	}

	nonce := []byte("123456789012") // Must be unique for each encryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}

	ciphertext := aesGCM.Seal(nil, nonce, []byte(data), nil)
	return base64.StdEncoding.EncodeToString(ciphertext)
}
