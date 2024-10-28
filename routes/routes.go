package routes

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID        int    `json:"_id" bson:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var collection *mongo.Collection

// Define handlers for each route
func GetToDo(c *fiber.Ctx) error {
	var todos []Todo

	// bson.M{} is an empty filter that fetches all documents in the collection
	cursor, err := collection.Find(context.Background(), bson.M{})
	HandleError(err)
	defer cursor.Close(context.Background()) // Ensures cursor closes after function execution

	// Iterate through each document in the cursor
	for cursor.Next(context.Background()) {
		var todo Todo
		err := cursor.Decode(&todo) // Decode each document into a Todo struct
		HandleError(err)
		todos = append(todos, todo) // Add decoded Todo to the todos slice
	}

	// Return the list of todos in JSON format with status 200
	return c.Status(200).JSON(todos)
}


// func PostToDo(c *fiber.Ctx) error {
// 	// Placeholder response for posting a todo
// 	return c.JSON(fiber.Map{"message": "Post Todo"})
// }

// func PatchToDo(c *fiber.Ctx) error {
// 	// Placeholder response for updating a todo
// 	return c.JSON(fiber.Map{"message": "Patch Todo"})
// }

// func DeleteToDo(c *fiber.Ctx) error {
// 	// Placeholder response for deleting a todo
// 	return c.JSON(fiber.Map{"message": "Delete Todo"})
// }

func HandleError(err error) {
	if err != nil {
		log.Fatal("There was an error in the code ", err)
	}
}
