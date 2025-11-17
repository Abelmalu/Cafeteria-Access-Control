package core

import "github.com/abelmalu/CafeteriaAccessControl/internal/models"

type MealAccessService interface {
	GetStudentByRfidTag(rfidTag string,cafeteriaId string) (*models.Student, error)
	//CheckValidCafeteria(studentBatchCafeteria, deviceCafeteria string) (bool, error)
	// CheckMealTime(currentTime string) (*models.Meal, error)
	// GrantOrDenyAccess(currentDate string, student *models.Student, mealId string, deviceId int) (bool, error)
	// GetAccessLog(date string) (*models.MealAccessLog, error)
	// AttemptAccess(rfidTag string)
}
