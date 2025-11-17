package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/abelmalu/CafeteriaAccessControl/internal/core"
	"github.com/go-chi/chi/v5"
)

type MealAccessHandler struct {
	service core.MealAccessService
}


func NewMealAccessHandler(svc core.MealAccessService) *MealAccessHandler{


	return  &MealAccessHandler{service: svc}
}



func (mh *MealAccessHandler) AttemptAccess(w http.ResponseWriter, r *http.Request){
	
	
	studentRfId := chi.URLParam(r, "sutdentRfid")
	cafeteriaId := chi.URLParam(r,"cafeteriaId")

	fmt.Printf("Received request for RFID Tag %s\n",studentRfId)
	if studentRfId == ""{
		http.Error(w,"invalid rfid tag",http.StatusBadRequest)
		return
	}
	student,err := mh.service.AttemptAccess(studentRfId,cafeteriaId)

	if err != nil{

		errStr := err.Error()

		http.Error(w,errStr,http.StatusBadRequest)
		return
		
	}
	
	json.NewEncoder(w).Encode(student)
	w.Write([]byte("student fetched successfully"))

	



}