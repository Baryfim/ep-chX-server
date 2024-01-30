package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sixfwa/fiber-api/database"
	"github.com/sixfwa/fiber-api/middlewares"
	"github.com/sixfwa/fiber-api/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}

func setupRoutes(app *fiber.App) {
	// welcome
	app.Get("/api", welcome)
	// Year endpoints
	app.Post("/api/years", middlewares.CheackAdminIsValid, routes.CreateYear)
	app.Get("/api/years", middlewares.CheackAdminIsValid, routes.GetYears)
	app.Get("/api/years/:id", routes.GetYear)
	app.Put("/api/years/:id", middlewares.CheackAdminIsValid, routes.UpdateYear)
	app.Delete("/api/years/:id", middlewares.CheackAdminIsValid, routes.DeleteYear)
	// Item endpoints
	app.Post("/api/items", middlewares.CheackAdminIsValid, routes.CreateItem)
	app.Get("/api/items", routes.GetItems)
	app.Get("/api/items/:id", routes.GetItem)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error dotenv file \n" + err.Error())
		return
	}

	log.Fatal(app.Listen(os.Getenv("PORT")))
}
