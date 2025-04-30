package main

// first task is tto do jwt authentication

import (
	"Chat-application/database"
	"Chat-application/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	_, err := database.Connect()
	if err != nil {
		panic("their is some error in connecting the database")
	}
	//connected the database
	fmt.Println("the connetion was successfull")

	app := fiber.New()
	// why cors-cross origin resourse sharing ---- API communication ,Blocks requests from different origins (domain, protocol, or port)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	//setting up routes
	routes.SetupRoutes(app)

	app.Listen(":8080")

}
