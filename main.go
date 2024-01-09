package main

import (
	"GikBazar/database"
	"GikBazar/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	dbRef, err := database.ConnectDB()
	if err != nil {
		println("Error in 1 db connection")
	} else {
		println("Connected sucsesfull")
	}
	db, err := dbRef.DB()
	if err != nil {
		println("Error in 2 db connection")
	} else {
		println("Connected sucsesfull")
	}
	defer db.Close()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	router.SetupRoutes(app)
	app.Listen(":8080")
}
