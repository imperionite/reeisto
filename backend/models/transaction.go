package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        uint           `gorm:"primaryKey"`
	UserID    uint           `gorm:"not null"` // FK to User
	User      User           `gorm:"foreignKey:UserID"`
	ElementID uint           `gorm:"not null"` // FK to REE
	Element   REE            `gorm:"foreignKey:ElementID"`
	Type      string         `gorm:"type:varchar(10);not null"` // 'buy' or 'sell'
	Quantity  float64        `gorm:"not null"`
	Price     float64        `gorm:"not null"`
	Timestamp time.Time      `gorm:"autoCreateTime"`
<<<<<<< HEAD
	WarehouseLocation string `gorm:"not null"`
=======
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}