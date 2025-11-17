package service

import (
	"errors"
	"strconv"

	"github.com/abelmalu/CafeteriaAccessControl/internal/core"
	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
)

type MealAccessService struct {
	repo core.MealAccessServiceRepository
}

func NewMealAccessService(repo core.MealAccessServiceRepository) *MealAccessService {

	return &MealAccessService{repo: repo}

}

func (ms *MealAccessService) GetStudentByRfidTag(rfidTag string,cafeteriaId string) (*models.Student, error) {

	if rfidTag == "" {

		return nil,errors.New("RFIDTag value empty")

	}
	if cafeteriaId == "" {

		return nil,errors.New("cafeteria id of the device is empty")

	}
	student,batch, err := ms.repo.GetStudentByRfidTag(rfidTag)

	

	if err != nil {

		return nil, err
	}

	cafeteriaIdInteger,_:= strconv.Atoi(cafeteriaId)

	if  cafeteriaIdInteger== batch.Cafeteria_id {

		return student, nil
	}else{

		return student,errors.New("Access Denied: Wrong Cafeteria.")
	}

	

	

}

// this method checks if the student can eat in the cafeteria
func CheckValidCafeteria(studentBatchCafeteria, deviceCafeteria string) (bool, error) {
	panic("unimplemented")
}

// checks if the current time is a meal time(breakfast,lunch,dinner)
func CheckMealTime(currentTime string) (*models.Meal, error) {
	panic("unimplemented")
}

// Grants or denies access to cafeteria for given student
func GrantOrDenyAccess(currentDate string, student *models.Student, mealId string, deviceId int) (bool, error) {
	panic("unimplemented")
}

// gets accesss logs
func GetAccessLog(date string) (*models.MealAccessLog, error) {
	panic("unimplemented")
}
