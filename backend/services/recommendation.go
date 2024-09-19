package services

import (
	"ecommerce-platform/backend/config"
	"ecommerce-platform/backend/models"
)

func GetRecommendations(userID uint) []models.Product {
	var user models.User
	var recommendations []models.Product

	config.DB.Preload("Products").First(&user, userID)

	// Simple recommendation based on previous purchases
	for _, product := range user.Products {
		var relatedProducts []models.Product
		config.DB.Where("id != ?", product.ID).Find(&relatedProducts)
		recommendations = append(recommendations, relatedProducts...)
	}

	return recommendations
}
