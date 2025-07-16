package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nocson47/go-hex-concept/internal/adapter/handler/http"
	"github.com/nocson47/go-hex-concept/internal/adapter/repository/postgres"
	service "github.com/nocson47/go-hex-concept/internal/core/service"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Get connection string from environment variable
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL not set in environment")
	}

	// Initialize database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Verify database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to database successfully")
	userRepository := postgres.NewUsersRepositoryDB(db)
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserHandler(userService)

	// Initialize Fiber app
	app := fiber.New()

	// Define API routes
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// User routes
	users := v1.Group("/users")
	users.Get("/", userHandler.GetAllUsers)
	users.Get("/:id", userHandler.GetUser)
	users.Post("/", userHandler.CreateUser)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
