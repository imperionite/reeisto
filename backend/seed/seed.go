package seed

import (
	"log"
	"time"

<<<<<<< HEAD
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/imperionite/reeisto/backend/config"
	"github.com/imperionite/reeisto/backend/models"
)

// Seed inserts initial data for dev and prod. Safe to run multiple times.
=======
	"gorm.io/gorm"

	"github.com/imperionite/reetis/backend/models"
)

>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
func Seed(db *gorm.DB) {
	log.Println("Starting database seeding...")

	seedUsers(db)
	seedREEs(db)
<<<<<<< HEAD
	seedInventories(db)
	seedTransactions(db)

	log.Println("Database seeding completed successfully!")
=======

	log.Println("Database seeding completed.")
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
}

// ---------------------- USERS ----------------------
func seedUsers(db *gorm.DB) {
<<<<<<< HEAD
	passwords := map[string]string{
		"admin":  config.App.BootstrapAdminPassword,
		"trader": "traderpass",
	}

	users := []models.User{
		{
			Username: config.App.BootstrapAdminUser, // admin from config
			Email:    "admin-reetis@grr.la",
			Role:     "admin",
		},
		{
			Username: "johnstreet",
			Email:    "johnstreet@maildrop.cc",
			Role:     "trader",
		},
		{
			Username: "lpmmaterials",
			Email:    "lpmmaterials@maildrop.cc",
			Role:     "trader",
		},
		{
			Username: "neptuneresearch",
			Email:    "neptuneresearch@maildrop.cc",
			Role:     "trader",
		},
=======
	users := []models.User{
		{Username: "admin", Password: "supersecure", Role: "admin"},
		{Username: "trader1", Password: "traderpass", Role: "trader"},
		{Username: "trader2", Password: "traderpass", Role: "trader"},
		{Username: "trader3", Password: "traderpass", Role: "trader"},
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
	}

	for _, u := range users {
		var existing models.User
		err := db.Where("username = ? AND deleted_at IS NULL", u.Username).First(&existing).Error
<<<<<<< HEAD
		if err == nil {
			log.Printf("ℹ️  User already exists: %s", u.Username)
			continue
		}

		// Hash password
		pw := passwords["trader"]
		if u.Role == "admin" {
			pw = passwords["admin"]
		}

		hashed, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("❌ Failed to hash password for %s: %v", u.Username, err)
			continue
		}
		u.Password = string(hashed)

		if err := db.Create(&u).Error; err != nil {
			log.Printf("❌ Failed to seed user %s: %v", u.Username, err)
			continue
		}
		log.Printf("✅ Seeded user: %s (role: %s)", u.Username, u.Role)
=======
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&u).Error; err != nil {
					log.Printf("Failed to seed user %s: %v", u.Username, err)
					continue
				}
				log.Printf("Seeded user: %s", u.Username)
			} else {
				log.Printf("Error checking user %s: %v", u.Username, err)
			}
		} else {
			log.Printf("User already exists: %s", u.Username)
		}
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
	}
}

