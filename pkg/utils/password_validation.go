package utils

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func NormalizePassword(p string) []byte {
	return []byte(p)
}

func GenerateHashedPassword(p string) (string, error) {
	bytePwd := NormalizePassword(p)

	// generate hashed password
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	// return
	return string(hash), nil
}

func ComparePassword(hashedPassword, inputedPassword string) bool {
	hashed := NormalizePassword(hashedPassword)
	input := NormalizePassword(inputedPassword)

	// return value
	if err := bcrypt.CompareHashAndPassword(hashed, input); err != nil {
		return false
	}
	return true
}

func SimplePasswordValidation(password string) bool {
	regex_, err := regexp.Compile(`^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*|[a-zA-Z0-9]*)$`)
	if err != nil {
		return false
	}

	return regex_.MatchString(password)
}
