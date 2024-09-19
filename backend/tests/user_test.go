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

func TestRegisterUser(t *testing.T) {
	r := setupRouter()

	newUser := models.User{
		Username: "testuser",
		Password: "password123",
	}

	body, _ := json.Marshal(newUser)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
