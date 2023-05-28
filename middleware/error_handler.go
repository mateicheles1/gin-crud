package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateicheles1/golang-crud/logs"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Next()

		for _, err := range ctx.Errors {

			switch ctx.Writer.Status() {

			case http.StatusBadRequest:
				logs.Logger.Error().
					Str("Method", ctx.Request.Method).
					Str("Path", ctx.Request.URL.Path).
					Int("Status code", http.StatusBadRequest).
					Msgf("Bad request: %s", err)

				ctx.JSON(http.StatusBadRequest, "invalid request body JSON syntax")
				return

			case http.StatusInternalServerError:
				logs.Logger.Error().
					Str("Method", ctx.Request.Method).
					Str("Path", ctx.Request.URL.Path).
					Int("Status code", http.StatusInternalServerError).
					Msgf("Internal server error: %s", err)

				ctx.JSON(http.StatusInternalServerError, "something went wrong")
				return

			case http.StatusUnauthorized:
				logs.Logger.Error().
					Str("Method", ctx.Request.Method).
					Str("Path", ctx.Request.URL.Path).
					Int("Status code", http.StatusUnauthorized).
					Msgf("Unauthorized: %s", err)

				ctx.JSON(http.StatusUnauthorized, "invalid credentials")
				return

			case http.StatusForbidden:
				logs.Logger.Error().
					Str("Method", ctx.Request.Method).
					Str("Path", ctx.Request.URL.Path).
					Int("Status code", http.StatusForbidden).
					Msgf("Forbidden: %s", err)

				ctx.JSON(http.StatusForbidden, "action not allowed")
				return
			}

		}
	}
}
