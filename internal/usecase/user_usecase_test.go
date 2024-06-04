package usecase

import (
	"github.com/SyamSolution/user-service/internal/model"
	"github.com/SyamSolution/user-service/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserUsecase_GetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	email := "test@test.com"

	mockUserRepo := mock.NewMockUserPersister(ctrl)
	mockUserRepo.EXPECT().GetUserByEmail(email).Return(model.User{
		UserID:      1,
		Username:    "testuser",
		Email:       "test@test.com",
		FullName:    "Test User",
		PhoneNumber: "1234567890",
		Address:     "Test Address",
		City:        "Test City",
		Country:     "Test Country",
		PostalCode:  "12345",
		NIK:         "1234567890123456",
	}, nil)

	usecase := &UserUsecase{
		UserRepo: mockUserRepo,
	}

	user, err := usecase.GetUserByEmail(email)

	assert.NoError(t, err)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "test@test.com", user.Email)
}
