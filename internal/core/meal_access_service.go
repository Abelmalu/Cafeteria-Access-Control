package core

import "github.com/abelmalu/CafeteriaAccessControl/internal/models"

type AccessService interface {
	GetStudentByRfidTag(RfidTag string) (*models.Student, error)
	CheckValidCafeteria(studentBatchCafeteria, deviceCafeteria string) (bool, error)
	CheckMealTime(currentTime string) (*models.Meal, error)
	GrantOrDenyAccess(currentDate string, student *models.Student, mealId string, deviceId int) (bool, error)
	GetAccessLog(date string) (*models.MealAccessLog, error)
}
