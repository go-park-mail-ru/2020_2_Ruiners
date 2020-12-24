package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (string, error) {
	salt := make([]byte, 8)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)
	hash = append(salt, hash...)
	return hex.EncodeToString(hash[:]), nil
}

func CheckPassword(password string, hash string) (bool, error) {
	hashByte, err := hex.DecodeString(hash)
	if err != nil {
		return false, err
	}

	salt := hashByte[:8]
	passwordToHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	passwordToHash = append(salt, passwordToHash...)
	return hex.EncodeToString(passwordToHash[:]) == hash, nil
}
