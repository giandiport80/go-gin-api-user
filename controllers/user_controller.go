package controllers

import (
	"net/http"

	"github.com/giandiport80/go-gin-api-user/config"
	"github.com/giandiport80/go-gin-api-user/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
	})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Data tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Terjadi kesalahan",
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Data berhasil disimpan",
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":  false,
			"messaage": err.Error(),
		})
		return
	}

	if err := config.DB.Model(&user).Updates(models.User{
		Name:  input.Name,
		Email: input.Email,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data berhasil disimpan",
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	}

	config.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data berhasil dihapus",
	})
}
