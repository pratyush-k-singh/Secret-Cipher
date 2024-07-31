package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"math/big"
	"strings"
	"sync"

	"cipher_interpreter/storage"
)

var cipherLock sync.Mutex

func GenerateCipher(message string, complexity int) (string, string) {
	key := make([]byte, complexity)
	for i := 0; i < complexity; i++ {
		randomChar, _ := rand.Int(rand.Reader, big.NewInt(26))
		key[i] = byte('a' + randomChar.Int64())
	}

	var cipherText strings.Builder
	for i, char := range message {
		cipherChar := (char + rune(key[i%complexity])) % 256
		cipherText.WriteRune(cipherChar)
	}

	encryptedKey, _ := encrypt(key)
	return cipherText.String(), encryptedKey
}

func Decipher(cipherText string) string {
	cipherLock.Lock()
	defer cipherLock.Unlock()

	for encryptedKey := range storage.GetCipherKeys() {
		key, err := decrypt(encryptedKey)
		if err != nil {
			continue
		}
		var originalText strings.Builder
		complexity := len(key)
		for i, char := range cipherText {
			originalChar := (char - rune(key[i%complexity]) + 256) % 256
			originalText.WriteRune(originalChar)
		}
		return originalText.String()
	}
	return ""
}

func encrypt(data []byte) (string, error) {
	block, err := aes.NewCipher(getHashKey())
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(data string) ([]byte, error) {
	block, err := aes.NewCipher(getHashKey())
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func getHashKey() []byte {
	key := "random-key"
	hasher := sha256.New()
	hasher.Write([]byte(key))
	return hasher.Sum(nil)
}
