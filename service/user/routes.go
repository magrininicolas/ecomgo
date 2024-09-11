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
	var payload types.LoginUserPayload

	err := utils.ParseJSON(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("MALFORMATED JSON BODY"))
		return
	}

	err = utils.Validate.Struct(payload)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("INVALID PAYLOAD %v", errors))
		return
	}

	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("USER NOT FOUND, INVALID EMAIL OR PASSWORD"))
		return
	}

	if !auth.ComparePassword(u.Password, payload.Password) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("USER NOT FOUND, INVALID EMAIL OR PASSWORD"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": "jwt auth under construction"})
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
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("INVALID PAYLOAD %v", errors))
		return
	}

	_, err = h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("USER ALREADY EXISTS FOR EMAIL: %s", payload.Email))
		return
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
