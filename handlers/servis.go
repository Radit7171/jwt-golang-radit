package handlers

import (
    "booking-bengkel-smkn1depok/config"
    "booking-bengkel-smkn1depok/models"
    "github.com/gofiber/fiber/v2"
)

// Create Servis
func CreateServis(c *fiber.Ctx) error {
    var servis models.Servis
    if err := c.BodyParser(&servis); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }
    config.DB.Create(&servis)
    return c.JSON(servis)
}

// Get All Servis
func GetServis(c *fiber.Ctx) error {
    var servis []models.Servis
    config.DB.Find(&servis)
    return c.JSON(servis)
}

// Get Servis by ID
func GetServisByID(c *fiber.Ctx) error {
    id := c.Params("id")
    var servis models.Servis
    if result := config.DB.First(&servis, id); result.Error != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Servis tidak ditemukan"})
    }
    return c.JSON(servis)
}

// Update Servis
func UpdateServis(c *fiber.Ctx) error {
    id := c.Params("id")
    var servis models.Servis
    if result := config.DB.First(&servis, id); result.Error != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Servis tidak ditemukan"})
    }
    if err := c.BodyParser(&servis); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }
    config.DB.Save(&servis)
    return c.JSON(servis)
}

// Delete Servis
func DeleteServis(c *fiber.Ctx) error {
    id := c.Params("id")
    var servis models.Servis
    if result := config.DB.Delete(&servis, id); result.Error != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Servis tidak ditemukan"})
    }
    return c.JSON(fiber.Map{"message": "Servis dihapus"})
}
