package migrations

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDBmain() error {
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

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbPassword)
	DB, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", dbName)
	DB.Exec(createDatabaseCommand)

	dbURL := "postgres://" + dbUsername + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=" + dbSSLMode
	cmd := exec.Command("migrate", "-path", "db/migrations", "-database", dbURL, "up")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}
	return nil
}
