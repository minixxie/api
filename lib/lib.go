package lib

import "github.com/dgrijalva/jwt-go"
import "time"

func GenJWT(userId int64) string {
	jwtToken := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := jwtToken.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	// claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Generate encoded token and send it as response.

	t, err := jwtToken.SignedString([]byte("hello123"))
	if err != nil {
	    return ""
	}

	return t
}