package handler

import (
	"github.com/SyamSolution/user-service/internal/model"
	"github.com/SyamSolution/user-service/mock"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserHandler_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := model.UserRequest{
		Username:    "testuser",
		Email:       "testuser@test.com",
		FullName:    "Test User",
		PhoneNumber: "1234567890",
		Address:     "Test Address",
		City:        "Test City",
		Country:     "Test Country",
		PostalCode:  "12345",
		NIK:         "1234567890123456",
	}

	mockUserUsecase := mock.NewMockUserExecutor(ctrl)
	mockUserUsecase.EXPECT().CreateUser(user).Return(1, nil)

	handler := &User{
		UserUsecase: mockUserUsecase,
	}

	app := fiber.New()
	app.Post("/users", handler.CreateUser)

	req := httptest.NewRequest("POST", "/users", strings.NewReader(`{
		"username": "testuser",
		"email": "testuser@test.com",
		"full_name": "Test User",
		"phone_number": "1234567890",
		"address": "Test Address",
		"city": "Test City",
		"country": "Test Country",
		"postal_code": "12345",
		"nik": "1234567890123456"
	}`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, 2*1024)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}

func TestUserHandler_ConfirmUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userCode := model.ConfirmCode{
		Email: "testuser@test.com",
		Code:  "123456",
	}

	mockUserUsecase := mock.NewMockUserExecutor(ctrl)
	mockUserUsecase.EXPECT().ConfirmUser(userCode).Return(nil, "success")

	handler := &User{
		UserUsecase: mockUserUsecase,
	}

	app := fiber.New()
	app.Post("/users/confirm", handler.ConfirmUser)

	req := httptest.NewRequest("POST", "/users/confirm", strings.NewReader(`{
		"email": "testuser@test.com",
		"code": "123456"
	}`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, 2*1024)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}

func TestUserHandler_SignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := model.SignIn{
		Email:    "testuser@test.com",
		Password: "testpassword",
	}

	mockUserUsecase := mock.NewMockUserExecutor(ctrl)
	mockUserUsecase.EXPECT().LoginUser(user).Return(nil, "success", &cognitoidentityprovider.InitiateAuthOutput{
		AuthenticationResult: &cognitoidentityprovider.AuthenticationResultType{
			IdToken:      new(string),
			RefreshToken: new(string),
		},
	})

	handler := &User{
		UserUsecase: mockUserUsecase,
	}

	app := fiber.New()
	app.Post("/users/signin", handler.SignIn)

	req := httptest.NewRequest("POST", "/users/signin", strings.NewReader(`{
  "email": "testuser@test.com",
  "password": "testpassword"
 }`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, 2*1024)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}

func TestUserHandler_RefreshToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	refreshToken := "testRefreshToken"

	mockUserUsecase := mock.NewMockUserExecutor(ctrl)
	mockUserUsecase.EXPECT().RefreshToken(refreshToken).Return(nil, "success", &cognitoidentityprovider.InitiateAuthOutput{
		AuthenticationResult: &cognitoidentityprovider.AuthenticationResultType{
			IdToken:      new(string),
			RefreshToken: new(string),
		},
	})

	handler := &User{
		UserUsecase: mockUserUsecase,
	}

	app := fiber.New()
	app.Post("/users/refresh", handler.RefreshToken)

	req := httptest.NewRequest("POST", "/users/refresh", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", refreshToken)

	resp, err := app.Test(req, 2*1024)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}

func TestUserHandler_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	email := "test@test.com"

	mockUserUsecase := mock.NewMockUserExecutor(ctrl)
	mockUserUsecase.EXPECT().GetUserByEmail(email).Return(model.UserResponse{}, nil)

	handler := &User{
		UserUsecase: mockUserUsecase,
	}

	app := fiber.New()
	app.Get("/users/:email", func(ctx *fiber.Ctx) error {
		ctx.Locals("email", email)
		return handler.GetUser(ctx)
	})

	req := httptest.NewRequest("GET", "/users/"+email, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, 2*1024)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}
