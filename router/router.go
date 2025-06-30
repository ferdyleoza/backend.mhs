package router

import (
	"inibackend/config/middleware"
	"inibackend/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api",)

	// Route untuk homepage
	api.Get("/", handler.Homepage)

	api.Get("/mahasiswa", handler.GetAllMahasiswa)

	api.Get("/mahasiswa/:npm",  middleware.Middlewares("admin"), handler.GetMahasiswaByNPM)

	// Route untuk menambah mahasiswa baru
	api.Post("/mahasiswa",  middleware.Middlewares("admin"),handler.CreateMahasiswa)

	// Route untuk mengupdate data mahasiswa berdasarkan NPM
	api.Put("/mahasiswa/:npm",  middleware.Middlewares("admin") ,handler.UpdateMahasiswa)

	// Route untuk menghapus data mahasiswa berdasarkan NPM
	api.Delete("/mahasiswa/:npm",  middleware.Middlewares("admin"),handler.DeleteMahasiswa)

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

}