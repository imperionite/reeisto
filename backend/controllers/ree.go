package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
<<<<<<< HEAD
	"github.com/go-redis/redis/v8"

	"github.com/imperionite/reeisto/backend/dto"
	"github.com/imperionite/reeisto/backend/services"
	"github.com/imperionite/reeisto/backend/utils"
)

func GetREEs(reeService *services.REEService, rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ---------------------------
		// Pagination params
		// ---------------------------
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "50"))
		if err != nil || limit <= 0 {
			limit = 50
		}

		offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
		if err != nil || offset < 0 {
			offset = 0
		}

		// Optional: max cap (protect API)
		if limit > 100 {
			limit = 100
		}

		// ---------------------------
		// Fetch from service (cached full list)
		// ---------------------------
		rees, source, err := reeService.ListWithCache(rdb)
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, "Failed to fetch REEs")
			return
		}

		total := len(rees)

		// ---------------------------
		// Apply pagination (safe slicing)
		// ---------------------------
		end := offset + limit
		if offset > total {
			offset = total
		}
		if end > total {
			end = total
		}

		paginated := rees[offset:end]

		// ---------------------------
		// Response
		// ---------------------------
		utils.Success(c, gin.H{
			"source": source,
			"data":   paginated,
			"meta": gin.H{
				"total":  total,
				"limit":  limit,
				"offset": offset,
				"count":  len(paginated),
			},
		})
	}
}

func CreateREE(reeService *services.REEService, rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.CreateREEDTO
		if err := c.ShouldBindJSON(&input); err != nil {
			utils.Error(c, http.StatusBadRequest, "Invalid input")
			return
		}

		ree, err := reeService.Create(&input, rdb)
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, "Failed to create REE")
			return
		}

		utils.Success(c, ree)
	}
}

func UpdateREE(reeService *services.REEService, rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			utils.Error(c, http.StatusBadRequest, "Invalid ID")
			return
		}

		var input dto.UpdateREEDTO
		if err := c.ShouldBindJSON(&input); err != nil {
			utils.Error(c, http.StatusBadRequest, "Invalid input")
			return
		}

		ree, err := reeService.Update(uint(id), &input, rdb)
		if err != nil {
			utils.Error(c, http.StatusNotFound, "REE not found")
			return
		}

		utils.Success(c, ree)
	}
}

func DeleteREE(reeService *services.REEService, rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			utils.Error(c, http.StatusBadRequest, "Invalid ID")
			return
		}

		if err := reeService.Delete(uint(id), rdb); err != nil {
			utils.Error(c, http.StatusNotFound, "REE not found")
			return
		}

		utils.Success(c, gin.H{"message": "REE deleted successfully"})
	}
}
=======

	"github.com/imperionite/reetis/backend/config"
	"github.com/imperionite/reetis/backend/dto"
	"github.com/imperionite/reetis/backend/models"
	"github.com/imperionite/reetis/backend/utils"
)

func GetREEs(c *gin.Context) {
	var rees []models.REE

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	offset := (page - 1) * limit

	if err := config.DB.Limit(limit).Offset(offset).Find(&rees).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to fetch REEs")
		return
	}

	utils.Success(c, rees)
}

func CreateREE(c *gin.Context) {
	var req dto.CreateREEDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	ree := models.REE{
		Name:        req.Name,
		Symbol:      req.Symbol,
		Category:    req.Category,
		MarketPrice: req.MarketPrice,
	}

	if err := config.DB.Create(&ree).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to create REE")
		return
	}

	utils.Success(c, ree)
}

func UpdateREE(c *gin.Context) {
	var req dto.UpdateREEDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var ree models.REE
	if err := config.DB.First(&ree, c.Param("id")).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "REE not found")
		return
	}

	if err := config.DB.Model(&ree).Updates(req).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Update failed")
		return
	}

	utils.Success(c, ree)
}

func DeleteREE(c *gin.Context) {
	if err := config.DB.Delete(&models.REE{}, c.Param("id")).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Delete failed")
		return
	}

	utils.Success(c, gin.H{"message": "REE deleted"})
}
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
