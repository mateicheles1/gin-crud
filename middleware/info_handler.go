package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InfoMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] - %s - %s - code: %d \n",
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.StatusCode,
		)
	})
}
