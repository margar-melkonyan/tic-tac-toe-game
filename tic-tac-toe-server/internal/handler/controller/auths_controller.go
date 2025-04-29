package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/helper"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/service"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) SingIn(w http.ResponseWriter, r *http.Request) {
	resp := helper.Response{}
	if resp.IsValidMediaType(w, r) {
		return
	}
	var form common.AuthSignInRequest
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		slog.Error("Error decoding JSON: " + err.Error())
		resp.ResponseWrite(w, r, http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(&form)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		humanReadableErrors, err := helper.LocalizedValidationMessages(
			r.Context(),
			errs,
		)
		if err != nil {
			slog.Error("Error localizing validation messages: " + err.Error())
			resp.ResponseWrite(w, r, http.StatusInternalServerError)
			return
		}
		resp.Data = humanReadableErrors
		resp.ResponseWrite(w, r, http.StatusUnprocessableEntity)
		return
	}
	token, err := h.service.SignIn(r.Context(), form)
	if err != nil {
		resp.ResponseWrite(w, r, http.StatusInternalServerError)
		return
	}
	resp.Data = token
	resp.ResponseWrite(w, r, http.StatusOK)
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	resp := helper.Response{}
	if resp.IsValidMediaType(w, r) {
		return
	}
	var form common.AuthSignUpRequest
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		slog.Error("Error decoding JSON: " + err.Error())
		resp.ResponseWrite(w, r, http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(&form)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		humanReadableErrors, err := helper.LocalizedValidationMessages(
			r.Context(),
			errs,
		)
		if err != nil {
			slog.Error("Error localizing validation messages: " + err.Error())
			resp.ResponseWrite(w, r, http.StatusInternalServerError)
			return
		}
		resp.Data = humanReadableErrors
		resp.ResponseWrite(w, r, http.StatusUnprocessableEntity)
		return
	}
	if err := h.service.SignUp(r.Context(), form); err != nil {
		resp.Message = err.Error()
		resp.ResponseWrite(w, r, http.StatusConflict)
		return
	}
	resp.Message = "Created!"
	resp.ResponseWrite(w, r, http.StatusOK)
}
