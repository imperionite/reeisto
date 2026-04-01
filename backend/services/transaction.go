package services

import (
	"errors"
<<<<<<< HEAD

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/imperionite/reeisto/backend/dto"
	"github.com/imperionite/reeisto/backend/models"
)

type Transaction struct {
	models.Transaction
}

=======
	"time"

	"github.com/imperionite/reetis/backend/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
type TransactionService struct {
	db *gorm.DB
}

func NewTransactionService(db *gorm.DB) *TransactionService {
	return &TransactionService{db: db}
}

<<<<<<< HEAD
func (s *TransactionService) Create(userID uint, input *dto.CreateTransactionDTO) (*Transaction, error) {

	// INPUT VALIDATION

	if input.Type != "buy" && input.Type != "sell" {
		return nil, errors.New("invalid transaction type")
	}

	if input.Quantity <= 0 {
		return nil, errors.New("quantity must be greater than 0")
	}

	if input.Price <= 0 {
		return nil, errors.New("price must be greater than 0")
	}

	if input.WarehouseLocation == "" {
		return nil, errors.New("warehouse location is required")
	}

	var resultTx *Transaction

	err := s.db.Transaction(func(tx *gorm.DB) error {

		// -----------------------------
		// Validate REE
		// -----------------------------
		var ree models.REE
		if err := tx.First(&ree, input.ElementID).Error; err != nil {
			return errors.New("REE element not found")
		}

		//  Validate user
		var user models.User
		if err := tx.First(&user, userID).Error; err != nil {
			return errors.New("user not found")
		}

		// Find inventory WITH LOCK
		var inventory models.Inventory
		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("element_id = ? AND warehouse_location = ?", input.ElementID, input.WarehouseLocation).
			First(&inventory).Error

		// SELL: no inventory
		if err != nil && input.Type == "sell" {
			return errors.New("no inventory available in this warehouse")
		}

		// BUY LOGIC
		if input.Type == "buy" {

			if err != nil {
				// Create new inventory row
				inventory = models.Inventory{
					ElementID:         input.ElementID,
					Quantity:          input.Quantity,
					WarehouseLocation: input.WarehouseLocation,
				}

				if err := tx.Create(&inventory).Error; err != nil {
					return err
				}

			} else {
				// Update existing inventory
				inventory.Quantity += input.Quantity

				if err := tx.Save(&inventory).Error; err != nil {
					return err
				}
			}
		}

		// SELL LOGIC
		if input.Type == "sell" {

			if inventory.Quantity < input.Quantity {
				return errors.New("insufficient inventory in this warehouse")
			}

			inventory.Quantity -= input.Quantity

			if err := tx.Save(&inventory).Error; err != nil {
				return err
			}
		}

		// Create transaction record
		newTx := &Transaction{
			Transaction: models.Transaction{
				UserID:            userID,
				ElementID:         input.ElementID,
				Type:              input.Type,
				Quantity:          input.Quantity,
				Price:             input.Price,
				WarehouseLocation: input.WarehouseLocation,
			},
		}

		if err := tx.Create(newTx).Error; err != nil {
			return err
		}

		resultTx = newTx
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Reload with relations
	s.db.Preload("User").Preload("Element").First(resultTx, resultTx.ID)

	return resultTx, nil
}

func (s *TransactionService) List(userID uint, limit, offset int) ([]Transaction, error) {
	var transactions []Transaction

	err := s.db.
		Preload("Element").
		Where("user_id = ?", userID).
		Limit(limit).
		Offset(offset).
		Order("timestamp desc").
		Find(&transactions).Error

	return transactions, err
=======
func (s *TransactionService) Create(tx *models.Transaction) error {
	return s.db.Transaction(func(txDB *gorm.DB) error {

		// VALIDATION
		if tx.Quantity <= 0 {
			return errors.New("quantity must be greater than zero")
		}

		if tx.Price <= 0 {
			return errors.New("price must be greater than zero")
		}

		if tx.Type != "buy" && tx.Type != "sell" {
			return errors.New("invalid transaction type")
		}

		var inventory models.Inventory

		// LOCK ROW
		if err := txDB.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("element_id = ?", tx.ElementID).
			First(&inventory).Error; err != nil {
			return errors.New("inventory not found")
		}

		switch tx.Type {
		case "sell":
			if inventory.Quantity < tx.Quantity {
				return errors.New("insufficient inventory")
			}
			inventory.Quantity -= tx.Quantity

		case "buy":
			inventory.Quantity += tx.Quantity
		}

		if err := txDB.Save(&inventory).Error; err != nil {
			return err
		}

		if tx.Timestamp.IsZero() {
			tx.Timestamp = time.Now()
		}

		return txDB.Create(tx).Error
	})
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
}
