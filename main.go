package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ShardenduMishra22/GoLangToDoList/database"
	"github.com/ShardenduMishra22/GoLangToDoList/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection // Declare a global variable to hold the collection

func main() {
	fmt.Println("ToDo List Project!!")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database and get the collection
	collection = database.ConnectToDatabase() // Store the collection

	app := fiber.New()
	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // Replace with your frontend's origin
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message1": "Welcome to ToDo List Project",
			"message2": "This is a Sample Response to test if the application",
		})
	})

	routes.SetupRoutes(app, collection) // Pass the collection to the routes

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}

// Sample Code for GoLang Project
// package main

// import (
// 	"fmt"
// 	"log"
// 	"strconv"

// 	"github.com/gofiber/fiber/v2"
// )

// type Todo struct {
// 	ID        int    `json:"id"`
// 	Body      string `json:"body"`
// 	Completed bool   `json:"completed"`
// }

// func main() {
// 	fmt.Println("Hi, This is my first GoLang Project")

// 	app := fiber.New()
// 	todos := make([]Todo, 0)

// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.JSON(fiber.Map{
// 			"message": "This is a Sample Application",
// 		})
// 	})

// 	app.Get("/todo", func(c *fiber.Ctx) error {
// 		return c.JSON(fiber.Map{
// 			"todos": todos,
// 		})
// 	})

// 	app.Post("/todo", func(c *fiber.Ctx) error {
// 		todo := new(Todo)
// 		if err := c.BodyParser(todo); err != nil {
// 			return c.Status(400).JSON(fiber.Map{
// 				"error": "Failed to parse body",
// 			})
// 		}

// 		if todo.Body == "" {
// 			return c.Status(400).JSON(fiber.Map{
// 				"error": "Body is required",
// 			})
// 		}

// 		todo.ID = len(todos) + 1
// 		todos = append(todos, *todo)

// 		return c.Status(201).JSON(todo)
// 	})

// 	app.Patch("/todo/:id", func(c *fiber.Ctx) error {
// 		idParam := c.Params("id")
// 		id, err := strconv.Atoi(idParam)
// 		if err != nil {
// 			return c.Status(400).JSON(fiber.Map{
// 				"error": "Invalid ID",
// 			})
// 		}

// 		var todo *Todo
// 		for i := range todos {
// 			if todos[i].ID == id {
// 				todo = &todos[i]
// 				break
// 			}
// 		}

// 		if todo == nil {
// 			return c.Status(404).JSON(fiber.Map{
// 				"error": "Todo not found",
// 			})
// 		}

// 		if err := c.BodyParser(todo); err != nil {
// 			return c.Status(400).JSON(fiber.Map{
// 				"error": "Failed to parse body",
// 			})
// 		}

// 		return c.JSON(todo)
// 	})

// 	app.Delete("/todo/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")
// 		for i, todo := range todos {
// 			if strconv.Itoa(todo.ID) == id {
// 				todos = append(todos[:i], todos[i+1:]...)
// 				return c.SendStatus(204)
// 			}
// 		}

// 		return c.Status(404).JSON(fiber.Map{
// 			"error": "Todo not found",
// 		})
// 	})

// 	log.Fatal(app.Listen(":3000"))
// }

// Install GoFiber (Express.JS of Go)
// go get github.com/gofiber/fiber/v2

// Install Air (Nodemon of Go))
// go install github.com/cosmtrek/air@v1.40.4
