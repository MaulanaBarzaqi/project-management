package routes

import (
	"log"

	"github.com/MaulanaBarzaqi/project-management/config"
	"github.com/MaulanaBarzaqi/project-management/controllers"
	"github.com/MaulanaBarzaqi/project-management/utils"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
)

func Setup(
	app *fiber.App, 
	uc *controllers.UserController,
	bc *controllers.BoardController,
	) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	} 
	app.Post("/v1/auth/register", uc.Register)
	app.Post("/v1/auth/login", uc.Login)

	// protected routes
	api := app.Group("/api/v1", jwtware.New(jwtware.Config{
		SigningKey: []byte(config.AppConfig.JWTSecret),
		ContextKey: "user",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return utils.Unauthorized(c,"error unauthorized", err.Error())
		},
	}))
	userGroup := api.Group("/users")
	userGroup.Get("/page", uc.GetUserPagination)
	userGroup.Get("/:id", uc.GetUser)
	userGroup.Put("/:id", uc.UpdateUser)
	userGroup.Delete("/:id", uc.DeleteUser)

	boarGroup := api.Group("/boards")
	boarGroup.Post("/", bc.CreateBoard)
	boarGroup.Put("/:id", bc.UpdateBoard)
	boarGroup.Post("/:id/members", bc.AddBoardMembers)
	boarGroup.Delete("/:id/members", bc.RemoveBoardMembers)
}