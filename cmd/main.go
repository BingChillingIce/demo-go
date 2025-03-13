package main

import (
	"log"

	"gihub.com/BingChillingIce/demo-go/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	log.Fatal(app.Listen(":3000"))

	db, err := database.Database()

	if err != nil {
		log.Fatal(&err)
	}
	log.Panic(db)
}
