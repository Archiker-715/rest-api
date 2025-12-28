package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Archiker-715/rest-api/internal/entity"
	"github.com/Archiker-715/rest-api/internal/repository/pg/items"
	"github.com/Archiker-715/rest-api/internal/usecase"
	"github.com/gorilla/mux"
)

type ItemHandler struct {
	items *usecase.ItemService
}

func NewItemHandler(repo *items.ItemRepository) *ItemHandler {
	return &ItemHandler{items: usecase.NewItemService(repo)}
}

func (h *ItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.items.GetItems()
	if err != nil {
		http.Error(w, "failed to fetch items", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// @Summary Создание элемента
// @Description Создает новый элемент в системе
// @Tags items
// @Accept json
// @Produce json
// @Param item body entity.ItemRequest true "Данные элемента"
// @Success 200 {object} entity.Item
// @Failure 400 {object} ErrorResponse
// @Router /api/v1/items [post]
func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var req entity.ItemRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	item, err := h.items.CreateItem(&req)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to create item: "+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	// maybe check id > 0

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed convert id: %v", err), http.StatusBadRequest)
	}

	var req *entity.ItemRequest
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	item, err := h.items.UpdateItem(req, id)
	if err != nil {
		http.Error(w, "failed update item", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	// maybe check id > 0

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed get id: %v", err), http.StatusInternalServerError)
	}

	if err := h.items.DeleteItem(id); err != nil {
		http.Error(w, "failed delete item", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
