package middlewares

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

var limiter *rate.Limiter

func init() {
	limiter = rate.NewLimiter(100, 1024)
}

func RateLimiter(c *gin.Context) {
	if !limiter.Allow() {
		c.JSON(http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests))
		c.Abort()
		return
	}
	c.Next()
}
