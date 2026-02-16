package main

import (
	"log"

	"github.com/MaulanaBarzaqi/project-management/config"
	"github.com/MaulanaBarzaqi/project-management/controllers"
	"github.com/MaulanaBarzaqi/project-management/database/seed"
	"github.com/MaulanaBarzaqi/project-management/repositories"
	"github.com/MaulanaBarzaqi/project-management/routes"
	"github.com/MaulanaBarzaqi/project-management/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()
	app := fiber.New()
	// user setup
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	// board setup
	boardRepo := repositories.NewBoardRepository()
	boardService := services.NewBoardService(boardRepo, userRepo)
	boardController := controllers.NewBoardController(boardService)

	routes.Setup(app,userController, boardController)

	port := config.AppConfig.AppPort
	log.Println("server is running on port :", port)
	log.Fatal(app.Listen(":" + port))
}