// ---------------------- REEs ----------------------
func seedREEs(db *gorm.DB) {
	rees := []models.REE{
<<<<<<< HEAD
		{Name: "Lanthanum", Symbol: "La", Category: "Light REE", MarketPrice: 8.50, Form: "oxide", PriceUnit: "USD/kg", Purity: 99.9},
		{Name: "Cerium", Symbol: "Ce", Category: "Light REE", MarketPrice: 7.80, Form: "oxide", PriceUnit: "USD/kg", Purity: 99.9},
		{Name: "Praseodymium", Symbol: "Pr", Category: "Light REE", MarketPrice: 95.00, Form: "oxide", PriceUnit: "USD/kg", Purity: 99.9},
		{Name: "Neodymium", Symbol: "Nd", Category: "Light REE", MarketPrice: 110.00, Form: "oxide", PriceUnit: "USD/kg", Purity: 99.9},
		{Name: "Samarium", Symbol: "Sm", Category: "Light REE", MarketPrice: 35.00, Form: "oxide", PriceUnit: "USD/kg", Purity: 99.9},
		{Name: "Europium", Symbol: "Eu", Category: "Light REE", MarketPrice: 450.00, Form: "oxide", PriceUnit: "USD/kg", Purity: 99.9},
		{Name: "Promethium", Symbol: "Pm", Category: "Light REE", MarketPrice: 4500.00, Form: "oxide", PriceUnit: "USD/kg", Purity: 99.9},
		{Name: "Gadolinium", Symbol: "Gd", Category: "Heavy REE", MarketPrice: 40.00, Form: "oxide", PriceUnit: "USD/kg", Purity: 99.9},
		{Name: "Terbium", Symbol: "Tb", Category: "Heavy REE", MarketPrice: 1400.00, Form: "oxide", PriceUnit: "USD/kg", Purity: 99.9},
		{Name: "Dysprosium", Symbol: "Dy", Category: "Heavy REE", MarketPrice: 750.00, Form: "oxide", PriceUnit: "USD/kg", Purity: 99.9},
		{Name: "Holmium", Symbol: "Ho", Category: "Heavy REE", MarketPrice: 120.00, Form: "oxide", PriceUnit: "USD/kg", Purity: 99.9},
		{Name: "Erbium", Symbol: "Er", Category: "Heavy REE", MarketPrice: 85.00, Form: "oxide", PriceUnit: "USD/kg", Purity: 99.9},
		{Name: "Thulium", Symbol: "Tm", Category: "Heavy REE", MarketPrice: 380.00, Form: "oxide", PriceUnit: "USD/kg", Purity: 99.9},
		{Name: "Ytterbium", Symbol: "Yb", Category: "Heavy REE", MarketPrice: 190.00, Form: "oxide", PriceUnit: "USD/kg", Purity: 99.9},
		{Name: "Lutetium", Symbol: "Lu", Category: "Heavy REE", MarketPrice: 950.00, Form: "oxide", PriceUnit: "USD/kg", Purity: 99.9},
		{Name: "Yttrium", Symbol: "Y", Category: "Heavy REE", MarketPrice: 25.00, Form: "oxide", PriceUnit: "USD/kg", Purity: 99.9},
		{Name: "Scandium", Symbol: "Sc", Category: "Heavy REE", MarketPrice: 3500.00, Form: "oxide", PriceUnit: "USD/kg", Purity: 99.9},
=======
		{Name: "Neodymium", Symbol: "Nd", Category: "Light REE", MarketPrice: 120},
		{Name: "Praseodymium", Symbol: "Pr", Category: "Light REE", MarketPrice: 90},
		{Name: "Samarium", Symbol: "Sm", Category: "Light REE", MarketPrice: 80},
		{Name: "Europium", Symbol: "Eu", Category: "Light REE", MarketPrice: 500},
		{Name: "Gadolinium", Symbol: "Gd", Category: "Light REE", MarketPrice: 45},
		{Name: "Terbium", Symbol: "Tb", Category: "Heavy REE", MarketPrice: 1500},
		{Name: "Dysprosium", Symbol: "Dy", Category: "Heavy REE", MarketPrice: 800},
		{Name: "Holmium", Symbol: "Ho", Category: "Heavy REE", MarketPrice: 120},
		{Name: "Erbium", Symbol: "Er", Category: "Heavy REE", MarketPrice: 90},
		{Name: "Thulium", Symbol: "Tm", Category: "Heavy REE", MarketPrice: 400},
		{Name: "Ytterbium", Symbol: "Yb", Category: "Heavy REE", MarketPrice: 200},
		{Name: "Lutetium", Symbol: "Lu", Category: "Heavy REE", MarketPrice: 1000},
		{Name: "Lanthanum", Symbol: "La", Category: "Light REE", MarketPrice: 10},
		{Name: "Cerium", Symbol: "Ce", Category: "Light REE", MarketPrice: 8},
		{Name: "Yttrium", Symbol: "Y", Category: "Heavy REE", MarketPrice: 30},
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
	}

	now := time.Now()
	for i := range rees {
		rees[i].CreatedAt = now
		rees[i].UpdatedAt = now
	}

	for _, r := range rees {
		var existing models.REE
<<<<<<< HEAD
		err := db.Where("symbol = ? AND deleted_at IS NULL", r.Symbol).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&r).Error; err != nil {
				log.Printf("❌ Failed to seed REE %s: %v", r.Name, err)
				continue
			}
			log.Printf("✅ Seeded REE: %s (%s)", r.Name, r.Symbol)
		} else {
			log.Printf("ℹ️  REE already exists: %s", r.Name)
		}
	}
}

