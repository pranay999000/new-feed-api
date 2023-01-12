package config

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func ValidateJWT(token string) (jwt.MapClaims, bool) {
	tk, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return SECRET, nil
	})

	if err != nil {
		return nil, false
	}

	if !tk.Valid {
		return nil, false
	}

	claims, ok := tk.Claims.(jwt.MapClaims)
	if !ok {
		return nil, false
	}
	return claims, true
}