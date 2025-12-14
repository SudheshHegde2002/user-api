package main

import (
	"context"
	"log"
	"os"
	"user-api/db/sqlc"
	"user-api/internal/handler"
	"user-api/internal/repository"
	"user-api/internal/routes"
	"user-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file", envErr)
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("Database url is not set")
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatal("Error parsing db config", err)
	}
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Some error when connecting to db", err)
	}

	queries := sqlc.New(db)

	app := fiber.New()

	repo := repository.NewUserRepository(queries)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)
	routes.RegisterUserRoutes(app, handler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("go fiber running")
	})

	log.Fatal(app.Listen(":3000"))
}
