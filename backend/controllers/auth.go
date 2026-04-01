package controllers

import (
	"net/http"

<<<<<<< HEAD

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/imperionite/reeisto/backend/dto"
	"github.com/imperionite/reeisto/backend/models"
	"github.com/imperionite/reeisto/backend/utils"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.LoginDTO
		if err := c.ShouldBindJSON(&input); err != nil {
			utils.Error(c, http.StatusBadRequest, "Invalid input")
			return
		}

		var user models.User
		if err := db.Where("username = ? AND deleted_at IS NULL", input.Username).
			First(&user).Error; err != nil {
			utils.Error(c, http.StatusUnauthorized, "Invalid credentials")
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			utils.Error(c, http.StatusUnauthorized, "Invalid credentials")
			return
		}

		token, err := utils.GenerateJWT(user.ID, user.Role)
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, "Failed to generate token")
			return
		}

		utils.Success(c, gin.H{
			"token":  token,
			"user":   gin.H{"id": user.ID, "username": user.Username, "role": user.Role},
			"expires_in": "1h",
		})
	}
}

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.RegisterDTO
		if err := c.ShouldBindJSON(&input); err != nil {
			utils.Error(c, http.StatusBadRequest, "Invalid input")
			return
		}

		// Check if user exists
		var existing models.User
		if err := db.Where("username = ? OR email = ?", input.Username, input.Email).First(&existing).Error; err == nil {
			utils.Error(c, http.StatusConflict, "Username or email already exists")
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, "Failed to create user")
			return
		}

		user := models.User{
			Username: input.Username,
			Email:    input.Email,
			Password: string(hashedPassword),
			Role:     "trader",
		}

		if err := db.Create(&user).Error; err != nil {
			utils.Error(c, http.StatusInternalServerError, "Failed to create user")
			return
		}

		token, err := utils.GenerateJWT(user.ID, user.Role)
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, "Failed to generate token")
			return
		}


		utils.Success(c, gin.H{
			"token": token,
			"user":  gin.H{"id": user.ID, "username": user.Username, "role": user.Role},
		})
	}
}

=======
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"

	"github.com/imperionite/reetis/backend/config"
	"github.com/imperionite/reetis/backend/dto"
	"github.com/imperionite/reetis/backend/models"
	"github.com/imperionite/reetis/backend/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req dto.LoginDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	db := config.DB
	var user models.User

	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.Error(c, http.StatusUnauthorized, "Invalid credentials")
			return
		}
		utils.Error(c, http.StatusInternalServerError, "Database error")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		utils.Error(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.Success(c, gin.H{
		"token": token,
	})
}
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
