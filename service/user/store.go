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
	rows, err := s.db.Queryx("select * from users where email = $1", email)
	if err != nil {
		return nil, err
	}

	u := &types.User{}
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == uuid.Nil {
		return nil, fmt.Errorf("USER NOT FOUND")
	}

	return u, nil
}

func (s *Store) GetUserById(id uuid.UUID) (*types.User, error) {
	rows, err := s.db.Queryx("select * from users where id = $1", id)
	if err != nil {
		return nil, err
	}

	u := &types.User{}
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == uuid.Nil {
		return nil, fmt.Errorf("USER NOT FOUND")
	}

	return u, nil
}

func (s *Store) CreateUser(user *types.User) error {
	return nil
}

func scanRowIntoUser(rows *sqlx.Rows) (*types.User, error) {
	u := &types.User{}
	err := rows.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return u, nil
}
