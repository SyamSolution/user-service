package main

import (
	"fmt"
	"github.com/SyamSolution/user-service/config"
	"github.com/SyamSolution/user-service/internal/handler"
	"github.com/SyamSolution/user-service/internal/repository"
	"github.com/SyamSolution/user-service/internal/usecase"
	"github.com/SyamSolution/user-service/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	baseDep := config.NewBaseDep()
	loadEnv(baseDep.Logger)
	db, err := config.NewDbPool(baseDep.Logger)
	if err != nil {
		os.Exit(1)
	}

	//=== repository lists start ===//
	userRepo := repository.UserRepository(repository.UserRepository{
		DB: db,
	})
	//=== repository lists end ===//

	//=== usecase lists start ===//
	userUsecase := usecase.NewUserUsecase(&usecase.UserUsecase{
		UserRepo: userRepo,
	})
	//=== usecase lists end ===//

	//=== handler lists start ===//
	userHandler := handler.NewUserHandler(handler.User{
		UserUsecase: userUsecase,
	})
	//=== handler lists end ===//

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//=== user routes ===//
	app.Post("/users/register", userHandler.CreateUser)
	app.Post("/users/confirm", userHandler.ConfirmUser)
	app.Post("/users/login", userHandler.SignIn)
	app.Get("/users/refresh-token", userHandler.RefreshToken)
	app.Get("/users/profile", middleware.Auth(), userHandler.GetUser)

	//=== listen port ===//
	if err := app.Listen(fmt.Sprintf(":%s", "3000")); err != nil {
		log.Fatal(err)
	}
}

func loadEnv(logger config.Logger) {
	_, err := os.Stat(".env")
	if err == nil {
		err = godotenv.Load()
		if err != nil {
			logger.Error("no .env files provided")
		}
	}
}
