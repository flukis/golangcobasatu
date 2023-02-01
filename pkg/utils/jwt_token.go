package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Tokens struct {
	Access  string
	Refresh string
}

func GenerateNewTokens(id string) (*Tokens, error) {
	// Generate JWT Access token.
	accessToken, err := generateNewAccessToken(id)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	// Generate JWT Refresh token.
	refreshToken, err := generateNewRefreshToken(id)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	return &Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func generateNewAccessToken(id string) (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")

	minuteCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	claims := jwt.MapClaims{}

	claims["id"] = id
	claims["expires"] = time.Now().Add(time.Minute * time.Duration(minuteCount))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// generate tokens
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func generateNewRefreshToken(id string) (string, error) {
	hash := sha256.New()

	refresh := os.Getenv("JWT_REFRESH_KEY")

	_, err := hash.Write([]byte(refresh))
	if err != nil {
		return "", err
	}

	hoursCount, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT"))

	expiredTime := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(hoursCount)))

	t := hex.EncodeToString(hash.Sum(nil)) + "." + expiredTime
	return t, nil
}

func ParseRefreshToken(refreshToken string) (int64, error) {
	return strconv.ParseInt(strings.Split(refreshToken, ".")[1], 0, 64)
}
