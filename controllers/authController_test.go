package controllers

import (
	"bytes"
	"encoding/json"
	"markeable/database"
	"markeable/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func mockDB() {
	database.DB = database.InitMockDB()
	database.DB.AutoMigrate(&models.User{})
}

func TestRegister(t *testing.T) {
	mockDB()

	app := fiber.New()

	app.Post("/register", Register)

	testData := map[string]string{
		"username": "testuser",
		"email":    "test@example.com",
		"password": "password123",
		"role":     "user",
	}
	payload, _ := json.Marshal(testData)

	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var user models.User
	json.NewDecoder(resp.Body).Decode(&user)
	assert.Equal(t, testData["username"], user.Username)
	assert.Equal(t, testData["email"], user.Email)
	assert.Equal(t, testData["role"], user.Role)

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(testData["password"]))
	assert.Nil(t, err)
}

func TestLogin(t *testing.T) {

	database.InitMockDB()

	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), 14)
	user := models.User{
		Username: "vishal",
		Email:    "vishal@example.com",
		Password: password,
		Role:     "user",
	}
	database.DB.Create(&user)

	app := fiber.New()

	app.Post("/login", Login)

	payload := map[string]string{
		"email":    "vishal@example.com",
		"password": "password123",
	}

	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response map[string]string
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, "success", response["message"])

	cookie := resp.Header.Get("Set-Cookie")
	assert.Contains(t, cookie, "jwt=")
}
