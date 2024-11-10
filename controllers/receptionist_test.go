package controllers

import (
	"bytes"
	"encoding/json"
	"markeable/database"
	"markeable/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreatePatient(t *testing.T) {
	database.InitMockDB()
	app := fiber.New()
	app.Post("/patients", CreatePatient)

	payload := models.Patient{
		FirstName: "vishal",
		LastName:  "keshri",
		Age:       24,
		DoctorID:  1,
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/patients", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var patient models.Patient
	json.NewDecoder(resp.Body).Decode(&patient)

	assert.Equal(t, "vishal", patient.FirstName)
	assert.Equal(t, "keshri", patient.LastName)
	assert.Equal(t, 24, patient.Age)
	assert.Equal(t, uint(1), patient.DoctorID)
}

func TestGetPatients(t *testing.T) {
	database.InitMockDB()
	app := fiber.New()
	app.Get("/patients", GetPatients)

	database.DB.Create(&models.Patient{FirstName: "vishal", LastName: "keshri", Age: 24, DoctorID: 1})

	req := httptest.NewRequest("GET", "/patients", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var patients []models.Patient
	json.NewDecoder(resp.Body).Decode(&patients)

	assert.Equal(t, 3, len(patients))
	assert.Equal(t, "vishal", patients[0].FirstName)

}

func TestUpdatePatient(t *testing.T) {
	database.InitMockDB()
	app := fiber.New()
	app.Put("/patients/:id", UpdatePatient)

	patient := models.Patient{FirstName: "vishal", LastName: "keshri", Age: 24, DoctorID: 1}
	database.DB.Create(&patient)

	updatedData := models.Patient{
		FirstName: "vishal Updated",
		LastName:  "keshri Updated",
		Age:       31,
		DoctorID:  1,
	}

	body, _ := json.Marshal(updatedData)
	req := httptest.NewRequest("PUT", "/patients/"+strconv.Itoa(int(patient.ID)), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var updatedPatient models.Patient
	json.NewDecoder(resp.Body).Decode(&updatedPatient)

	assert.Equal(t, "vishal Updated", updatedPatient.FirstName)
	assert.Equal(t, "keshri Updated", updatedPatient.LastName)
	assert.Equal(t, 31, updatedPatient.Age)
	assert.Equal(t, uint(1), updatedPatient.DoctorID)
}

func TestDeletePatient(t *testing.T) {
	database.InitMockDB()
	app := fiber.New()
	app.Delete("/patients/:id", DeletePatient)

	patient := models.Patient{FirstName: "vishal", LastName: "keshri", Age: 24, DoctorID: 1}
	database.DB.Create(&patient)

	req := httptest.NewRequest("DELETE", "/patients/"+strconv.Itoa(int(patient.ID)), nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response map[string]string
	json.NewDecoder(resp.Body).Decode(&response)

	assert.Equal(t, "Patient deleted", response["message"])

	var deletedPatient models.Patient
	result := database.DB.First(&deletedPatient, patient.ID)
	assert.Error(t, result.Error)
	assert.Equal(t, uint(0), deletedPatient.ID)
}
