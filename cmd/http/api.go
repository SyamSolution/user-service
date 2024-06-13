package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SyamSolution/user-service/config"
	"github.com/SyamSolution/user-service/config/middleware"
	"github.com/SyamSolution/user-service/internal/handler"
	"github.com/SyamSolution/user-service/internal/repository"
	"github.com/SyamSolution/user-service/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	baseDep := config.NewBaseDep()
	loadEnv(baseDep.Logger)
	db, err := config.NewDbPool(baseDep.Logger)
	if err != nil {
		os.Exit(1)
	}

	dbCollector := middleware.NewStatsCollector("assesment", db)
	prometheus.MustRegister(dbCollector)
	fiberProm := middleware.NewWithRegistry(prometheus.DefaultRegisterer, "user-service", "", "", map[string]string{})

	//=== repository lists start ===//
	userRepo := repository.NewUserRepository(repository.UserRepository{
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
	
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(pprof.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeInterval: time.Millisecond,
		TimeFormat:   "02-01-2006 15:04:05",
		TimeZone:     "Indonesia/Jakarta",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("UP")
	})

	//=== metrics route
	fiberProm.RegisterAt(app, "/metrics")
	app.Use(fiberProm.Middleware)

	api := app.Group("/api/v1")
	//=== user routes ===//
	api.Post("/users/register", userHandler.CreateUser)
	api.Post("/users/confirm", userHandler.ConfirmUser)
	api.Post("/users/login", userHandler.SignIn)
	api.Get("/users/refresh-token", userHandler.RefreshToken)
	api.Get("/users/profile", middleware.Auth(), userHandler.GetUser)

	//=== listen port ===//
	if err := app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))); err != nil {
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
