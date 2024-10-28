package routes

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

// Initialize routes with the collection
func SetupRoutes(app *fiber.App, coll *mongo.Collection) {
	collection = coll // Assign the passed collection to the global variable

	app.Get("/api/todo", GetToDo)
	app.Post("/api/todo", PostToDo)
}

// Define handlers for each route
// Get ToDo's Requests
func GetToDo(c *fiber.Ctx) error {
	var todos []Todo

	// bson.M{} is an empty filter that fetches all documents in the collection
	cursor, err := collection.Find(context.Background(), bson.M{})
	if HandleError(c, err) {
		return nil // Error handled; exit function
	}
	defer cursor.Close(context.Background()) // Ensures cursor closes after function execution

	// Iterate through each document in the cursor
	for cursor.Next(context.Background()) {
		var todo Todo
		err := cursor.Decode(&todo) // Decode each document into a Todo struct
		if HandleError(c, err) {
			return nil // Error handled; exit function
		}
		todos = append(todos, todo) // Add decoded Todo to the todos slice
	}

	// Return the list of todos in JSON format with status 200
	return c.Status(200).JSON(todos)
}

// Post ToDo's Requests
func PostToDo(c *fiber.Ctx) error {
	todo := Todo{}

	err := c.BodyParser(&todo)
	if HandleError(c, err) {
		return nil
	}

	if todo.Body == "" {
		// could be handeled in front end
		return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)
	if HandleError(c, err) {
		return nil
	}

	// Set the ID of the inserted Todo
	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	// Return the inserted Todo with status 201
	return c.Status(201).JSON(todo)
}

// func PatchToDo(c *fiber.Ctx) error {
// 	// Placeholder response for updating a todo
// 	return c.JSON(fiber.Map{"message": "Patch Todo"})
// }

// func DeleteToDo(c *fiber.Ctx) error {
// 	// Placeholder response for deleting a todo
// 	return c.JSON(fiber.Map{"message": "Delete Todo"})
// }

func HandleError(c *fiber.Ctx, err error) bool {
	if err != nil {
		log.Println("There was an error in the code:", err)
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
		return true
	}
	return false
}
