package handler

import (
	"inibackend/model"
	"inibackend/repository"
	"strconv"

	 _"inibackend/docs"
	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	return c.SendString("Welcome to the jungle!")
}

// GetAllMahasiswa godoc
// @Summary Get All Data Mahasiswa.
// @Description Mengambil semua data mahasiswa.
// @Tags Mahasiswa
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} model.Mahasiswa
// @Failure 401 "Unauthorized"
// @Failure 500
// @Router /api/mahasiswa [get]
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
// GetMahasiswaByNPM godoc
// @Summary Get By NPM Data Mahasiswa.
// @Description Ambil per NPM data mahasiswa.
// @Tags Mahasiswa
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param npm path int true "Masukan NPM"
// @Success 200 {object} model.Mahasiswa
// @Failure 400 "NPM harus berupa angka"
// @Failure 401 "Unauthorized"
// @Failure 404 "Data tidak ditemukan"
// @Router /api/mahasiswa/{npm} [get]
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
// InsertMahasiswa godoc
// @Summary Insert data Mahasiswa.
// @Description Input data Mahasiswa.
// @Tags Mahasiswa
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body model.MahasiswaRequest true "Payload Body [RAW]"
// @Success 201 {object} model.Mahasiswa
// @Failure 400 "Invalid request data"
// @Failure 401 "Unauthorized"
// @Failure 409 "Gagal menambahkan mahasiswa"
// @Router /api/mahasiswa [post]
func InsertMahasiswa(c *fiber.Ctx) error {
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
// UpdateMahasiswa godoc
// @Summary Update data Mahasiswa.
// @Description Ubah data Mahasiswa.
// @Tags Mahasiswa
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param npm path integer true "Masukan NPM"
// @Param request body model.MahasiswaRequest true "Payload Body [RAW]"
// @Success 200 {object} model.Mahasiswa
// @Failure 400 "Invalid request data or Invalid NPM format"
// @Failure 401 "Unauthorized"
// @Failure 404 "Error Update Data Mahasiswa"
// @Router /api/mahasiswa/{npm} [put]
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

// DeleteMahasiswa godoc
// @Summary Delete data Mahasiswa.
// @Description Hapus data Mahasiswa.
// @Tags Mahasiswa
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param npm path integer true "Masukan NPM"
// @Success 200 "Mahasiswa berhasil dihapus"
// @Failure 400 "Invalid NPM format"
// @Failure 401 "Unauthorized"
// @Failure 404 "Mahasiswa tidak ditemukan"
// @Router /api/mahasiswa/{npm} [delete]
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