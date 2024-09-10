package account_entity

import (
	"time"
)

type Account struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateAccountRequestDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateAccountResponseDTO struct {
	Message string      `json:"message"`
	Data    DataAccount `json:"data"`
}

type LoginAccountRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginAccountResponseDTO struct {
	Message  string      `json:"message"`
	Data     DataAccount `json:"data"`
	JWTToken string      `json:"jwt_token"`
}

type DataAccount struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
