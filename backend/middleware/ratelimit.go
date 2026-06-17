package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/opencti-lite/backend/database"
	"github.com/gin-gonic/gin"
)

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if database.RDB == nil {
			c.Next()
			return
		}

		ip := c.ClientIP()
		key := fmt.Sprintf("rate:%s", ip)
		ctx := context.Background()

		count, err := database.RDB.Incr(ctx, key).Result()
		if err != nil {
			// Fail open on Redis connectivity issue
			c.Next()
			return
		}

		if count == 1 {
			database.RDB.Expire(ctx, key, 1*time.Minute)
		}

		if count > 100 {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded. 100 requests per minute allowed.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
