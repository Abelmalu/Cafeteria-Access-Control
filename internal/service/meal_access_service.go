package service

import (
	"github.com/abelmalu/CafeteriaAccessControl/internal/core"
	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
)

type MealAccessService struct {
	repo core.AccessRepository
}

func GetStudentByRfidTag(RfidTag string) (*models.Student, error) {
	panic("unimplemented")
}
func CheckValidCafeteria(studentBatchCafeteria, deviceCafeteria string) (bool, error) {
	panic("unimplemented")
}
func CheckMealTime(currentTime string) (*models.Meal, error) {
	panic("unimplemented")
}
func GrantOrDenyAccess(currentDate string, student *models.Student, mealId string, deviceId int) (bool, error) {
	panic("unimplemented")
}
func GetAccessLog(date string) (*models.MealAccessLog, error) {
	panic("unimplemented")
}
