package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPasword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPaswordHash(password string, hash string) (bool, error) {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err != nil {
        return false, err
    }
    return true, err
}