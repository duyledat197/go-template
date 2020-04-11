package http

import (
	"net/http"

	"github.com/go-chi/chi"
	kitlog "github.com/go-kit/kit/log"
	"github.com/duyledat197/go-template/service/user"
)

type userHandler struct {
	userService user.Service
	logger      kitlog.Logger
}

func (h *userHandler) router() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.getListUser)
	r.Get("/{userID}", h.getUserByID)
	r.Put("/{userID}", h.UpdateByID)
	r.Delete("/{userID}", h.DeleteByID)

	return r
}

func (h *userHandler) getListUser(w http.ResponseWriter, r *http.Request) {
}

func (h *userHandler) getUserByID(w http.ResponseWriter, r *http.Request) {
}

func (h *userHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {

}

func (h *userHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {

}
