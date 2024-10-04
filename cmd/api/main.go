package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"tickets/internal/application/usecases"
	"tickets/internal/domain/entities"
	"tickets/internal/handlers"
	"tickets/internal/repositories"
	"tickets/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	engine := gin.Default()

	httpHandlers := buildHttpHandlers()
	server.DefineRoutes(engine, httpHandlers)

	err = engine.Run()
	if err != nil {
		return
	}
}

func buildHttpHandlers() *handlers.HttpHandlers {
	db := connectDatabase()

	userRepository := repositories.NewUserRepository(db)
	useUseCase := usecases.NewUserUseCase(userRepository)

	UserHandler := handlers.NewUserHandler(useUseCase)
	return handlers.NewHttpHandlers(*UserHandler)
}

func connectDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	//dsn := "host=postgres user=user password=password dbname=quest_db port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}

	err = database.AutoMigrate(&entities.User{})
	if err != nil {
		panic("Failed to migrate database")
	}

	return database
}
