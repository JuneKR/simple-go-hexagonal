package main

import (
	"neptune/core"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	// Initialize the database connection
	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&core.Order{})

	// Start the server
	app.Listen(":8000")
}
