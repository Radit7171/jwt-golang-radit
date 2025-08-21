package handlers

import (
	"booking-bengkel-smkn1depok/database"
	"booking-bengkel-smkn1depok/models"

	"github.com/gofiber/fiber/v2"
)

// Create Teknisi
func CreateTeknisi(c *fiber.Ctx) error {
	teknisi := new(models.Teknisi)

	if err := c.BodyParser(teknisi); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	if err := database.DB.Create(&teknisi).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(teknisi)
}

// Get All Teknisi
func GetTeknisi(c *fiber.Ctx) error {
	var teknisi []models.Teknisi
	if err := database.DB.Find(&teknisi).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(teknisi)
}

// Get Teknisi By ID
func GetTeknisiByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var teknisi models.Teknisi
	if err := database.DB.First(&teknisi, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "teknisi not found",
		})
	}
	return c.JSON(teknisi)
}

// Update Teknisi
func UpdateTeknisi(c *fiber.Ctx) error {
	id := c.Params("id")
	var teknisi models.Teknisi

	// cek data lama
	if err := database.DB.First(&teknisi, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "teknisi not found",
		})
	}

	// parsing update data
	if err := c.BodyParser(&teknisi); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	if err := database.DB.Save(&teknisi).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(teknisi)
}

// Delete Teknisi
func DeleteTeknisi(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Teknisi{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "teknisi deleted successfully",
	})
}
