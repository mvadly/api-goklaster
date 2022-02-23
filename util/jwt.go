package util

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type M map[string]interface{}

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte(os.Getenv("JWT_SIGN_KEY"))

type DataJwt struct {
	Username   string `json:"username"`
	Permission int    `json:"permission"`
}

type MyClaims struct {
	jwt.StandardClaims
	Jwtsession interface{}
}

func GetToken(dataSession DataJwt) (string, error) {
	var APPLICATION_NAME = os.Getenv("JWT_APPNAME")
	var exp, _ = strconv.Atoi(os.Getenv("JWT_EXP"))
	var LOGIN_EXPIRATION_DURATION = time.Duration(exp) * time.Minute
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Jwtsession: dataSession,
	}
	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func CheckToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return JWT_SIGNATURE_KEY, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Token is not valid")
	}

	return claims, nil

}
