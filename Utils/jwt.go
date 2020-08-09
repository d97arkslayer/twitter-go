package Utils

import (
	"github.com/d97arkslayer/twitter-go/Models"
	 "github.com/dgrijalva/jwt-go"
	"time"
)

/**
 * GenerateJWT
 * Generate de JSON WEB TOKEN
 */
func GenerateJWT(u Models.User)(string, error)  {
	secret := []byte("VerificarComoUsarDotEnvsEnGO")
	payload := jwt.MapClaims{
		"email": u.Email,
		"name": u.Name,
		"lastname": u.Lastname,
		"birth_date": u.BirthDate,
		"biography": u.Biography,
		"location": u.Location,
		"website": u.Website,
		"_id": u.Id.Hex(),
		"exp": time.Now().Add(time.Hour*24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}