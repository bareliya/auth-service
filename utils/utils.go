package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func GetHashFromPassword(userPassword string) string {
	data := []byte(userPassword)
	hash := sha512.Sum512(data)
	hashString := hex.EncodeToString(hash[:])
	return hashString
}
func CheckPassword(userPassword, dbPasswordHash string) bool {
	userPassHash := GetHashFromPassword(userPassword)
	if userPassHash == dbPasswordHash {
		return true
	}
	return false
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"

func GenerateRandomPassword(length int) string {
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}
