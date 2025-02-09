package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"
)

// EncryptSecret encrypts a given secret using AES-GSM.
func EncryptSecret(secret string, key string) (string, error) {
	// Ensure the key is of valid size
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.New("invalid key size: key must be 16, 24, or 32 bytes")
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal("NewCipher: ", err)
		return "", err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal("ReadFull: ", err)
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal("NewGCM: ", err)
		return "", err
	}

	encrypted := aesGCM.Seal(nonce, nonce, []byte(secret), nil)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// DecryptSecret decrypts a given encrypted secret using AES-Gcm.
func DecryptSecret(encryptedSecret string, key string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encryptedSecret)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("invalid encrypted data")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	decrypted, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}
