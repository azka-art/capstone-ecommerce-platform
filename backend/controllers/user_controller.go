package controllers

import (
	"ecommerce-platform/backend/config"
	"ecommerce-platform/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers retrieves all users from the database
func GetUsers(c *gin.Context) {
	var users []models.User
	if result := config.DB.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID retrieves a user by their ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if result := config.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser updates user details by ID
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if result := config.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Username = updatedUser.Username
	user.Password = updatedUser.Password // Note: In a real application, you should handle password hashing separately
	if result := config.DB.Save(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user by their ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if result := config.DB.Delete(&models.User{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
