package core

import "github.com/abelmalu/CafeteriaAccessControl/internal/models"

type MealAccessService interface {
	AttemptAccess(rfidTag string, cafeteriaId string) (*models.Student, string, error)
	GetCafeterias() ([]models.Cafeteria, error)
	VerifyDevice(SerialNumber string) bool
}
