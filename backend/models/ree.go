package models

import (
	"time"

	"gorm.io/gorm"
)

type REE struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"not null;index"`
	Symbol      string  `gorm:"unique;not null"`
	Category    string  `gorm:"not null"`
	MarketPrice float64 `gorm:"not null"`
<<<<<<< HEAD

	Form      string  `gorm:"default:oxide"`
	PriceUnit string  `gorm:"default:USD/kg"`
	Purity    float64 `gorm:"default:99.9"`

	Inventories []Inventory `gorm:"foreignKey:ElementID"` 

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
=======
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
}
