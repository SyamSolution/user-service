package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SyamSolution/user-service/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectExec("INSERT INTO user").
		WithArgs("testuser", "testuser@test.com", "Test User", "1234567890", "Test Address", "Test City", "Test Country", "12345", "1234567890123456").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewUserRepository(UserRepository{DB: db})

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

	_, err = repo.CreateUser(user)

	assert.NoError(t, err)
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"user_id", "username", "email", "full_name", "phone_number", "address", "city", "country", "postal_code", "nik", "created_at"}).
		AddRow(1, "testuser", "testuser@test.com", "Test User", "1234567890", "Test Address", "Test City", "Test Country", "12345", "1234567890123456", time.Now())

	mock.ExpectQuery("SELECT user_id, username, email, full_name, phone_number, address, city, country, postal_code, nik, created_at FROM user WHERE email = ?").
		WithArgs("testuser@test.com").
		WillReturnRows(rows)

	repo := NewUserRepository(UserRepository{DB: db})

	_, err = repo.GetUserByEmail("testuser@test.com")

	assert.NoError(t, err)
}
