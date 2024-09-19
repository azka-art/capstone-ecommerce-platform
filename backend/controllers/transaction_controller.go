package controllers

import (
	"ecommerce-platform/backend/config"
	"ecommerce-platform/backend/models"
	"net/http"

	"ecommerce-platform/backend/services"

	"github.com/gin-gonic/gin"
)

// CreateTransaction records a new transaction
func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := config.DB.Create(&transaction); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction created successfully"})
}

func GetCartItems(c *gin.Context) {
	username, _ := c.Get("username")

	cartItems, err := services.GetCartItemsByUsername(username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart items"})
		return
	}

	c.JSON(http.StatusOK, cartItems)
}

func AddToCart(c *gin.Context) {
	username, _ := c.Get("username")
	var requestBody struct {
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := services.AddToCart(username.(string), requestBody.ProductID, requestBody.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

func RemoveFromCart(c *gin.Context) {
	username, _ := c.Get("username")
	itemID := c.Param("id")

	err := services.RemoveFromCart(username.(string), itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item from cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}
