package repository

import (
	"database/sql"
	"github.com/SyamSolution/user-service/internal/model"
)

type UserRepository struct {
	DB *sql.DB
}

type UserPersister interface {
	CreateUser(user model.UserRequest) (int, error)
	GetUserByEmail(email string) (model.User, error)
}

func NewUserRepository(user UserRepository) UserPersister {
	return &user
}

func (r *UserRepository) CreateUser(user model.UserRequest) (int, error) {
	query := `
		INSERT INTO user (username, email, full_name, phone_number, address, city, country, postal_code, nik, created_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, NOW())`

	result, err := r.DB.Exec(query, user.Username, user.Email, user.FullName, user.PhoneNumber, user.Address, user.City,
		user.Country, user.PostalCode, user.NIK)
	if err != nil {
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastInsertID), nil
}

func (r *UserRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	query := `SELECT user_id, username, email, full_name, phone_number, address, city, country, postal_code, nik, created_at 
		FROM user WHERE email = ?`

	err := r.DB.QueryRow(query, email).Scan(&user.UserID, &user.Username, &user.Email, &user.FullName, &user.PhoneNumber,
		&user.Address, &user.City, &user.Country, &user.PostalCode, &user.NIK, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}
