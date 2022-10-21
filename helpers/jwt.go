package helpers

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte(os.Getenv("SECRETKEY"))

type JWTClaim struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(id int, role string) (tokenString string, err error) {
	expirationTime := time.Now().Add(5 * time.Hour)
	claims := &JWTClaim{
		ID:   id,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}

func ClaimToken(c *gin.Context) JWTClaim {

	hmacSecret := os.Getenv("SECRETKEY")
	reqToken := c.GetHeader(os.Getenv("HEADERAUTH"))

	prefix := "Bearer "
	tokenString := strings.TrimPrefix(reqToken, prefix)

	claims := &JWTClaim{}
	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	return *claims
}