// ---------------------- INVENTORIES ----------------------
func seedInventories(db *gorm.DB) {
	inventories := []models.Inventory{
		{ElementID: 1, Quantity: 150.5, WarehouseLocation: "WH-A-01"},
		{ElementID: 2, Quantity: 89.0, WarehouseLocation: "WH-A-02"},
		{ElementID: 3, Quantity: 200.75, WarehouseLocation: "WH-B-01"},
		{ElementID: 7, Quantity: 75.25, WarehouseLocation: "WH-B-02"},
		{ElementID: 8, Quantity: 300.0, WarehouseLocation: "WH-C-01"},
		{ElementID: 1, Quantity: 50.0, WarehouseLocation: "WH-C-02"},
		{ElementID: 14, Quantity: 5.0, WarehouseLocation: "WH-D-01"},
		{ElementID: 14, Quantity: 3.5, WarehouseLocation: "WH-D-02"},
	}

	for _, inv := range inventories {
		var existing models.Inventory
		err := db.Where("element_id = ? AND warehouse_location = ?", inv.ElementID, inv.WarehouseLocation).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&inv).Error; err != nil {
				log.Printf("❌ Failed to seed inventory for element %d: %v", inv.ElementID, err)
				continue
			}
			log.Printf("✅ Seeded inventory: Element %d @ %s (%.2f)", inv.ElementID, inv.WarehouseLocation, inv.Quantity)
		} else {
			log.Printf("ℹ️  Inventory already exists: Element %d @ %s", inv.ElementID, inv.WarehouseLocation)
		}
	}
}

// ---------------------- TRANSACTIONS ----------------------
func seedTransactions(db *gorm.DB) {
	transactions := []models.Transaction{
		{UserID: 2, ElementID: 1, Type: "buy", Quantity: 10.5, Price: 125.50, WarehouseLocation: "WH-A-01"},
		{UserID: 2, ElementID: 7, Type: "sell", Quantity: 5.0, Price: 1520.00, WarehouseLocation: "WH-B-02"},
		{UserID: 3, ElementID: 2, Type: "buy", Quantity: 25.0, Price: 92.30, WarehouseLocation: "WH-A-02"},
		{UserID: 3, ElementID: 8, Type: "buy", Quantity: 15.75, Price: 1010.25, WarehouseLocation: "WH-C-01"},
		{UserID: 4, ElementID: 3, Type: "sell", Quantity: 8.0, Price: 82.50, WarehouseLocation: "WH-B-01"},
		{UserID: 4, ElementID: 1, Type: "buy", Quantity: 20.0, Price: 122.75, WarehouseLocation: "WH-C-02"},
	}

	for _, tx := range transactions {
		var existing models.Transaction
		err := db.Where(
			"user_id = ? AND element_id = ? AND quantity = ? AND price = ?",
			tx.UserID, tx.ElementID, tx.Quantity, tx.Price,
		).First(&existing).Error

		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&tx).Error; err != nil {
				log.Printf("❌ Failed to seed transaction (user:%d, elem:%d): %v", tx.UserID, tx.ElementID, err)
				continue
			}
			log.Printf("✅ Seeded transaction: User %d %s %.2f of Element %d @ $%.2f (%s)",
				tx.UserID, tx.Type, tx.Quantity, tx.ElementID, tx.Price, tx.WarehouseLocation)
		} else {
			log.Printf("ℹ️  Transaction already exists: User %d, Element %d", tx.UserID, tx.ElementID)
		}
	}
}
=======
		err := db.Where("name = ? AND deleted_at IS NULL", r.Name).First(&existing).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&r).Error; err != nil {
					log.Printf("Failed to seed REE %s: %v", r.Name, err)
					continue
				}
				log.Printf("Seeded REE: %s", r.Name)
			} else {
				log.Printf("Error checking REE %s: %v", r.Name, err)
			}
		} else {
			log.Printf("REE already exists: %s", r.Name)
		}
	}
}
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
