package controllers

import (
	"fmt"
	"markeable/database"
	"markeable/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hi, I am Vishal keshri!")
}

// Register      Register user by Code
//
//	@Summary		Register a new user by provided data (username, email, password, role)
//	@Description	This endpoint allows you to register a new user by providing their username, email, password, and role. The password is hashed before storing it in the database.
//	@Tags			Register
//	@Accept			json
//	@Produce		json
//	@Param			user	body		map[string]string	true	"User registration details"
//	@Success		200		{object}	models.User			"Successfully created user"
//	@Failure		400		{object}	map[string]string	"Invalid input"
//	@Failure		409		{object}	map[string]string	"Email already exists"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/register [post]
func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(map[string]string{"message": "Invalid input"})
	}

	if data["username"] == "" || data["email"] == "" || data["password"] == "" || data["role"] == "" {
		return c.Status(400).JSON(map[string]string{"message": "All fields (username, email, password, role) are required"})
	}

	var existingUser models.User
	database.DB.Where("email = ?", data["email"]).First(&existingUser)
	if existingUser.Id != 0 {
		return c.Status(409).JSON(map[string]string{"message": "Email already exists"})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Username: data["username"],
		Email:    data["email"],
		Password: password,
		Role:     data["role"],
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.Status(500)
		return c.JSON(map[string]string{"message": "Internal server error"})
	}

	return c.JSON(user)
}

// Login         Login user by email and password
//
//	@Summary		Log in a user using email and password
//	@Description	This endpoint allows a user to log in by providing their email and password. If the credentials are correct, a JWT token is generated and sent as a cookie in the response.
//	@Tags			Login
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		map[string]string	true	"User login credentials"
//	@Success		200			{object}	map[string]string	"Login successful, JWT token set in cookie"
//	@Failure		400			{object}	map[string]string	"Incorrect password"
//	@Failure		404			{object}	map[string]string	"User not found"
//	@Failure		500			{object}	map[string]string	"Internal server error"
//	@Router			/login [post]
func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"issuer": user.Id,
		"role":   user.Role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	fmt.Println(os.Getenv("JWT_SECRET"))
	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
