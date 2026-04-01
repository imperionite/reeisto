package config

import (
<<<<<<< HEAD
	"context"
=======
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

<<<<<<< HEAD
	"github.com/imperionite/reeisto/backend/models"
=======
	"github.com/imperionite/reetis/backend/models"
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
)

var (
	DB   *gorm.DB
	once sync.Once
)

func ConnectDatabase() *gorm.DB {
	once.Do(func() {
<<<<<<< HEAD
		var db *gorm.DB
		var err error

		if App.DBURL != "" {
			// Use full URL (Render, Docker, K8s)
			db, err = gorm.Open(postgres.Open(App.DBURL), &gorm.Config{})
		} else {
			// fallback to manual config (local dev)
			dsn := fmt.Sprintf(
				"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
				App.DBHost,
				App.DBUser,
				App.DBPass,
				App.DBName,
				App.DBPort,
				App.DBSSLMode,
			)
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		}

=======
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			App.DBHost,
			App.DBUser,
			App.DBPass,
			App.DBName,
			App.DBPort,
			App.DBSSLMode,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}

		log.Println("Database connected successfully")
<<<<<<< HEAD
=======

>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
		DB = db
	})

	return DB
}

// AutoMigrateModels migrates all tables
<<<<<<< HEAD
func AutoMigrateModels(db *gorm.DB) error {
	return db.AutoMigrate(
=======
func AutoMigrateModels(db *gorm.DB) {
	err := db.AutoMigrate(
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
		&models.User{},
		&models.REE{},
		&models.Inventory{},
		&models.Transaction{},
	)
<<<<<<< HEAD
}

func PingDB(ctx context.Context) error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.PingContext(ctx)
=======
	if err != nil {
		log.Fatal("Auto-migration failed:", err)
	}

	log.Println("Database migration completed")
}

func ResetDatabase(db *gorm.DB) {
	db.Migrator().DropTable(
		&models.Transaction{},
		&models.Inventory{},
		&models.REE{},
		&models.User{},
	)
	log.Println("Database reset complete")
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
}
