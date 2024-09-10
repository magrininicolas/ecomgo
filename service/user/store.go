package user

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/magrininicolas/ecomgo/types"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	u := &types.User{}

	err := s.db.Get(u, "select * from users where email=$1", email)
	if err != nil {
		return nil, err
	}

	if u.ID == uuid.Nil {
		return nil, fmt.Errorf("USER NOT FOUND")
	}

	return u, nil
}

func (s *Store) GetUserById(id uuid.UUID) (*types.User, error) {
	u := &types.User{}

	err := s.db.Get(u, "select * from users where id=$1", id)
	if err != nil {
		return nil, err
	}

	if u.ID == uuid.Nil {
		return nil, fmt.Errorf("USER NOT FOUND")
	}

	return u, nil
}

func (s *Store) CreateUser(user *types.User) error {
	query := `insert into users (id, first_name, last_name, email, password, created_at, updated_at)
		values($1, $2, $3, $4, $5, $6, $7)`

	_, err := s.db.Exec(query,
		user.ID, user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
