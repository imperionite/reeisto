package middleware

import (
<<<<<<< HEAD
	"fmt"
	"log"
=======
	"context"
	"fmt"
	"log"
	"os"
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

<<<<<<< HEAD
// RateLimitMiddleware returns a configurable rate limiter
func RateLimit(rdb *redis.Client, prefix string, limit int64, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		key := fmt.Sprintf("%s:%s", prefix, c.ClientIP())

		count, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			log.Println("Redis error (rate limit):", err)
			c.Next()
			return
		}

		if count == 1 {
			rdb.Expire(ctx, key, window)
		}

		if count > limit {
			c.JSON(429, gin.H{
				"success": false,
				"error":   "Too many requests",
			})
=======
var rdb *redis.Client

// Initialize Redis client
func init() {
	// Connect to Redis Cloud (provide your Redis Cloud credentials)
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),      // Read from the .env file
		Password: os.Getenv("REDIS_PASSWORD"), // Read from the .env file
		DB:       0,                           // Default DB is 0
	})
}

// RateLimiter middleware with Redis
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := fmt.Sprintf("rate_limit:%s", ip) // Redis key for each client IP

		limit := 10                   // Max requests allowed per second (adjust as needed)
		windowDuration := time.Minute // Time window for rate limit (e.g., 1 minute)

		ctx := context.Background()

		// Get the current request count from Redis
		count, err := rdb.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			log.Printf("Error fetching rate limit for IP %s: %v", ip, err)
			c.JSON(500, gin.H{"error": "Internal server error"})
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
			c.Abort()
			return
		}

<<<<<<< HEAD
		c.Next()
	}
}
=======
		// If the IP has exceeded the rate limit, deny the request
		if count >= limit {
			c.JSON(429, gin.H{"error": "Too many requests, try again later"})
			c.Abort()
			return
		}

		// Increment the request counter for this IP
		_, err = rdb.Incr(ctx, key).Result()
		if err != nil {
			log.Printf("Error incrementing rate limit for IP %s: %v", ip, err)
			c.JSON(500, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		}

		// Set an expiration time for the key, so the counter resets after the time window
		_, err = rdb.Expire(ctx, key, windowDuration).Result()
		if err != nil {
			log.Printf("Error setting expiration for IP %s: %v", ip, err)
			c.JSON(500, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		}

		// Allow the request to continue
		c.Next()
	}
}
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
