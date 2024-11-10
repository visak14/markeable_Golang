package controllers

import (
	"encoding/json"
	"markeable/database"
	"markeable/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAllPatients(t *testing.T) {

	database.InitMockDB()
	app := fiber.New()
	app.Get("/patients", GetAllPatients)

	database.DB.Create(&models.Patient{FirstName: "vishal", LastName: "keshri", Age: 24, DoctorID: 1})

	req := httptest.NewRequest("GET", "/patients", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var patients []models.Patient
	json.NewDecoder(resp.Body).Decode(&patients)

	assert.Equal(t, 1, len(patients))

	assert.Equal(t, "vishal", patients[0].FirstName)
	assert.Equal(t, "keshri", patients[0].LastName)
	assert.Equal(t, 24, patients[0].Age)
	assert.Equal(t, uint(1), patients[0].DoctorID)
}
