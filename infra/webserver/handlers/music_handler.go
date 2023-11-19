package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/vitorconti/loader-pagination/infra/database"
)

type MusicHandler struct {
	MusicDB database.MusicInterface
}

func NewProductHandler(db database.MusicInterface) *MusicHandler {
	return &MusicHandler{
		MusicDB: db,
	}
}

func (h *MusicHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}
	products, err := h.MusicDB.FindAll(pageInt, limitInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
