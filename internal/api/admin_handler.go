package api

import (
	"encoding/json"
	"github.com/abelmalu/CafeteriaAccessControl/internal/core"
	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
	"net/http"
)

type AdminHandler struct {
	service core.AdminService
}

func NewAdminHandler(service core.AdminService) *AdminHandler {
	return &AdminHandler{service: service}
}

// CreateStudent handles the api/admin/create/student route
func (h *AdminHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {

		http.Error(w, "invalid input", http.StatusBadRequest)

		return
	}

	created, err := h.service.CreateStudent(r.Context(), &student)
	if err != nil {
		http.Error(w, "failed to create student", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(created)

}

// CreateCafeteria handles the api/admin/create/Cafeteria route
func (h *AdminHandler) CreateCafeteria(w http.ResponseWriter, r *http.Request) {
	var cafeteria models.Cafeteria
	err := json.NewDecoder(r.Body).Decode(&cafeteria)

	if err != nil {

		http.Error(w, "Invalid data sent", http.StatusBadRequest)
		return
	}

	created, err := h.service.CreateCafeteria(r.Context(), &cafeteria)

	if err != nil {
		http.Error(w, "failed to create cafeteria", http.StatusInternalServerError)
		return
	}
	message := "Cafeteria successfully created"

	json.NewEncoder(w).Encode(created)
	json.NewEncoder(w).Encode(message)

}
