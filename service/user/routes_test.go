package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/magrininicolas/ecomgo/types"
)

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUserStore{}
	h := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := &types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "123",
			Email:     "",
			Password:  "asd",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := chi.NewRouter()

		router.HandleFunc("/register", h.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should successfully create an user", func(t *testing.T) {
		payload := &types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "123",
			Email:     "user@email.com",
			Password:  "asd",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := chi.NewRouter()

		router.HandleFunc("/register", h.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserById(id uuid.UUID) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) CreateUser(*types.User) error {
	return nil
}
