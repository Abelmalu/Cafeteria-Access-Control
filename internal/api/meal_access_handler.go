package api

import (
	"net/http"

	"github.com/abelmalu/CafeteriaAccessControl/internal/core"
)

type MealAccessHandler struct {
	service core.MealAccessService
}


func NewMealAccessHandler(svc core.MealAccessService) *MealAccessHandler{


	return  &MealAccessHandler{service: svc}
}



func (mh *MealAccessHandler) GetStudentByRfidTag(w http.ResponseWriter, r *http.Request){



}