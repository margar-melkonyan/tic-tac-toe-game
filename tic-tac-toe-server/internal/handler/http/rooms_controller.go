package http_handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/helper"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/service"
)

type RoomHandler struct {
	service service.RoomService
}

func NewRoomHandler(service service.RoomService) *RoomHandler {
	return &RoomHandler{
		service: service,
	}
}

func (h *RoomHandler) GetRooms(w http.ResponseWriter, r *http.Request) {
	resp := helper.Response{}
	data := h.service.GetAll(r.Context())
	resp.Data = data
	resp.ResponseWrite(w, r, http.StatusOK)
}

func (h *RoomHandler) GetRoom(w http.ResponseWriter, r *http.Request) helper.Response {
	resp := helper.Response{}
	param := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		resp.Errors = err
		return resp
	}
	data, err := h.service.GetById(r.Context(), id)
	if err != nil {
		resp.Errors = err.Error()
		return resp
	}
	resp.Data = data
	return resp
}

func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	resp := helper.Response{}
	if resp.IsValidMediaType(w, r) {
		return
	}
	var form common.RoomRequest
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
		resp.Errors = humanReadableErrors
		resp.ResponseWrite(w, r, http.StatusUnprocessableEntity)
		return
	}
	if err := h.service.Create(r.Context(), form); err != nil {
		resp.ResponseWrite(w, r, http.StatusInternalServerError)
		return
	}
	resp.Message = "Created!"
	resp.ResponseWrite(w, r, http.StatusOK)
}

func (h *RoomHandler) DestroyRoom(w http.ResponseWriter, r *http.Request) {
	resp := helper.Response{}
	param := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		resp.ResponseWrite(w, r, http.StatusNotFound)
		return
	}
	if err := h.service.DeleteById(r.Context(), id); err != nil {
		resp.ResponseWrite(w, r, http.StatusNoContent)
		return
	}
	resp.ResponseWrite(w, r, http.StatusOK)
}
