package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/argon2"
	"log"
	"strings"
)

func GenerateSalt(size int) ([]byte, error) {
	salt := make([]byte, size)

	_, err := rand.Read(salt)

	if err != nil {
		return nil, err
	}

	return salt, nil
}

func HashCreate(input string) (string, error) {
	salt, err := GenerateSalt(16)
	if err != nil {
		log.Fatal(err)
	}

	hash := argon2.IDKey([]byte(input), salt, 1, 64*1024, 4, 32)
	hashWithSalt := append(salt, hash...)
	return base64.RawStdEncoding.EncodeToString(hashWithSalt), nil
}

func HashValidate(input string, encodedHash string) (bool, error) {
	hashWithSalt, err := base64.RawStdEncoding.DecodeString(encodedHash)
	if err != nil {
		return false, err
	}

	salt := hashWithSalt[:16]
	hash := hashWithSalt[16:]

	inputHash := argon2.IDKey([]byte(input), salt, 1, 64*1024, 4, 32)
	return strings.Compare(string(hash), string(inputHash)) == 0, nil
}
