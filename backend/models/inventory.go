package models

import (
	"time"

	"gorm.io/gorm"
)

type Inventory struct {
	ID                uint    `gorm:"primaryKey"`
<<<<<<< HEAD
	ElementID         uint    `gorm:"not null;index"` // FK to REE
=======
	ElementID         uint    `gorm:"not null;unique"` // FK to REE
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
	Element           REE     `gorm:"foreignKey:ElementID"`
	Quantity          float64 `gorm:"not null"`
	WarehouseLocation string  `gorm:"not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
<<<<<<< HEAD


=======
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
