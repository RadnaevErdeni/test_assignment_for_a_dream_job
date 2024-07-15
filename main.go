package main

import (
	"os"
	"tt/db/migrations"
	_ "tt/docs"
	"tt/handler"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"tt/repository"
	"tt/service"
	"tt/testtask"
)

func init() {
	testtask.Init()
}

func main() {
	logrus.Info("Starting application")

	err := migrations.StartDBmain()
	if err != nil {
		logrus.Fatalf("Error initializing database 1: %v", err)
	}
	logrus.Info("Successfully initialized database 1")

	err = godotenv.Load("connect.env")
	if err != nil {
		logrus.Fatalf("Failed to load .env file: %v", err)
	}
	logrus.Info("Loaded .env file successfully")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")
	cnPort := os.Getenv("CON_PORT")
	cnHost := os.Getenv("CON_HOST")

	logrus.WithFields(logrus.Fields{
		"dbHost": dbHost,
		"dbPort": dbPort,
		"dbName": dbName,
		"cnHost": cnHost,
		"cnPort": cnPort,
	}).Debug("Loaded environment variables")

	db, err := repository.DBC(repository.Conf{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUsername,
		Password: dbPassword,
		BDname:   dbName,
		SSLMode:  dbSSLMode,
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize the database: %v", err)
	}
	logrus.Info("Successfully initialized the database")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(testtask.Server)
	if err := srv.Run(cnHost, cnPort, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Launch error HTTP server: %v", err)
	}

	logrus.Info("Application started successfully")
}
