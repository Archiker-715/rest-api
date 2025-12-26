package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Archiker-715/rest-api/internal/entity"
	"github.com/Archiker-715/rest-api/internal/repository/pg/items"
	"github.com/gorilla/mux"
)

type ItemHandler struct {
	repo *items.ItemRepository
}

func NewItemHandler(repo *items.ItemRepository) *ItemHandler {
	return &ItemHandler{repo: repo}
}

func (h *ItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.repo.GetItems()
	if err != nil {
		http.Error(w, "failed to fetch items", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var req entity.ItemRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	item := entity.Item{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}

	if err := h.repo.Create(&item); err != nil {
		http.Error(w, "failed to create item", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// maybe check id > 0

	var req entity.ItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	item := entity.Item{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}

	if err := h.repo.Update(&item); err != nil {
		http.Error(w, "failed update item", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// maybe check id > 0

	numId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed get id: %w", err), http.StatusInternalServerError)
	}

	if err := h.repo.Delete(uint(numId)); err != nil {
		http.Error(w, "failed delete item", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
