package handler

import (
	"github.com/SyamSolution/user-service/internal/model"
	"github.com/SyamSolution/user-service/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	UserUsecase usecase.UserExecutor
}

type UserHandler interface {
	CreateUser(c *fiber.Ctx) error
	ConfirmUser(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
}

func NewUserHandler(handler User) UserHandler {
	return &handler
}

func (handler *User) CreateUser(c *fiber.Ctx) error {
	var user model.UserRequest
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	userID, err := handler.UserUsecase.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user_id": userID,
	})
}

func (handler *User) ConfirmUser(c *fiber.Ctx) error {
	var userCode model.ConfirmCode
	if err := c.BodyParser(&userCode); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	err, _ := handler.UserUsecase.ConfirmUser(userCode)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User confirmed successfully",
	})
}

func (handler *User) SignIn(c *fiber.Ctx) error {
	var user model.SignIn
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	err, _, initiateAuthOutput := handler.UserUsecase.LoginUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":       "login successfully",
		"token":         *initiateAuthOutput.AuthenticationResult.IdToken,
		"refresh_token": *initiateAuthOutput.AuthenticationResult.RefreshToken,
	})
}

func (handler *User) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Get("Authorization")
	if refreshToken == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	err, _, initiateAuthOutput := handler.UserUsecase.RefreshToken(refreshToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "refresh token successfully",
		"token":   *initiateAuthOutput.AuthenticationResult.IdToken,
		//"refresh_token": *initiateAuthOutput.AuthenticationResult.RefreshToken,
	})
}

func (handler *User) GetUser(c *fiber.Ctx) error {
	email := c.Locals("email")
	user, err := handler.UserUsecase.GetUserByEmail(email.(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"user":    user,
	})
}
