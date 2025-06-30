package handler

import (
	"inibackend/model"
	"inibackend/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	return c.SendString("Welcome to the jungle!")
}

func GetAllMahasiswa(c *fiber.Ctx) error {
	// Call the repository function to get all Mahasiswa
	mahasiswa, err := repository.GetAllMahasiswa(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": "Failed to retrieve data",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Data retrieved successfully",
		"data":    mahasiswa,
	})
}

func GetMahasiswaByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm")
	if npm == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "NPM is required",
		})
	}

	npmInt, err := strconv.Atoi(npm)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "NPM must be a valid number",
		})
	}

	mhs, err := repository.GetMahasiswaByNPM(c.Context(), npmInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": "Failed to retrieve Mahasiswa",
		})
	}
	if mhs.NPM == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "Mahasiswa not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Data retrieved successfully",
		"data":    mhs,
	})
}

func CreateMahasiswa(c *fiber.Ctx) error {
	var mahasiswaData model.Mahasiswa
	if err := c.BodyParser(&mahasiswaData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	insertedID, err := repository.InsertMahasiswa(c.Context(), mahasiswaData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "Mahasiswa created successfully",
		"id":      insertedID,
	})
}

func UpdateMahasiswa(c *fiber.Ctx) error {
	npm := c.Params("npm")
	npmInt, err := strconv.Atoi(npm)
	if err != nil || npmInt <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "Invalid NPM parameter. It must be a positive integer.",
		})
	}

	var mahasiswaData model.Mahasiswa
	if err := c.BodyParser(&mahasiswaData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "Invalid request body format",
		})
	}

	mahasiswaData.NPM = npmInt

	// Validate required fields
	if mahasiswaData.Nama == "" || mahasiswaData.Prodi == "" || mahasiswaData.Fakultas == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "Nama, Prodi, and Fakultas are required fields",
		})
	}

	updatedCount, err := repository.UpdateMahasiswa(c.Context(), npmInt, mahasiswaData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": "Failed to update Mahasiswa",
		})
	}
	if updatedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "Mahasiswa not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Mahasiswa updated successfully",
	})
}


func DeleteMahasiswa(c *fiber.Ctx) error {
	npm := c.Params("npm")
	npmInt, err := strconv.Atoi(npm)
	if err != nil || npmInt <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "Invalid NPM parameter. It must be a positive integer.",
		})
	}

	success, err := repository.DeleteMahasiswa(c.Context(), npmInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": "Failed to delete Mahasiswa",
		})
	}
	if success == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "Mahasiswa not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Mahasiswa deleted successfully",
	})
}