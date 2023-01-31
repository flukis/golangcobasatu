package utils

import "golang.org/x/crypto/bcrypt"

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
