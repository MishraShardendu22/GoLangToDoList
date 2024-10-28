package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ShardenduMishra22/GoLangToDoList/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("ToDo List Project!!")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectToDatabase()
	app := fiber.New()

	// Just a sample import file
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message1": "Welcome to ToDo List Project",
			"message2": "This is a Sample Response to test if the application",
		})
	})

	// Api calls to perform CRUD operations	

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	log.Fatal(app.Listen(":" + port))
}
