package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

<<<<<<< HEAD
	"github.com/imperionite/reeisto/backend/dto"
	"github.com/imperionite/reeisto/backend/services"
	"github.com/imperionite/reeisto/backend/utils"
)

func GetTransactions(txService *services.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get userID safely as uint
		userIDVal, exists := c.Get("userID")
		if !exists {
			utils.Error(c, http.StatusUnauthorized, "Unauthorized")
			return
		}

		userID, ok := userIDVal.(uint)
		if !ok {
			utils.Error(c, http.StatusInternalServerError, "Invalid user context")
			return
		}

		// Pagination
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
		offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

		transactions, err := txService.List(userID, limit, offset)
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, "Failed to fetch transactions")
			return
		}

		utils.Success(c, gin.H{
			"data": transactions,
			"meta": gin.H{
				"limit":  limit,
				"offset": offset,
				"count":  len(transactions),
			},
		})
	}
}

func CreateTransaction(txService *services.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get userID safely as uint
		userIDVal, exists := c.Get("userID")
		if !exists {
			utils.Error(c, http.StatusUnauthorized, "Unauthorized")
			return
		}

		userID, ok := userIDVal.(uint)
		if !ok {
			utils.Error(c, http.StatusInternalServerError, "Invalid user context")
			return
		}

		var input dto.CreateTransactionDTO
		if err := c.ShouldBindJSON(&input); err != nil {
			utils.Error(c, http.StatusBadRequest, "Invalid input")
			return
		}

		transaction, err := txService.Create(userID, &input)
		if err != nil {
			utils.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		utils.Success(c, transaction)
	}
}
=======
	"github.com/imperionite/reetis/backend/config"
	"github.com/imperionite/reetis/backend/dto"
	"github.com/imperionite/reetis/backend/models"
	"github.com/imperionite/reetis/backend/services"
	"github.com/imperionite/reetis/backend/utils"
)

var txService = services.NewTransactionService(config.DB)

func GetTransactions(c *gin.Context) {
	db := config.DB
	var txs []models.Transaction

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	offset := (page - 1) * limit

	txType := c.Query("type")
	userID := c.GetUint("userID")
	role := c.GetString("role")

	query := db.Preload("Element").Preload("User").
		Limit(limit).Offset(offset)

	if role != "admin" {
		query = query.Where("user_id = ?", userID)
	}

	if txType != "" {
		query = query.Where("type = ?", txType)
	}

	if err := query.Find(&txs).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to fetch transactions")
		return
	}

	utils.Success(c, txs)
}

func CreateTransaction(c *gin.Context) {
	var req dto.CreateTransactionDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	tx := models.Transaction{
		ElementID: req.ElementID,
		Type:      req.Type,
		Quantity:  req.Quantity,
		Price:     req.Price,
		UserID:    c.GetUint("userID"),
	}

	if err := txService.Create(&tx); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, tx)
}
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
