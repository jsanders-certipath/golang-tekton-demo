package ratelimiter

import (
	"time"

	"github.com/axiaoxin-com/ratelimiter"
	"github.com/gin-gonic/gin"
)

type Service interface {
	InMemRateLimiter() gin.HandlerFunc
}

type service struct{}

func New() Service {
	return &service{}
}

func (s service) InMemRateLimiter() gin.HandlerFunc {
	rateLimiter := ratelimiter.GinMemRatelimiter(ratelimiter.GinRatelimiterConfig{
		LimitKey: func(c *gin.Context) string {
			return c.ClientIP()
		},
		LimitedHandler: func(c *gin.Context) {
			c.JSON(500, "too many requests from: "+c.ClientIP())
			c.Abort()
			return
		},
		TokenBucketConfig: func(*gin.Context) (time.Duration, int) {
			return time.Second * 1, 1
		},
	})
	return rateLimiter
}
