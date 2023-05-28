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

		header := ctx.GetHeader("Authorization")

		if header == "" {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("no authorization header provided"))
			return
		}

		headerToken := strings.TrimPrefix(header, "Bearer ")

		token, err := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		ctx.Set("userId", claims["userId"])
		ctx.Set("username", claims["username"])
		ctx.Set("role", claims["role"])
		ctx.Next()
	}
}
