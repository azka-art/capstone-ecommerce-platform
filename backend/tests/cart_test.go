package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ecommerce-platform/backend/models"

	"github.com/stretchr/testify/assert"
)

func TestAddToCart(t *testing.T) {
	r := setupRouter()

	cartItem := models.Transaction{
		UserID:    1,
		ProductID: 1,
		Quantity:  1,
	}

	body, _ := json.Marshal(cartItem)

	req, _ := http.NewRequest("POST", "/cart", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
