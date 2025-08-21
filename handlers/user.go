package handlers

import (
	"booking-bengkel-smkn1depok/database"
	"booking-bengkel-smkn1depok/models"

	"github.com/gofiber/fiber/v2"
)

// CREATE user
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// simpan ke DB
	if result := database.DB.Create(&user); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(201).JSON(user)
}

// READ all users
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

// READ single user
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	if result := database.DB.First(&user, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

// UPDATE user
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	if result := database.DB.First(&user, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	database.DB.Save(&user)
	return c.JSON(user)
}

// DELETE user
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	if result := database.DB.First(&user, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	database.DB.Delete(&user)
	return c.JSON(fiber.Map{"message": "User deleted"})
}
