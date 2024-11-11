package jwt

import (
	"aws/models"
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUser string

func ProcessToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	myKey := []byte(JWTSign)
	var claims models.Claim
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, "", errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		//Check againts db
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("token invalido")
	}

	return &claims, false, string(""), nil
}
