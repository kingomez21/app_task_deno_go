package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/kingomez21/app_task_deno_go/database"
	"github.com/kingomez21/app_task_deno_go/routers"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error al cargar el .env")
	}

	database.GetConectionDB()

	app := fiber.New()

	app.Use(cors.New())

	routers.UserRoute(app)

	app.Listen(":4000")
}
