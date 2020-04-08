package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	kitlog "github.com/go-kit/kit/log"
	"github.com/stamp-server/models"
	"github.com/stamp-server/service/wallet"
)

type walletHandler struct {
	walletService wallet.Service
	logger        kitlog.Logger
}

func (h *walletHandler) router() chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.createWallet)

	return r
}

func (h *walletHandler) createWallet(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	user, _ := r.Context().Value("user").(*models.User)
	address, privateKey, err := h.walletService.Create(user.ID)
	if err != nil {
		encodeError(ctx, err, w)
	}
	var response = struct {
		Address    string `json:"address"`
		PrivateKey string `json:"privateKey"`
		Success    bool   `json:"success"`
	}{
		Address:    address,
		PrivateKey: privateKey,
		Success:    true,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err = json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Log("error", err)
		encodeError(ctx, err, w)
		return
	}
}
