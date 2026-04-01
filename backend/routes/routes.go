package routes

import (
<<<<<<< HEAD
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"github.com/imperionite/reeisto/backend/controllers"
	"github.com/imperionite/reeisto/backend/middleware"
	"github.com/imperionite/reeisto/backend/services"
)

func SetupRoutes(router *gin.RouterGroup, db *gorm.DB, rdb *redis.Client) {
	// Initialize services
	invService := services.NewInventoryService(db)
	reeService := services.NewREEService(db)
	txService := services.NewTransactionService(db)

	// Public routes
	public := router.Group("/")

	// 5 requests per minute for login/register
	authLimiter := middleware.RateLimit(rdb, "rl:auth", 5, time.Minute)

	public.POST("/login", authLimiter, controllers.Login(db))
	public.POST("/register", authLimiter, controllers.Register(db))

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())

	// REE routes
	protected.GET("/rees", controllers.GetREEs(reeService, rdb))
	protected.POST("/rees", middleware.RequireRole("admin"), controllers.CreateREE(reeService, rdb))
	protected.PUT("/rees/:id", middleware.RequireRole("admin"), controllers.UpdateREE(reeService, rdb))
	protected.DELETE("/rees/:id", middleware.RequireRole("admin"), controllers.DeleteREE(reeService, rdb))

	// Inventory routes
	protected.GET("/inventories", controllers.GetInventories(invService))
	protected.POST("/inventories", middleware.RequireRole("admin"), controllers.CreateInventory(invService))

	// Transaction routes
	protected.GET("/transactions", controllers.GetTransactions(txService))
	protected.POST("/transactions", controllers.CreateTransaction(txService))
}
=======
	"github.com/gin-gonic/gin"
	"github.com/imperionite/reetis/backend/controllers"
	"github.com/imperionite/reetis/backend/middleware"
)

func SetupRoutes(router *gin.RouterGroup) {
	// Public
	router.POST("/login", controllers.Login)

	// Protected
	auth := router.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())

	// REE
	auth.GET("/rees", controllers.GetREEs)
	auth.POST("/rees", middleware.RequireRole("admin"), controllers.CreateREE)
	auth.PUT("/rees/:id", middleware.RequireRole("admin"), controllers.UpdateREE)
	auth.DELETE("/rees/:id", middleware.RequireRole("admin"), controllers.DeleteREE)

	// Inventory
	auth.GET("/inventories", controllers.GetInventories)
	auth.POST("/inventories", middleware.RequireRole("admin"), controllers.CreateInventory)

	// Transactions
	auth.GET("/transactions", controllers.GetTransactions)
	auth.POST("/transactions", controllers.CreateTransaction)
}
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
