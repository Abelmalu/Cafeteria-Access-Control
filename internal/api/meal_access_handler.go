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



func (mh *MealAccessHandler) GetStudentByRfidTag(w http.ResponseWriter, r *http.Request){

	studentRfid :=chi.URLParam(r, "sutdentRfid")

	fmt.Printf("Received request for RFID Tag %s\n",studentRfid)
	if studentRfid == ""{
		http.Error(w,"invalid rfid tag",http.StatusBadRequest)
		return
	}
	student,err := mh.service.GetStudentByRfidTag(studentRfid)

	if err != nil{

		errStr := err.Error()

		http.Error(w,errStr,http.StatusBadRequest)
		return
		
	}
	
	json.NewEncoder(w).Encode(student)
	w.Write([]byte("student fetched successfully"))

	



}