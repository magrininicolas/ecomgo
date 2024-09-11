package types

import (
	"time"

	"github.com/google/uuid"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id uuid.UUID) (*User, error)
	CreateUser(*User) error
}

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	FirstName string    `json:"firstName" db:"first_name"`
	LastName  string    `json:"lastName" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=50"`
}

type RegisterUserResponse struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=50"`
}

func NewUser(firstName, lastName, email, password string) *User {
	return &User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
