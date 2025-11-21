package core

import "github.com/abelmalu/CafeteriaAccessControl/internal/models"

type MealAccessServiceRepository interface {
	AttemptAccess(rfidTag string) (*models.Student, *models.Batch, error)
	GetMeals() ([]models.Meal, error)
	GrantOrDenyAccess(currentDate string, studentId int, mealId int, cafeteriaId int) (string, error)
	GetCafeterias() ([]models.Cafeteria, error)
	VerifyDevice(SerialNumber string) bool

	// CheckValidCafeteria(studentBatchCafeteria, deviceCafeteria string) (bool, error)
	// CheckMealTime(currentTime string) (*models.Meal, error)
	// GrantOrDenyAccess(currentDate string, student *models.Student, mealId string, deviceId int) (bool, error)
	// GetAccessLog(date string) (*models.MealAccessLog, error)
}
