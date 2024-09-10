package user

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/magrininicolas/ecomgo/service/auth"
	"github.com/magrininicolas/ecomgo/types"
	"github.com/magrininicolas/ecomgo/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Post("/login", h.handleLogin)
	r.Post("/register", h.handleRegister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload

	err := utils.ParseJSON(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("MALFORMATED JSON BODY"))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	if ok, err := h.emailExists(payload.Email); ok {
		utils.WriteJSON(w, http.StatusBadRequest, err)
	}

	hashedPasswd, err := auth.HashPasswd(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("ERROR WHILE HASHING PASSWORD %s", err.Error()))
		return
	}

	newUser := types.NewUser(payload.FirstName, payload.LastName, payload.Email, hashedPasswd)
	err = h.store.CreateUser(newUser)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, types.RegisterUserResponse{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
	})
}

func (h *Handler) emailExists(email string) (bool, error) {
	_, err := h.store.GetUserByEmail(email)
	if err != nil {
		return false, fmt.Errorf("USER WITH EMAIL %s DONT EXISTS", email)
	}
	return true, fmt.Errorf("USER WITH EMAIL %s ALREADY EXISTS", email)
}
