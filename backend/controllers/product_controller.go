package controllers

import (
	"ecommerce-platform/backend/config"
	"ecommerce-platform/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProducts retrieves all products
func GetProducts(c *gin.Context) {
	var products []models.Product
	if result := config.DB.Find(&products); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// CreateProduct adds a new product
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := config.DB.Create(&product); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
}
