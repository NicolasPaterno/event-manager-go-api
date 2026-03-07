package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const secretkey = "secretKey"

func GenerateToken(email string, userid int64) (string, error) {
	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      email,
		"userId":     userid,
		"expiration": time.Now().Add(time.Hour * 2).Unix(),
	})
	return tokenJwt.SignedString([]byte(secretkey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretkey), nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}
	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("invalid Token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalide jwt claims")
	}

	userid := int64(claims["userId"].(float64))
	// email := claims["email"].(string)
	return userid, nil
}
