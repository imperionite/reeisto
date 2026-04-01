package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
<<<<<<< HEAD
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"

	"github.com/imperionite/reeisto/backend/config"
	"github.com/imperionite/reeisto/backend/middleware"
	"github.com/imperionite/reeisto/backend/routes"
	"github.com/imperionite/reeisto/backend/seed"
=======
	"github.com/joho/godotenv"

	"github.com/imperionite/reetis/backend/config"
	"github.com/imperionite/reetis/backend/middleware"
	"github.com/imperionite/reetis/backend/routes"
	"github.com/imperionite/reetis/backend/seed"
	"github.com/imperionite/reetis/backend/utils"
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
)

func main() {
	_ = godotenv.Load()
	config.LoadConfig()

<<<<<<< HEAD
=======
	// Set Gin mode
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

<<<<<<< HEAD
	// ---------------- DB ----------------
	db := config.ConnectDatabase()

	if config.App.Env == "development" {
		if err := config.AutoMigrateModels(db); err != nil {
			log.Fatal("Auto-migration failed:", err)
		}
		log.Println("Database migration completed")
	}

	// ---------------- REDIS ----------------
	rdb := initRedis()

	// ---------------- ENV-BASED BEHAVIOR ----------------
	env := os.Getenv("APP_ENV")
	seedData := os.Getenv("SEED_DATA") == "true"

	switch env {
	case "development":
		// Always seed full dev data
		seed.Seed(db)

	case "production":
		if seedData {
			// Seed full dev-like data in production
			seed.Seed(db)
		} else {
			// Only bootstrap the admin user but set the  BOOTSTRAP_ADMIN_USER and BOOTSTRAP_ADMIN_PASSWORD env variables
			config.BootstrapAdmin(db)
		}
	}

	// ---------------- ROUTER ----------------
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.Use(setupCORS())

	// Inject Redis into middleware
	router.Use(middleware.RateLimit(rdb, "rate_limit", 100, time.Minute))

	// Health check
	router.GET("/health", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()

		dbErr := config.PingDB(ctx)
		redisErr := rdb.Ping(ctx).Err()

		status := "ok"
		if dbErr != nil || redisErr != nil {
			status = "degraded"
		}

		c.JSON(200, gin.H{
			"status": status,
		})
	})

	// Routes with DI
	api := router.Group("/api/v1")
	routes.SetupRoutes(api, db, rdb)

	// ---------------- SERVER ----------------
	srv := &http.Server{
		Addr:         ":" + config.App.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("Server running on port %s", config.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
=======
	// Initialize Redis
	// Ensure you have the correct Redis initialization in middleware

	db := config.ConnectDatabase()

	// Reset DB (dev only)
	if len(os.Args) > 1 && os.Args[1] == "--reset" {
		config.ResetDatabase(db)
	}

	config.AutoMigrateModels(db)
	seed.Seed(db)

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get DB instance:", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Database not reachable:", err)
	}

	// Router
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	// CORS origins from env
	originsStr := config.App.CORSOrigins
	if originsStr == "" {
		log.Fatal("CORS_ORIGINS environment variable is required")
	}

	allowOrigins := strings.Split(strings.TrimSpace(originsStr), ",")
	for i := range allowOrigins {
		allowOrigins[i] = strings.TrimSpace(allowOrigins[i])
	}

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Apply Redis-based Rate Limiter
	router.Use(middleware.RateLimit())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		sqlDB, _ := config.DB.DB()
		err := sqlDB.Ping()

		status := "ok"
		if err != nil {
			status = "degraded"
		}

		utils.Success(c, gin.H{"status": status})
	})

	// API versioning
	api := router.Group("/api/v1")
	routes.SetupRoutes(api)

	srv := &http.Server{
		Addr:    ":" + config.App.Port,
		Handler: router,
	}

	log.Printf("event=server_start port=%s", config.App.Port)

	// Start server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

<<<<<<< HEAD
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down...")
	srv.Shutdown(ctx)
}

func initRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.App.RedisAddr,
		Password: config.App.RedisPassword,
		DB:       0,
	})
=======
	log.Println("Shutting down server...")
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

<<<<<<< HEAD
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal("Redis connection failed:", err)
	}

	log.Println("Redis connected")
	return rdb
}

func setupCORS() gin.HandlerFunc {
	origins := strings.Split(config.App.CORSOrigins, ",")

	return cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
=======
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
