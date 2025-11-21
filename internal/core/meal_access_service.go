package core

import "github.com/abelmalu/CafeteriaAccessControl/internal/models"

type MealAccessService interface {
	AttemptAccess(rfidTag string, cafeteriaId string) (*models.Student, error)
	GetCafeterias() ([]models.Cafeteria, error)
	VerifyDevice(SerialNumber string) bool
}
