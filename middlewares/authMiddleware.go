package middlewares

import (
	"net/http"
	"os"
	"sosialMedia/helpers"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		hmacSecret := os.Getenv("SECRETKEY")
		reqToken := c.GetHeader(os.Getenv("HEADERAUTH"))

		prefix := "Bearer "
		tokenString := strings.TrimPrefix(reqToken, prefix)

		if tokenString == "" || reqToken == tokenString {
			helpers.ResponseJson(c, http.StatusUnauthorized, false, "request does not contain an access token", nil)
			c.Abort()
			return
		}

		err := helpers.ValidateToken(tokenString)
		if err != nil {
			helpers.ResponseJson(c, http.StatusUnauthorized, false, "request does not contain an valid token", nil)
			c.Abort()
			return
		}

		claims := &helpers.JWTClaim{}
		jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return hmacSecret, nil
		})

		c.Next()
	}
}
