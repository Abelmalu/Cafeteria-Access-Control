package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/abelmalu/CafeteriaAccessControl/internal/core"
	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
	"github.com/go-chi/chi/v5"
)

type MealAccessHandler struct {
	service core.MealAccessService
}

type APIResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	// omitempty ensures the field is not present in the JSON if it's nil
	Data *models.Student `json:"data,omitempty"`
}

func NewMealAccessHandler(svc core.MealAccessService) *MealAccessHandler {

	return &MealAccessHandler{service: svc}
}

func (mh *MealAccessHandler) AttemptAccess(w http.ResponseWriter, r *http.Request) {

	studentRfId := chi.URLParam(r, "sutdentRfid")
	cafeteriaId := chi.URLParam(r, "cafeteriaId")

	fmt.Printf("Received request for RFID Tag %s\n", studentRfId)
	if studentRfId == "" {
		http.Error(w, "invalid rfid tag", http.StatusBadRequest)
		return
	}
	student, accessStatus, err := mh.service.AttemptAccess(studentRfId, cafeteriaId)

	if err != nil {

		errStr := err.Error()
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		errMap := map[string]string{

			"status":  "error",
			"message": errStr,
		}

		errJson, _ := json.Marshal(errMap)
		w.Write(errJson)

		return

	}
	switch accessStatus {

	case "Granted":
		response := APIResponse{
			Status:  "success",
			Message: "Granted",
			Data:    student,
		}
		json.NewEncoder(w).Encode(response)

	case "Denied":
		response := APIResponse{
			Status:  "success",
			Message: "Denied",
			Data:    student,
		}
		json.NewEncoder(w).Encode(response)
	case "Not Meal Time":
		response := APIResponse{
			Status:  "success",
			Message: "Not Meal Time",
			Data:    student,
		}
		json.NewEncoder(w).Encode(response)
	case "Wrong Cafeteria":
		response := APIResponse{
			Status:  "success",
			Message: "Wrong Cafeteria",
			Data:    student,
		}
		json.NewEncoder(w).Encode(response)

	}

	// json.NewEncoder(w).Encode(student)
	// w.Write([]byte("student fetched successfully"))

}

func (mh *MealAccessHandler) GetCafeterias(w http.ResponseWriter, r *http.Request) {

	cafeterias, err := mh.service.GetCafeterias()

	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"status":"error","message":"something went wrong"}`))

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	cafteriasJson, _ := json.Marshal(cafeterias)
	w.Write(cafteriasJson)

}

func (mh *MealAccessHandler) VerifyDevice(w http.ResponseWriter, r *http.Request) {

	SerialNumber := chi.URLParam(r, "SerialNumber")
	fmt.Println(SerialNumber)

	exists := mh.service.VerifyDevice(SerialNumber)

	if exists {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"status":"success","message":"the device is a valid registered device"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"status":"error","message":"the device is not a valid registered device"}`))

}
