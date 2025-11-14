package api

import (
	"encoding/json"
	//"fmt"

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
