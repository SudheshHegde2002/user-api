package main

import (
	"os"
	"context"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
	"user-api/db/sqlc"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
)

func main(){
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file", envErr)
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("Database url is not set")
	}

	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Some error when connecting to db", err)
	}

	queries := sqlc.New(db)

	app := fiber.New()

	app.Get("/", func(c*fiber.Ctx)error{
		return c.SendString("go fiber running")
	})

	log.Fatal(app.Listen(":3000"))
}