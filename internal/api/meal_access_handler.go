package api

import (
	"encoding/json"
	"fmt"
	"github.com/abelmalu/CafeteriaAccessControl/internal/core"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type MealAccessHandler struct {
	service core.MealAccessService
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
	student, err := mh.service.AttemptAccess(studentRfId, cafeteriaId)

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

	json.NewEncoder(w).Encode(student)
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
