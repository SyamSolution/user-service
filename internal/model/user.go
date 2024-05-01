package model

import "time"

type User struct {
	UserID      int       `json:"user_id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	PostalCode  string    `json:"postal_code"`
	NIK         string    `json:"nik"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Country     string `json:"country"`
	PostalCode  string `json:"postal_code"`
	NIK         string `json:"nik"`
}

type UserResponse struct {
	UserID      int    `json:"user_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Country     string `json:"country"`
	PostalCode  string `json:"postal_code"`
	NIK         string `json:"nik"`
}

type ConfirmCode struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
