package middleware

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthorizeRole(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		cookie := c.Cookies("jwt")

		token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			fmt.Println(os.Getenv("JWT_SECRET"))
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			log.Println("JWT Parse Error:", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		// Extract role from claims
		claims := token.Claims.(*jwt.MapClaims)
		userRole, ok := (*claims)["role"].(string)
		if !ok {
			log.Println("Role not found in claims")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized, Role missing",
			})
		}

		for _, role := range roles {
			if userRole == role {
				c.Locals("role", userRole)
				return c.Next()
			}
		}

		log.Println("Access denied: role mismatch")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Access denied",
		})
	}
}
