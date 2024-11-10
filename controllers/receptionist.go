package controllers

import (
	"markeable/database"
	"markeable/models"

	"github.com/gofiber/fiber/v2"
)

// CreatePatient  Create a new patient
//
//	@Summary		Create a new patient record
//	@Description	This endpoint allows the creation of a new patient record by providing details like first name, last name, email, phone number, etc.
//	@Tags			Patients
//	@Accept			json
//	@Produce		json
//	@Param			patient	body		models.Patient		true	"Patient data"
//	@Success		201		{object}	models.Patient		"Successfully created patient"
//	@Failure		400		{object}	map[string]string	"Invalid input"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/receptionist/patient [post]
func CreatePatient(c *fiber.Ctx) error {
	var patient models.Patient
	if err := c.BodyParser(&patient); err != nil {
		return err
	}
	database.DB.Create(&patient)
	return c.JSON(patient)
}

// GetPatients  Retrieve all patients
//
//	@Summary		Get all patient records
//	@Description	This endpoint retrieves all patient records from the database and returns them as a JSON array.
//	@Tags			Patients
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Patient		"Successfully retrieved all patients"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/receptionist/patients [get]
func GetPatients(c *fiber.Ctx) error {
	var patients []models.Patient
	database.DB.Find(&patients)
	return c.JSON(patients)
}

// UpdatePatient  Update an existing patient by ID
//
//	@Summary		Update an existing patient record by ID
//	@Description	This endpoint allows you to update an existing patient record. You need to provide the patient's ID in the URL and the new details in the request body.
//	@Tags			Patients
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Patient ID"
//	@Param			patient	body		models.Patient		true	"Updated patient data"
//	@Success		200		{object}	models.Patient		"Successfully updated patient"
//	@Failure		400		{object}	map[string]string	"Invalid input"
//	@Failure		404		{object}	map[string]string	"Patient not found"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/receptionist/patient/{id} [put]
func UpdatePatient(c *fiber.Ctx) error {

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

// DeletePatient  Delete a patient by ID
//
//	@Summary		Delete a patient record by ID
//	@Description	This endpoint allows you to delete a patient record by providing the patient's ID in the URL.
//	@Tags			Patients
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int					true	"Patient ID"
//	@Success		200	{object}	map[string]string	"Successfully deleted patient"
//	@Failure		404	{object}	map[string]string	"Patient not found"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/receptionist/patient/{id} [delete]
func DeletePatient(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.Patient{}, id)
	return c.JSON(fiber.Map{"message": "Patient deleted"})
}
