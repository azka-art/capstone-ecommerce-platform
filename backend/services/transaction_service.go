// backend/services/transaction_service.go
package services

import (
	"ecommerce-platform/backend/config"
	"ecommerce-platform/backend/models"

	"gorm.io/gorm"
)

var db *gorm.DB = config.DB

// GetCartItemsByUsername retrieves cart items for a user
func GetCartItemsByUsername(username string) ([]models.Product, error) {
	var user models.User
	err := db.Where("username = ?", username).Preload("Products").First(&user).Error
	if err != nil {
		return nil, err
	}
	return user.Products, nil
}

// backend/services/transaction_service.go
// AddToCart adds a product to the user's cart
func AddToCart(username, productID string, quantity int) error {
	var user models.User
	var product models.Product

	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return err
	}

	err = db.Where("id = ?", productID).First(&product).Error
	if err != nil {
		return err
	}

	transaction := models.Transaction{
		UserID:    user.ID,
		ProductID: product.ID,
		Quantity:  quantity,
	}

	return db.Create(&transaction).Error
}

// backend/services/transaction_service.go
// RemoveFromCart removes a product from the user's cart
func RemoveFromCart(username, itemID string) error {
	var transaction models.Transaction

	err := db.Where("id = ? AND user_id = (SELECT id FROM users WHERE username = ?)", itemID, username).Delete(&transaction).Error
	if err != nil {
		return err
	}

	return nil
}
