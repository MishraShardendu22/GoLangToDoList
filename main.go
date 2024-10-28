package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hi This is my first GoLang Project")

	app := fiber.New()
	log.Fatal(app.Listen(":3000"))
}

// Install GoFiber
// go get github.com/gofiber/fiber/v2
