package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

<<<<<<< HEAD
	"github.com/imperionite/reeisto/backend/dto"
	"github.com/imperionite/reeisto/backend/models"
	"github.com/imperionite/reeisto/backend/services"
	"github.com/imperionite/reeisto/backend/utils"
)

func GetInventories(invService *services.InventoryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
		offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

		inventories, err := invService.List(limit, offset)
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, "Failed to fetch inventories")
			return
		}

		// utils.Success(c, inventories)

		utils.Success(c, gin.H{
			"data": inventories,
			"meta": gin.H{
				"limit":  limit,
				"offset": offset,
				"count":  len(inventories),
			},
		})
	}
}

func CreateInventory(invService *services.InventoryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.CreateInventoryDTO
		if err := c.ShouldBindJSON(&input); err != nil {
			utils.Error(c, http.StatusBadRequest, "Invalid input")
			return
		}

		// FIXED: Use models.Inventory directly
		invData := &models.Inventory{
			ElementID:         input.ElementID,
			Quantity:          input.Quantity,
			WarehouseLocation: input.WarehouseLocation,
		}

		inventory, err := invService.Create(invData)
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, "Failed to create inventory")
			return
		}

		utils.Success(c, inventory)
	}
}
=======
	"github.com/imperionite/reetis/backend/config"
	"github.com/imperionite/reetis/backend/dto"
	"github.com/imperionite/reetis/backend/models"
	"github.com/imperionite/reetis/backend/utils"
)

func GetInventories(c *gin.Context) {
	db := config.DB
	var inv []models.Inventory

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	offset := (page - 1) * limit

	if err := db.Preload("Element").
		Limit(limit).Offset(offset).
		Find(&inv).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to fetch inventories")
		return
	}

	utils.Success(c, inv)
}

func CreateInventory(c *gin.Context) {
	var req dto.CreateInventoryDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	inv := models.Inventory{
		ElementID:         req.ElementID,
		Quantity:          req.Quantity,
		WarehouseLocation: req.WarehouseLocation,
	}

	if err := config.DB.Create(&inv).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to create inventory")
		return
	}

	utils.Success(c, inv)
}
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
