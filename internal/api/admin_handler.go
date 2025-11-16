package api

import (
	"encoding/json"
	"net/http"
	"github.com/abelmalu/CafeteriaAccessControl/internal/core"
	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
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
		errString := err.Error()
		http.Error(w, errString, http.StatusInternalServerError)
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
		errorString := err.Error()

		http.Error(w, errorString, http.StatusInternalServerError)
		return
	}
	message := "Cafeteria successfully created"

	json.NewEncoder(w).Encode(created)
	json.NewEncoder(w).Encode(message)

}

func (h *AdminHandler) CreateBatch(w http.ResponseWriter, r *http.Request) {
	var batch models.Batch
	// decode the request body
	err := json.NewDecoder(r.Body).Decode(&batch)
	if err != nil {

		http.Error(w, "Bad Request", http.StatusBadRequest)

		return
	}
	created, serviceErr := h.service.CreateBatch(r.Context(), &batch)
	if serviceErr != nil {

		errorString := serviceErr.Error()

		http.Error(w, errorString, http.StatusBadRequest)

		return

	}

	w.Write([]byte("successfully created a batch"))
	json.NewEncoder(w).Encode(created)

}

func (h *AdminHandler) CreateMeal(w http.ResponseWriter, r *http.Request) {
	var meal models.Meal

	decodingErr := json.NewDecoder(r.Body).Decode(&meal)
	if decodingErr != nil {
		errorString := decodingErr.Error()

		http.Error(w, errorString, http.StatusBadRequest)
		return

	}
	_, err := h.service.CreateMeal(r.Context(), &meal)

	if err != nil {

		errorString := err.Error()

		http.Error(w, errorString, http.StatusBadRequest)
		return

	}
	json.NewEncoder(w).Encode([]byte("Successfully created a meal"))


}

func (h *AdminHandler) RegisterDevice(w http.ResponseWriter, r *http.Request){

	var device models.Device

	decodingErr := json.NewDecoder(r.Body).Decode(&device)


	if decodingErr != nil {
		errorString := decodingErr.Error()
		http.Error(w, errorString, http.StatusBadRequest)
		return
		

	}
	_,err := h.service.RegisterDevice(r.Context(),&device)

	if err != nil{

		errString := err.Error()

		http.Error(w,errString,http.StatusBadRequest)
		return
		
	}
	w.Write([]byte("successfully Registered a device"))



}
