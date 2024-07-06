package main

import (
	"log"
	"os"
	"os/exec"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"tt/handler"
	"tt/repository"
	"tt/service"
	"tt/testtask"
)

func main() {
	err := godotenv.Load("connect.env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	db, err := repository.DBC(repository.Conf{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUsername,
		Password: dbPassword,
		BDname:   dbName,
		SSLMode:  dbSSLMode,
	})
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	dbURL := "postgresql://" + dbUsername + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=" + dbSSLMode
	cmd := exec.Command("migrate", "-path", "db/migrations", "-database", dbURL, "up")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(testtask.Server)
	if err := srv.Run("localhost", "8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Launch error HTTP server: %v", err)
	}
}
