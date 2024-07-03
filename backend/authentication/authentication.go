package authentication

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY = "secret"

func GenerateJWT(id string) (string, error) { 
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
  	claims["user"] = id
  	claims["exp"] = time.Now().Add(time.Minute * 5).Unix() 
  	tokenString, err := token.SignedString([]byte(SECRET_KEY))

    return tokenString, err
}

func ValidateJWT(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return "", err
	}
	expiryTime := token.Claims.(jwt.MapClaims)["exp"].(float64)
	if float64(time.Now().Unix()) > expiryTime {
		err := errors.New("token expired")
		return "", err
	} 

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", err
	}
	return claims["user"].(string), nil
// 	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(SECRET_KEY), nil
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !token.Valid {
// 		return "", err
// 	}
// 	return claims["id"].(string), nil
}