package core

import (
	"context"

	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
)

type AccessRepository interface {
	GetStudentByRfidTag(tag string) (*models.Student, error)
	CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error)
	CreateBatch(ctx context.Context, student *models.Student) (*models.Student, error)
	RegisterDevice(ctx context.Context, device *models.Device) (*models.Device, error)
	CreateMeal(ctx context.Context, student *models.Student) (*models.Student, error)
	CreateCafeteria(ctx context.Context, location *models.Device) (*models.Cafeteria, error)
}
