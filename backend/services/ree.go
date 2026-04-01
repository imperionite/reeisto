package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"github.com/imperionite/reeisto/backend/dto"
	"github.com/imperionite/reeisto/backend/models"
)

type REE struct {
	models.REE
}

type REEService struct {
	db *gorm.DB
}

func NewREEService(db *gorm.DB) *REEService {
	return &REEService{db: db}
}

func (s *REEService) ListWithCache(rdb *redis.Client) ([]REE, string, error) {
	ctx := context.Background()
	cacheKey := "rees:all"

	// Try cache first
	val, err := rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		var rees []REE
		if json.Unmarshal([]byte(val), &rees) == nil {
			return rees, "cache", nil
		}
	}

	// Database fallback
	var rees []REE
	if err := s.db.Preload("Inventories").Find(&rees).Error; err != nil {
		return nil, "", err
	}

	// Cache results
	data, _ := json.Marshal(rees)
	rdb.Set(ctx, cacheKey, data, 5*time.Minute)

	return rees, "database", nil
}

func (s *REEService) Create(input *dto.CreateREEDTO, rdb *redis.Client) (*REE, error) {
	ree := &REE{REE: models.REE{
		Name:        input.Name,
		Symbol:      input.Symbol,
		Category:    input.Category,
		MarketPrice: input.MarketPrice,
		Form:        defaultIfEmpty(input.Form, "oxide"),
		PriceUnit:   defaultIfEmpty(input.PriceUnit, "USD/kg"),
		Purity:      defaultIfZero(input.Purity, 99.9),
	}}

	if err := s.db.Create(ree).Error; err != nil {
		return nil, err
	}

	s.invalidateCache(rdb)

	return ree, nil
}

func (s *REEService) Update(id uint, input *dto.UpdateREEDTO, rdb *redis.Client) (*REE, error) {
	ree := &REE{}
	if err := s.db.First(ree, id).Error; err != nil {
		return nil, err
	}

	updated := false
	if input.Name != "" {
		ree.Name = input.Name
		updated = true
	}
	if input.Symbol != "" {
		ree.Symbol = input.Symbol
		updated = true
	}
	if input.Category != "" {
		ree.Category = input.Category
		updated = true
	}
	if input.MarketPrice > 0 {
		ree.MarketPrice = input.MarketPrice
		updated = true
	}

	if input.Form != "" {
		ree.Form = input.Form
		updated = true
	}
	if input.PriceUnit != "" {
		ree.PriceUnit = input.PriceUnit
		updated = true
	}

	if input.Purity > 0 {
		ree.Purity = input.Purity
		updated = true
	}

	if updated {
		s.db.Save(ree)
	}

	s.invalidateCache(rdb)

	s.db.Preload("Inventories").First(ree, id)
	return ree, nil
}

func (s *REEService) Delete(id uint, rdb *redis.Client) error {
	if err := s.db.Delete(&models.REE{}, id).Error; err != nil {
		return err
	}

	s.invalidateCache(rdb)
	return nil
}

// helper
func defaultIfZero(val float64, def float64) float64 {
	if val == 0 {
		return def
	}
	return val
}

func defaultIfEmpty(val, def string) string {
	if val == "" {
		return def
	}
	return val
}

func (s *REEService) invalidateCache(rdb *redis.Client) {
	ctx := context.Background()
	rdb.Del(ctx, "rees:all")
}

