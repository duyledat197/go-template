package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	kitlog "github.com/go-kit/kit/log"
	"github.com/stamp-server/models"
	"github.com/stamp-server/service/auth"
)

type authHandler struct {
	authService auth.Service
	logger      kitlog.Logger
}

func (h *authHandler) router() chi.Router {
	r := chi.NewRouter()
	r.Post("/register", h.register)
	r.Post("/login", h.login)
	r.Post("/forgot-password", h.forgotPassword)
	r.Post("/reset-password", h.resetPassword)

	return r
}

func (h *authHandler) register(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var request struct {
		UserName        string `json:"userName"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
		Address         string `json:"address"`
		Name            string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.logger.Log("error", err)
		encodeError(ctx, err, w)
		return
	}

	u := models.User{
		UserName: request.UserName,
		Address:  request.Address,
		Name:     request.Name,
	}

	err := h.authService.Register(&u, request.Password)

	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	var response = struct {
		Status string `json:"status"`
	}{
		Status: "success",
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err = json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Log("error", err)
		encodeError(ctx, err, w)
		return
	}
}

func (h *authHandler) login(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var request struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.logger.Log("error", err)
		encodeError(ctx, err, w)
		return
	}
	user, token, err := h.authService.Login(request.UserName, request.Password)
	// if user exist error equal null
	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	var response = struct {
		User   models.User `json:"user"`
		Token  string      `json:"token"`
		Status string      `json:"status"`
	}{
		User:   user,
		Token:  token,
		Status: "success",
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err = json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Log("error", err)
		encodeError(ctx, err, w)
		return
	}
}

func (h *authHandler) forgotPassword(w http.ResponseWriter, r *http.Request) {

}

func (h *authHandler) resetPassword(w http.ResponseWriter, r *http.Request) {

}
