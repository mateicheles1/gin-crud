package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("no authorization header provided"))
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("invalid header"))
			return
		}

		authToken := strings.TrimPrefix(authHeader, "Bearer ")

		parsedToken, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)

		if !ok || !parsedToken.Valid {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		ctx.Set("userId", claims["userId"])
		ctx.Set("username", claims["username"])
		ctx.Set("role", claims["role"])
		ctx.Next()
	}
}
