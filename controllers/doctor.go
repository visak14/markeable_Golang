package controllers

import (
	"markeable/database"
	"markeable/models"

	"github.com/gofiber/fiber/v2"
)

// CreatePatientForDoctor  Create a new patient (Doctor Role)
//
//	@Summary		Create a new patient record (Doctor Role)
//	@Description	This endpoint allows a doctor to create a new patient record by providing details like first name, last name, age, doctor_id, etc.
//	@Tags			Patients
//	@Accept			json
//	@Produce		json
//	@Param			patient	body		models.Patient		true	"Patient data"
//	@Success		201		{object}	models.Patient		"Successfully created patient"
//	@Failure		400		{object}	map[string]string	"Invalid input"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/doctor/patient [post]
func CreatePatientForDoctor(c *fiber.Ctx) error {
	var patient models.Patient

	if err := c.BodyParser(&patient); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input"})
	}

	if err := database.DB.Create(&patient).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error creating patient"})
	}

	return c.Status(fiber.StatusCreated).JSON(patient)
}

// UpdatePatientForDoctor  Update an existing patient by ID (Doctor Role)
//
//	@Summary		Update an existing patient record by ID (Doctor Role)
//	@Description	This endpoint allows a doctor to update an existing patient record. You need to provide the patient's ID in the URL and the new details in the request body.
//	@Tags			Patients
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Patient ID"
//	@Param			patient	body		models.Patient		true	"Updated patient data"
//	@Success		200		{object}	models.Patient		"Successfully updated patient"
//	@Failure		400		{object}	map[string]string	"Invalid input"
//	@Failure		404		{object}	map[string]string	"Patient not found"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/doctor/patient/{id} [put]
func UpdatePatientForDoctor(c *fiber.Ctx) error {
	id := c.Params("id")
	var patient models.Patient

	if err := database.DB.First(&patient, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Patient not found"})
	}

	var updatedPatient models.Patient
	if err := c.BodyParser(&updatedPatient); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input"})
	}

	patient.FirstName = updatedPatient.FirstName
	patient.LastName = updatedPatient.LastName
	patient.Age = updatedPatient.Age
	patient.DoctorID = updatedPatient.DoctorID

	if err := database.DB.Save(&patient).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error updating patient"})
	}

	return c.JSON(patient)
}

// DeletePatientForDoctor  Delete a patient by ID (Doctor Role)
//
//	@Summary		Delete a patient record by ID (Doctor Role)
//	@Description	This endpoint allows a doctor to delete a patient record by providing the patient's ID in the URL.
//	@Tags			Patients
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int					true	"Patient ID"
//	@Success		200	{object}	map[string]string	"Successfully deleted patient"
//	@Failure		404	{object}	map[string]string	"Patient not found"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/doctor/patient/{id} [delete]
func DeletePatientForDoctor(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := database.DB.Delete(&models.Patient{}, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Patient not found"})
	}

	return c.JSON(fiber.Map{"message": "Patient deleted"})
}

// GetAllPatients  Get all patients
//
//	@Summary		Retrieve a list of all patients
//	@Description	This endpoint retrieves all patient records from the database and returns them as a JSON array. This operation does not require authentication.
//	@Tags			Patients
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Patient		"Successfully retrieved all patients"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/doctor/patients [get]
func GetAllPatients(c *fiber.Ctx) error {
	var patients []models.Patient
	database.DB.Find(&patients)
	return c.JSON(patients)
}
