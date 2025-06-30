package main

import (
	"fmt"
	"inibackend/config"
	"inibackend/router"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)	

func init() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Gagal memuat file .env:", err)
		} else {
			fmt.Println("File .env berhasil dimuat")
		}
	}
}


// @title TES SWAGGER PEMROGRAMAN III
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/ferdyleoza
// @contact.email ferdyleoza123@gmail.com

// @BasePath /
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New()

	//logging request di terminal
	app.Use(logger.New())

	// basic middleware cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join(config.GetALLowedOrigins(), ","),
		AllowCredentials: true,
		AllowMethods: "GET,POST,PUT,DELETE",
}))

	//Route mahasiswa
	router.SetupRoutes(app)

	//handler 404
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": fiber.StatusNotFound,
			"message": "Endpoint Tidak Ditemukan",
		})
	})

	//Baca port yang ada di env
	port :=os.Getenv("PORT")
	if port == "" {
		port = "3000"//default port kalau tidak ada di env
	}
	
	//untuk log cek konek di port mana
	log.Printf("Server running on port %s\n", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	
	//koneksi terputus
}

