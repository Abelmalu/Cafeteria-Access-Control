package core

import "github.com/abelmalu/CafeteriaAccessControl/internal/models"

type MealAccessServiceRepository interface {
	AttemptAccess(rfidTag string) (*models.Student, *models.Batch, error)
	GetMeals() ([]models.Meal, error)
	GrantOrDenyAccess(currentDate string, studentId int, mealId int, cafeteriaId int) (string, error)
	GetCafeterias() ([]models.Cafeteria, error)
	VerifyDevice(SerialNumber string) bool
}
