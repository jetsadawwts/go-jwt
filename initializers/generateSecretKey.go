package initializers

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateSecretKey(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	secretKey := base64.StdEncoding.EncodeToString(randomBytes)
	return secretKey, nil
}
