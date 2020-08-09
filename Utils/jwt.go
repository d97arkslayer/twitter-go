package Utils

import (
	"errors"
	"github.com/d97arkslayer/twitter-go/Models"
	"github.com/d97arkslayer/twitter-go/Types"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

var secret = []byte("VerificarComoUsarDotEnvsEnGO")

/**
 * GenerateJWT
 * Generate de JSON WEB TOKEN
 */
func GenerateJWT(u Models.User)(string, error)  {
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

/**
 * ProcessToken
 * This functions decode JWT
 */
func ProcessToken(token string)(*Types.Claim, error){
	claims := &Types.Claim{}
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken)!= 2 {
		return claims, errors.New("invalid jwt format")
	}
	token = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(token, claims, func(tk *jwt.Token)(interface{}, error){
		return secret, nil
	})
	if err != nil {
		return claims, err
	}
	if !tkn.Valid {
		return claims, errors.New("invalid token")
	}
	return claims, nil
}