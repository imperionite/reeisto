package services

import (
<<<<<<< HEAD
	"github.com/imperionite/reeisto/backend/models"
	"gorm.io/gorm"
)

type Inventory struct {
	models.Inventory
}

=======
	"github.com/imperionite/reetis/backend/models"
	"gorm.io/gorm"
)

>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
type InventoryService struct {
	db *gorm.DB
}

func NewInventoryService(db *gorm.DB) *InventoryService {
	return &InventoryService{db: db}
}

<<<<<<< HEAD
// Create inventory - FIXED
func (s *InventoryService) Create(invData *models.Inventory) (*Inventory, error) {
	inv := &Inventory{Inventory: *invData}
	
	if err := s.db.Create(inv).Error; err != nil {
		return nil, err
	}
	
	// Reload with relations
	s.db.Preload("Element").First(&inv.Inventory, inv.ID)
	return inv, nil
}

// List inventories
func (s *InventoryService) List(limit, offset int) ([]Inventory, error) {
	var inventories []Inventory
	err := s.db.Model(&models.Inventory{}).
		Preload("Element").
		Limit(limit).Offset(offset).
		Order("created_at desc").
		Find(&inventories).Error
	return inventories, err
=======
func (s *InventoryService) Create(inv *models.Inventory) error {
	return s.db.Create(inv).Error
}

func (s *InventoryService) List(limit, offset int) ([]models.Inventory, error) {
	var inv []models.Inventory
	err := s.db.Preload("Element").Limit(limit).Offset(offset).Find(&inv).Error
	return inv, err
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
}