package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Authenticate validates the JWT for the requests and set the user info in the context.
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. get JWT from headers
		token := getJwt(c)

		// 2. parse token
		jwtToken, err := parseToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 3. Valid token
		if !isValidateToken(jwtToken) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 4. Get claims nd set them to the context
		setClaims(c, jwtToken)
	}
}

func getJwt(c *gin.Context) string {
	authorizationHeader := c.Request.Header.Get("Authorization")
	jwt := strings.TrimPrefix(authorizationHeader, "Bearer ")
	return jwt
}

func parseToken(token string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte("textcomparisonapipass754267"), nil
	})

	if err != nil {
		return nil, errors.New("Failed to unmarshal jwt")
	}

	return jwtToken, nil
}

func isValidateToken(jwToken *jwt.Token) bool {
	return jwToken.Valid
}

func setClaims(c *gin.Context, jwtToken *jwt.Token) {
	claims := jwtToken.Claims.(jwt.MapClaims)

	c.Set("user", claims["user"])
	c.Set("userID", claims["userID"])
}
