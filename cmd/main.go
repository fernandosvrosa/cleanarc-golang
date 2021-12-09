package main

import (
	"database/sql"
	"github.com/fernandosvrosa/cleanarc-golang/config"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	factory := config.NewFactory(db)
	studentHandler := factory.CreateStudentHandler()

	app := fiber.New()
	app.Post("/v1/students", studentHandler.SaveStudent)

	log.Fatal(app.Listen(":3000"))
}
