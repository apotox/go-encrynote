package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func GenerateAesKey() (string, error) {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}

func Encrypt(key string, plaintext string) (string, error) {

	keyBytes := []byte(key)
	c, err := aes.NewCipher(keyBytes)

	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	encrypted := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	return hex.EncodeToString(encrypted), nil
}

func Decrypt(key string, encrypted string) (string, error) {

	ciphertext, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	keyBytes := []byte(key)
	c, err := aes.NewCipher(keyBytes)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()

	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("encrypted text too short")
	}

	// split the nonce and the ciphertext
	nonce, cipher := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), cipher, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
